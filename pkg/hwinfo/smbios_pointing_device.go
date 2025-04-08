package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPointingDevice captures pointing device (aka 'mouse') information.
type SmbiosPointingDevice struct {
	Type      SmbiosType `json:"-"`
	Handle    int        `json:"handle"`
	MouseType *ID        `json:"mouse_type"`
	Interface *ID        `json:"interface"`
	Buttons   uint16     `json:"buttons"`
}

func (s SmbiosPointingDevice) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMouse(info C.smbios_mouse_t) (*SmbiosPointingDevice, error) {
	return &SmbiosPointingDevice{
		Type:      SmbiosTypePointingDevice,
		Handle:    int(info.handle),
		MouseType: NewID(info.mtype),
		Interface: NewID(info._interface),
		Buttons:   uint16(info.buttons),
	}, nil
}
