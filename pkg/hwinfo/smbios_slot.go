package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosSlot captures system slot information.
type SmbiosSlot struct {
	Type        SmbiosType `json:"-"`
	Handle      int        `json:"handle"`
	Designation string     `json:"designation,omitempty"`
	SlotType    *ID        `json:"slot_type"`
	BusWidth    *ID        `json:"bus_width"`
	Usage       *ID        `json:"usage"`
	Length      *ID        `json:"length"`
	ID          uint16     `json:"id"`
	Features    []string   `json:"features,omitempty"`
}

func (s SmbiosSlot) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosSlot(info C.smbios_slot_t) (*SmbiosSlot, error) {
	return &SmbiosSlot{
		Type:        SmbiosTypeSlot,
		Handle:      int(info.handle),
		Designation: C.GoString(info.desig),
		SlotType:    NewID(info.slot_type),
		BusWidth:    NewID(info.bus_width),
		Usage:       NewID(info.usage),
		Length:      NewID(info.length),
		ID:          uint16(info.id),
		Features:    ReadStringList(info.feature.str),
	}, nil
}
