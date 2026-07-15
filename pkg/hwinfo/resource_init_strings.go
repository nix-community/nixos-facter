package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_init_strings_t hd_res_get_init_strings(hd_res_t *res) { return res->init_strings; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceInitStrings struct {
	Type  ResourceType
	Init1 string `json:"init_1,omitempty"`
	Init2 string `json:"init_2,omitempty"`
}

func NewResourceInitStrings(res *C.hd_res_t, resType ResourceType) (*ResourceInitStrings, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeInitStrings {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeInitStrings, resType)
	}

	init := C.hd_res_get_init_strings(res)

	return &ResourceInitStrings{
		Type:  resType,
		Init1: C.GoString(init.init1),
		Init2: C.GoString(init.init2),
	}, nil
}

func (r ResourceInitStrings) ResourceType() ResourceType {
	return ResourceTypeInitStrings
}
