package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "unsafe"

// SmbiosBoard captures motherboard related information.
type SmbiosBoard struct {
	Type         SmbiosType `json:"-"`
	Handle       int        `json:"handle"`
	Manufacturer string     `json:"manufacturer"`
	Product      string     `json:"product"`
	Version      string     `json:"version"`
	Serial       string     `json:"-"` // omit from json output
	AssetTag     string     `json:"-"`
	BoardType    *ID        `json:"board_type"`
	Features     []string   `json:"features"`
	Location     string     `json:"location"`          // location in chassis
	Chassis      int        `json:"chassis"`           // handle of chassis
	Objects      []int      `json:"objects,omitempty"` // array of object handles
}

func NewSmbiosBoardInfo(info C.smbios_boardinfo_t) (*SmbiosBoard, error) {
	return &SmbiosBoard{
		Type:         SmbiosTypeBoard,
		Handle:       int(info.handle),
		Manufacturer: C.GoString(info.manuf),
		Product:      C.GoString(info.product),
		Version:      C.GoString(info.version),
		Serial:       C.GoString(info.serial),
		AssetTag:     C.GoString(info.asset),
		BoardType:    NewID(info.board_type),
		Features:     ReadStringList(info.feature.str),
		Location:     C.GoString(info.location),
		Chassis:      int(info.chassis),
		Objects:      ReadIntArray(unsafe.Pointer(info.objects), int(info.objects_len)),
	}, nil
}

func (s SmbiosBoard) SmbiosType() SmbiosType {
	return s.Type
}
