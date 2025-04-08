package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_irq_get_enabled(res_irq_t res) { return res.enabled; }

// CGO cannot access union type fields, so we do this as a workaround
res_irq_t hd_res_get_irq(hd_res_t *res) { return res->irq; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceIrq struct {
	Type      ResourceType `json:"type"`
	Base      uint16       `json:"base"`
	Triggered uint16       `json:"triggered"`
	Enabled   bool         `json:"enabled"`
}

func (r ResourceIrq) ResourceType() ResourceType {
	return ResourceTypeIrq
}

func NewResourceIrq(res *C.hd_res_t, resType ResourceType) (*ResourceIrq, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeIrq {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeIrq, resType)
	}

	irq := C.hd_res_get_irq(res)

	return &ResourceIrq{
		Type:      resType,
		Base:      uint16(irq.base),
		Triggered: uint16(irq.triggered),
		Enabled:   bool(C.hd_res_irq_get_enabled(irq)),
	}, nil
}
