package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosOEMStrings captures OEM information.
type SmbiosOEMStrings struct {
	Type    SmbiosType `json:"-"`
	Handle  int        `json:"handle"`
	Strings []string   `json:"strings,omitempty"`
}

func NewSmbiosOEM(info C.smbios_oem_t) (*SmbiosOEMStrings, error) {
	return &SmbiosOEMStrings{
		Type:    SmbiosTypeOEMStrings,
		Handle:  int(info.handle),
		Strings: ReadStringList(info.oem_strings),
	}, nil
}

func (s SmbiosOEMStrings) SmbiosType() SmbiosType {
	return s.Type
}
