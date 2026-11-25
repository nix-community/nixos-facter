//nolint:testpackage
package hwinfo

import (
	"testing"

	"github.com/numtide/nixos-facter/pkg/linux/input"
)

// Regression test for https://github.com/nix-community/nixos-facter/issues/339:
// Apple attaches touchpads via SPI on Intel MacBooks and Apple Silicon alike,
// and an unrecognised input bus must never abort the whole scan.
func TestMouseBusIDs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		bus          input.Bus
		wantBus      Bus
		wantSubClass SubClassMouse
	}{
		{name: "i2c", bus: input.BusI2c, wantBus: BusSerial, wantSubClass: SubClassMouseSer},
		{name: "usb", bus: input.BusUsb, wantBus: BusUsb, wantSubClass: SubClassMouseUsb},
		{name: "ps2", bus: input.BusI8042, wantBus: BusPs2, wantSubClass: SubClassMousePs2},
		{name: "rmi", bus: input.BusRmi, wantBus: BusOther, wantSubClass: SubClassMouseOther},
		{name: "bluetooth", bus: input.BusBluetooth, wantBus: BusBluetooth, wantSubClass: SubClassMouseOther},
		{name: "host", bus: input.BusHost, wantBus: BusHost, wantSubClass: SubClassMouseOther},
		// issue #339: MacBook touchpads
		{name: "spi", bus: input.BusSpi, wantBus: BusOther, wantSubClass: SubClassMouseOther},
		// unknown buses degrade gracefully instead of aborting the scan
		{name: "gameport", bus: input.BusGameport, wantBus: BusOther, wantSubClass: SubClassMouseOther},
		{name: "intel ishtp", bus: input.BusIntelIshtp, wantBus: BusOther, wantSubClass: SubClassMouseOther},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			busType, subClass := mouseBusIDs(tc.bus)

			if busType == nil || subClass == nil {
				t.Fatal("expected non-nil bus type and sub class")
			}

			if busType.Value != uint16(tc.wantBus) {
				t.Fatalf("bus type: expected %d, got %d", tc.wantBus, busType.Value)
			}

			if subClass.Value != uint16(tc.wantSubClass) {
				t.Fatalf("sub class: expected %d, got %d", tc.wantSubClass, subClass.Value)
			}
		})
	}
}
