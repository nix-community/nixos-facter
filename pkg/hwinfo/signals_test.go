//nolint:testpackage
package hwinfo

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"unsafe"
)

// kernelSigaction mirrors the kernel's struct sigaction on linux/amd64.
type kernelSigaction struct {
	handler  uintptr
	flags    uint64
	restorer uintptr
	mask     uint64
}

const saOnStack = 0x08000000

// mangleSigchld replicates libhd's bug (hd.c hd_fork): it re-registers the
// current SIGCHLD handler with the SA_ONSTACK flag stripped, exactly as
// restoring a handler through ISO C signal() does.
func mangleSigchld(t *testing.T) {
	t.Helper()

	var sa kernelSigaction

	// read the current registration; raw rt_sigaction is deliberate - this
	// replicates libhd's flag-stripping bug, which no safe API can express
	if _, _, errno := syscall.Syscall6(
		syscall.SYS_RT_SIGACTION, uintptr(syscall.SIGCHLD),
		0, uintptr(unsafe.Pointer(&sa)), 8, 0, 0, //nolint:gosec
	); errno != 0 {
		t.Fatalf("rt_sigaction read failed: %v", errno)
	}

	if sa.flags&saOnStack == 0 {
		t.Fatal("expected the Go runtime's SIGCHLD handler to have SA_ONSTACK")
	}

	sa.flags &^= saOnStack

	if _, _, errno := syscall.Syscall6(
		syscall.SYS_RT_SIGACTION, uintptr(syscall.SIGCHLD),
		uintptr(unsafe.Pointer(&sa)), 0, 8, 0, 0, //nolint:gosec
	); errno != 0 {
		t.Fatalf("rt_sigaction write failed: %v", errno)
	}
}

// Regression test for https://github.com/nix-community/nixos-facter/issues/216:
// libhd's hd_fork (modem, mouse and braille probes) restores the Go runtime's
// SIGCHLD handler without SA_ONSTACK; the next exec'd child then kills the
// process. The scan must snapshot and restore signal handlers around hd_scan.
//
// A process whose handler is broken dies unrecoverably, so the scenario runs
// in a re-exec'd copy of the test binary and the parent asserts its fate.
func TestSignalHandlerRestore(t *testing.T) {
	switch os.Getenv("FACTER_SIGNAL_TEST") {
	case "mangled":
		mangleSigchld(t)

		_ = exec.CommandContext(t.Context(), "/proc/self/exe", "-test.run=NONE").Run() // deliver a SIGCHLD

		os.Exit(0)
	case "restored":
		saveSignalHandlers()
		mangleSigchld(t)
		restoreSignalHandlers()

		_ = exec.CommandContext(t.Context(), "/proc/self/exe", "-test.run=NONE").Run()

		os.Exit(0)
	}

	t.Parallel()

	run := func(mode string) (string, error) {
		// re-exec this test binary; the mode is a fixed string, not tainted input
		cmd := exec.CommandContext(t.Context(), os.Args[0], "-test.run=TestSignalHandlerRestore") //nolint:gosec

		cmd.Env = append(os.Environ(), "FACTER_SIGNAL_TEST="+mode)
		out, err := cmd.CombinedOutput()

		return string(out), err
	}

	// control: the mangled handler must reproduce the crash from #216,
	// otherwise this test is not exercising anything
	out, err := run("mangled")
	if err == nil || !strings.Contains(out, "SA_ONSTACK") {
		t.Fatalf("expected SA_ONSTACK crash in mangled child, got err=%v output:\n%s", err, out)
	}

	// with save/restore around the mangle, the process must survive
	out, err = run("restored")
	if err != nil {
		t.Fatalf("restored child should survive: %v output:\n%s", err, out)
	}
}
