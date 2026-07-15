package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_link_get_connected(res_link_t res) { return res.state == 1; }

// CGO cannot access union type fields, so we do this as a workaround
res_link_t hd_res_get_link(hd_res_t *res) { return res->link; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceLink struct {
	Type      ResourceType `json:"type"`
	Connected bool         `json:"connected"`
}

func NewResourceLink(res *C.hd_res_t, resType ResourceType) (*ResourceLink, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeLink {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeLink, resType)
	}

	link := C.hd_res_get_link(res)

	return &ResourceLink{
		Type:      resType,
		Connected: bool(C.hd_res_link_get_connected(link)),
	}, nil
}

func (r ResourceLink) ResourceType() ResourceType {
	return ResourceTypeLink
}
