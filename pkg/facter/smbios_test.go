//nolint:testpackage
package facter

import (
	"testing"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

// Regression test for https://github.com/nix-community/nixos-facter/issues/623:
// some firmware (e.g. Apple EFI on a 2014 Mac Mini) presents malformed SMBIOS
// tables from which libhd decodes duplicate singleton structures. A duplicate
// must not abort the scan; we keep the first entry and warn.
func TestSmbiosAddDuplicateSingletons(t *testing.T) {
	t.Parallel()

	var s Smbios

	first := &hwinfo.SmbiosBios{Type: hwinfo.SmbiosTypeBios, Handle: 0x0B, Vendor: "Apple Inc."}
	second := &hwinfo.SmbiosBios{Type: hwinfo.SmbiosTypeBios, Handle: 0x0C, Vendor: "Phantom"}

	if err := s.add(first); err != nil {
		t.Fatalf("add first bios: %v", err)
	}

	if err := s.add(second); err != nil {
		t.Fatalf("duplicate bios entry must not error: %v", err)
	}

	if s.Bios != first {
		t.Fatalf("expected first bios entry to be kept, got %+v", s.Bios)
	}

	if err := s.add(&hwinfo.SmbiosSystem{Type: hwinfo.SmbiosTypeSystem}); err != nil {
		t.Fatalf("add first system: %v", err)
	}

	if err := s.add(&hwinfo.SmbiosSystem{Type: hwinfo.SmbiosTypeSystem}); err != nil {
		t.Fatalf("duplicate system entry must not error: %v", err)
	}

	if err := s.add(&hwinfo.SmbiosBoard{Type: hwinfo.SmbiosTypeBoard}); err != nil {
		t.Fatalf("add first board: %v", err)
	}

	if err := s.add(&hwinfo.SmbiosBoard{Type: hwinfo.SmbiosTypeBoard}); err != nil {
		t.Fatalf("duplicate board entry must not error: %v", err)
	}

	if err := s.add(&hwinfo.SmbiosConfig{Type: hwinfo.SmbiosTypeConfig}); err != nil {
		t.Fatalf("add first config: %v", err)
	}

	if err := s.add(&hwinfo.SmbiosConfig{Type: hwinfo.SmbiosTypeConfig}); err != nil {
		t.Fatalf("duplicate config entry must not error: %v", err)
	}
}
