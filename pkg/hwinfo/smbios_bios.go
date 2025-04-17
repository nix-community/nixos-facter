package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"fmt"
)

// SmbiosBios captures BIOS related information.
type SmbiosBios struct {
	Type         SmbiosType `json:"-"`
	Handle       int        `json:"handle"`
	Vendor       string     `json:"vendor"`
	Version      string     `json:"version"`
	Date         string     `json:"date"`
	Features     []string   `json:"features"`
	StartAddress string     `json:"start_address"`
	RomSize      uint32     `json:"rom_size"`
}

func (s SmbiosBios) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosBiosInfo(info C.smbios_biosinfo_t) (*SmbiosBios, error) {
	return &SmbiosBios{
		Type:         SmbiosTypeBios,
		Handle:       int(info.handle),
		Vendor:       C.GoString(info.vendor),
		Version:      C.GoString(info.version),
		Date:         C.GoString(info.date),
		Features:     ReadStringList(info.feature.str),
		StartAddress: fmt.Sprintf("0x%x", uint(info.start)),
		RomSize:      uint32(info.rom_size),
	}, nil
}
