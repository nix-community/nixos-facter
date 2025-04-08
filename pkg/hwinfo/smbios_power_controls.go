package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPowerControls captures system power controls information.
type SmbiosPowerControls struct {
	Type   SmbiosType `json:"-"`
	Handle int        `json:"handle"`
	Month  uint8      `json:"month"`  // next scheduled power-on month
	Day    uint16     `json:"day"`    // dto, day
	Hour   uint8      `json:"hour"`   // dto, hour
	Minute uint8      `json:"minute"` // dto, minute
	Second uint8      `json:"second"` // dto, second
}

func (s SmbiosPowerControls) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosPower(info C.smbios_power_t) (*SmbiosPowerControls, error) {
	return &SmbiosPowerControls{
		Type:   SmbiosTypePowerControls,
		Handle: int(info.handle),
		Month:  uint8(info.month),
		Day:    uint16(info.day),
		Hour:   uint8(info.hour),
		Minute: uint8(info.minute),
		Second: uint8(info.second),
	}, nil
}
