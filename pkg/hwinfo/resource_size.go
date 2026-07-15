package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_size_t hd_res_get_size(hd_res_t *res) { return res->size; }

*/
import "C"

import (
	"errors"
	"fmt"
)

//go:generate enumer -type=SizeUnit -json --transform=snake -trimprefix SizeUnit --output=./resource_enum_size_unit.go
type SizeUnit uint //nolint:recvcheck

const (
	SizeUnitCm SizeUnit = iota
	SizeUnitCinch
	SizeUnitByte
	SizeUnitSectors

	SizeUnitKbyte
	SizeUnitMbyte
	SizeUnitGbyte
	SizeUnitMm
)

type ResourceSize struct {
	Type   ResourceType `json:"type"`
	Unit   SizeUnit     `json:"unit"`
	Value1 uint64       `json:"value_1"`
	Value2 uint64       `json:"value_2,omitempty"`
}

func NewResourceSize(res *C.hd_res_t, resType ResourceType) (*ResourceSize, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeSize {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeSize, resType)
	}

	size := C.hd_res_get_size(res)

	return &ResourceSize{
		Type:   resType,
		Unit:   SizeUnit(size.unit),
		Value1: uint64(size.val1),
		Value2: uint64(size.val2),
	}, nil
}

func (r ResourceSize) ResourceType() ResourceType {
	return ResourceTypeSize
}
