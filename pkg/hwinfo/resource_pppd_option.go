package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_pppd_option_t hd_res_get_pppd_option(hd_res_t *res) { return res->pppd_option; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourcePppdOption struct {
	Type   ResourceType `json:"type"`
	Option byte         `json:"option"`
}

func (r ResourcePppdOption) ResourceType() ResourceType {
	return ResourceTypePppdOption
}

func NewResourcePppdOption(res *C.hd_res_t, resType ResourceType) (*ResourcePppdOption, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypePppdOption {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypePppdOption, resType)
	}

	pppd := C.hd_res_get_pppd_option(res)
	var option byte
	if pppd.option != nil {
		option = byte(*pppd.option)
	}

	return &ResourcePppdOption{
		Type:   resType,
		Option: option,
	}, nil
}
