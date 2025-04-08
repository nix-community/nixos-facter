package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_dma_get_enabled(res_dma_t res) { return res.enabled; }

// CGO cannot access union type fields, so we do this as a workaround
res_dma_t hd_res_get_dma(hd_res_t *res) { return res->dma; }

*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceDma struct {
	Type    ResourceType `json:"type"`
	Base    string       `json:"base"`
	Enabled bool         `json:"enabled"`
}

func (r ResourceDma) ResourceType() ResourceType {
	return ResourceTypeDma
}

func NewResourceDma(res *C.hd_res_t, resType ResourceType) (*ResourceDma, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeDma {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeDma, resType)
	}

	dma := C.hd_res_get_dma(res)

	return &ResourceDma{
		Type:    resType,
		Base:    fmt.Sprintf("0x%x", uint(dma.base)),
		Enabled: bool(C.hd_res_dma_get_enabled(dma)),
	}, nil
}
