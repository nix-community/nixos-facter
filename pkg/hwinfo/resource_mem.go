package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_mem_get_enabled(res_mem_t res) { return res.enabled; }
unsigned hd_res_mem_get_access(res_mem_t res) { return res.access; }
unsigned hd_res_mem_get_prefetch(res_mem_t res) { return res.prefetch; }

// CGO cannot access union type fields, so we do this as a workaround
res_mem_t hd_res_get_mem(hd_res_t *res) { return res->mem; }

*/
import "C"

import (
	"errors"
	"fmt"
)

//go:generate enumer -type=AccessFlags -json -transform=snake -trimprefix AccessFlags -output=./resource_enum_access_flags.go
type AccessFlags uint //nolint:recvcheck

const (
	AccessFlagsUnknown AccessFlags = iota
	AccessFlagsReadOnly
	AccessFlagsWriteOnly
	AccessFlagsReadWrite
)

//go:generate enumer -type=YesNoFlags -json -transform=snake -trimprefix YesNoFlags -output=./resource_enum_yes_no_flags.go
type YesNoFlags uint //nolint:recvcheck

const (
	YesNoFlagsUnknown YesNoFlags = iota
	YesNoFlagsNo
	YesNoFlagsYes
)

type Resource interface {
	ResourceType() ResourceType
}

type ResourceMemory struct {
	Type     ResourceType `json:"type"`
	Base     uint64       `json:"base"`
	Range    uint64       `json:"range"`
	Enabled  bool         `json:"enabled"`
	Access   AccessFlags  `json:"access"`
	Prefetch YesNoFlags   `json:"prefetch"`
}

func NewResourceMemory(res *C.hd_res_t, resType ResourceType) (*ResourceMemory, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeMem {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeMem, resType)
	}

	mem := C.hd_res_get_mem(res)

	return &ResourceMemory{
		Type:     resType,
		Base:     uint64(mem.base),
		Range:    uint64(mem._range),
		Enabled:  bool(C.hd_res_mem_get_enabled(mem)),
		Access:   AccessFlags(C.hd_res_mem_get_access(mem)),
		Prefetch: YesNoFlags(C.hd_res_mem_get_prefetch(mem)),
	}, nil
}

func (r ResourceMemory) ResourceType() ResourceType {
	return ResourceTypeMem
}
