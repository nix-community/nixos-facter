package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosSystem captures overall system related information.
type SmbiosSystem struct {
	Type         SmbiosType `json:"-"`
	Handle       int        `json:"handle"`
	Manufacturer string     `json:"manufacturer"`
	Product      string     `json:"product"`
	Version      string     `json:"version"`
	Serial       string     `json:"-"`       // omit from json output
	UUID         string     `json:"-"`       // universal unique id; all 0x00: undef, all 0xff: undef but settable
	WakeUp       *ID        `json:"wake_up"` // wake-up type
}

func NewSmbiosSysInfo(info C.smbios_sysinfo_t) (*SmbiosSystem, error) {
	return &SmbiosSystem{
		Type:         SmbiosTypeSystem,
		Handle:       int(info.handle),
		Manufacturer: C.GoString(info.manuf),
		Product:      C.GoString(info.product),
		Version:      C.GoString(info.version),
		Serial:       C.GoString(info.serial),
		// todo uuid
		WakeUp: NewID(info.wake_up),
	}, nil
}

func (s SmbiosSystem) SmbiosType() SmbiosType {
	return s.Type
}
