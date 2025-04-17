package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_baud_t hd_res_get_baud(hd_res_t *res) { return res->baud; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceBaud struct {
	Type      ResourceType `json:"type"`
	Speed     uint32       `json:"speed"`
	Bits      uint32       `json:"bits"`
	StopBits  uint32       `json:"stop_bits"`
	Parity    byte         `json:"parity"`
	Handshake byte         `json:"handshake"`
}

func (r ResourceBaud) ResourceType() ResourceType {
	return ResourceTypeBaud
}

func NewResourceBaud(res *C.hd_res_t, resType ResourceType) (*ResourceBaud, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeBaud {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeBaud, resType)
	}

	baud := C.hd_res_get_baud(res)

	return &ResourceBaud{
		Type:      resType,
		Speed:     uint32(baud.speed),
		Bits:      uint32(baud.bits),
		StopBits:  uint32(baud.stopbits),
		Parity:    byte(baud.parity),
		Handshake: byte(baud.handshake),
	}, nil
}
