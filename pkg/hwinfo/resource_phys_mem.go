package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_phys_mem_t hd_res_get_phys_mem(hd_res_t *res) { return res->phys_mem; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourcePhysicalMemory struct {
	Type  ResourceType `json:"type"`
	Range uint64       `json:"range"`
}

func (r ResourcePhysicalMemory) ResourceType() ResourceType {
	return ResourceTypePhysMem
}

func NewResourcePhysicalMemory(res *C.hd_res_t, resType ResourceType) (*ResourcePhysicalMemory, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypePhysMem {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypePhysMem, resType)
	}

	physMem := C.hd_res_get_phys_mem(res)

	return &ResourcePhysicalMemory{
		Type:  resType,
		Range: uint64(physMem._range),
	}, nil
}
