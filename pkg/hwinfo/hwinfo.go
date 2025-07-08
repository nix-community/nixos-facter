// Package hwinfo provides functionality for scanning hardware devices and reading SMBIOS information.
// It does this using the [hwinfo] C library.
//
// [hwinfo]: https://github.com/numtide/hwinfo
package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdlib.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_smbios_t* hd_smbios_next(hd_smbios_t *sm) { return sm->next; }
*/
import "C"

import (
	"unsafe"
)

func excludeDevice(item *HardwareDevice) bool {
	if item.Class == HardwareClassNetworkInterface {
		for _, driver := range item.Drivers {
			// devices that are not mapped to hardware should be not included in the hardware report
			if virtualNetworkDevices[driver] {
				return true
			}
		}
	}

	return false
}

// Scan returns a list of SMBIOS entries and detected hardware devices based on the provided probe features.
func Scan(probes []ProbeFeature, ephemeral bool) ([]Smbios, []HardwareDevice, error) {
	// initialise the struct to hold scan data
	data := (*C.hd_data_t)(C.calloc(1, C.size_t(unsafe.Sizeof(C.hd_data_t{}))))

	// ProbeFeatureInt needs to always be set, otherwise we don't get pci and usb vendor id lookups.
	// https://github.com/openSUSE/hwinfo/blob/c87f449f1d4882c71b0a1e6dc80638224a5baeed/src/hd/hd.c#L597-L605
	C.hd_set_probe_feature(data, C.enum_probe_feature(ProbeFeatureInt))

	// set the hardware probes to run
	for _, probe := range probes {
		C.hd_set_probe_feature(data, C.enum_probe_feature(probe))
	}

	// scan
	C.hd_scan(data)

	defer C.hd_free_hd_data(data)

	var smbiosItems []Smbios

	for sm := data.smbios; sm != nil; sm = C.hd_smbios_next(sm) {
		item, err := NewSmbios(sm)
		if err != nil {
			return nil, nil, err
		} else if item == nil {
			continue
		}

		smbiosItems = append(smbiosItems, item)
	}

	var (
		deviceIdx     uint16
		hardwareItems []HardwareDevice
	)

	for hd := data.hd; hd != nil; hd = hd.next {
		item, err := NewHardwareDevice(hd, ephemeral)
		if err != nil {
			return nil, nil, err
		}

		if item.Index > deviceIdx {
			deviceIdx = item.Index
		}

		if excludeDevice(item) {
			continue
		}

		hardwareItems = append(hardwareItems, *item)
	}

	// probe for additional inputs that hwinfo does not support
	touchpads, err := captureTouchpads(deviceIdx + 1)
	if err != nil {
		return nil, nil, err
	}

	hardwareItems = append(hardwareItems, touchpads...)

	return smbiosItems, hardwareItems, nil
}
