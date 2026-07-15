package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_io_get_enabled(res_io_t res) { return res.enabled; }
unsigned hd_res_io_get_access(res_io_t res) { return res.access; }

// CGO cannot access union type fields, so we do this as a workaround
res_io_t hd_res_get_io(hd_res_t *res) { return res->io; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceIO struct {
	Type    ResourceType `json:"type"`
	Base    uint64       `json:"base"`
	Range   uint64       `json:"range"`
	Enabled bool         `json:"enabled"`
	Access  AccessFlags  `json:"access"`
}

func NewResourceIO(res *C.hd_res_t, resType ResourceType) (*ResourceIO, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeIo {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeIo, resType)
	}

	io := C.hd_res_get_io(res)

	return &ResourceIO{
		Type:    resType,
		Base:    uint64(io.base),
		Range:   uint64(io._range),
		Enabled: bool(C.hd_res_io_get_enabled(io)),
		Access:  AccessFlags(C.hd_res_io_get_access(io)),
	}, nil
}

func (r ResourceIO) ResourceType() ResourceType {
	return ResourceTypeIo
}
