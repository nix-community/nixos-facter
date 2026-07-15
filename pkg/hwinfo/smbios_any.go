package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// SmbiosAny captures generic smbios data.
type SmbiosAny struct {
	Type    SmbiosType `json:"-"`
	Handle  int        `json:"handle"`
	Data    string     `json:"data"`
	Strings []string   `json:"strings,omitempty"`
}

func NewSmbiosAny(smbiosType SmbiosType, info C.smbios_any_t) (*SmbiosAny, error) {
	return &SmbiosAny{
		Type:    smbiosType,
		Handle:  int(info.handle),
		Data:    fmt.Sprintf("0x%x", ReadByteArray(unsafe.Pointer(info.data), int(info.data_len))),
		Strings: ReadStringList(info.strings),
	}, nil
}

func (s SmbiosAny) SmbiosType() SmbiosType {
	return s.Type
}
