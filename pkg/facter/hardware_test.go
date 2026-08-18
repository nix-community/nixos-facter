//nolint:testpackage
package facter

import (
	"testing"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

func cpuDevice(physicalID uint16) hwinfo.HardwareDevice {
	return hwinfo.HardwareDevice{
		Class:  hwinfo.HardwareClassCpu,
		Detail: &hwinfo.DetailCPU{PhysicalID: physicalID},
	}
}

// Regression test for https://github.com/nix-community/nixos-facter/issues/276:
// sparse physical ids (e.g. VMware with one core per socket assigns ids 0, 2, 4, ...)
// must not leave nil holes in Hardware.CPU, which serialise as null in the report.
func TestHardwareAddCPUSparsePhysicalIDs(t *testing.T) {
	t.Parallel()

	var h Hardware

	for _, id := range []uint16{4, 0, 2} {
		if err := h.add(cpuDevice(id)); err != nil {
			t.Fatalf("add: %v", err)
		}
	}

	if len(h.CPU) != 3 {
		t.Fatalf("expected 3 cpu entries, got %d", len(h.CPU))
	}

	for i, want := range []uint16{0, 2, 4} {
		if h.CPU[i] == nil {
			t.Fatalf("cpu entry %d is nil", i)
		}

		if h.CPU[i].PhysicalID != want {
			t.Fatalf("cpu entry %d: expected physical id %d, got %d", i, want, h.CPU[i].PhysicalID)
		}
	}
}

// One entry per physical id: sibling logical processors must not duplicate sockets.
func TestHardwareAddCPUDeduplicatesPhysicalID(t *testing.T) {
	t.Parallel()

	var h Hardware

	for _, id := range []uint16{0, 0, 1, 1} {
		if err := h.add(cpuDevice(id)); err != nil {
			t.Fatalf("add: %v", err)
		}
	}

	if len(h.CPU) != 2 {
		t.Fatalf("expected 2 cpu entries, got %d", len(h.CPU))
	}
}

// CPU devices with unavailable detail data (old hypervisors) are skipped.
func TestHardwareAddCPUNilDetail(t *testing.T) {
	t.Parallel()

	var h Hardware

	err := h.add(hwinfo.HardwareDevice{
		Class:  hwinfo.HardwareClassCpu,
		Detail: (*hwinfo.DetailCPU)(nil),
	})
	if err != nil {
		t.Fatalf("add: %v", err)
	}

	if len(h.CPU) != 0 {
		t.Fatalf("expected no cpu entries, got %d", len(h.CPU))
	}
}
