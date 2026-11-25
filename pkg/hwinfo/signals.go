package hwinfo

/*
#include <signal.h>

// libhd's hd_fork saves and restores the SIGCHLD handler with ISO C signal()
// (src/hd/hd.c), which re-registers the Go runtime's handler *without* the
// SA_ONSTACK and SA_SIGINFO flags it was installed with. The next SIGCHLD
// delivered after a scan - e.g. from exec'ing systemd-detect-virt - then
// aborts the process with "non-Go code set up signal handler without
// SA_ONSTACK flag" (#216). We cannot fix libhd's internals, but we own the
// boundary: snapshot the full sigaction for every standard signal before
// hd_scan and restore it verbatim afterwards.

#define FACTER_NSIG 32

static struct sigaction facter_saved_handlers[FACTER_NSIG];

static void facter_save_signal_handlers(void) {
	for (int sig = 1; sig < FACTER_NSIG; sig++) {
		sigaction(sig, NULL, &facter_saved_handlers[sig]);
	}
}

static void facter_restore_signal_handlers(void) {
	for (int sig = 1; sig < FACTER_NSIG; sig++) {
		// SIGKILL and SIGSTOP cannot be queried or set; sigaction fails for
		// them on both save and restore, which is harmless.
		sigaction(sig, &facter_saved_handlers[sig], NULL);
	}
}
*/
import "C"

// saveSignalHandlers snapshots the process signal handlers so they can be
// restored after libhd has run.
func saveSignalHandlers() {
	C.facter_save_signal_handlers()
}

// restoreSignalHandlers restores the signal handlers captured by
// saveSignalHandlers, undoing any registrations libhd performed or mangled.
func restoreSignalHandlers() {
	C.facter_restore_signal_handlers()
}
