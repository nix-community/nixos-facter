package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_wlan_t hd_res_get_wlan(hd_res_t *res) { return res->wlan; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceWlan struct {
	Type        ResourceType `json:"type"`
	Channels    []string     `json:"channels,omitempty"`
	Frequencies []string     `json:"frequencies,omitempty"`
	BitRates    []string     `json:"bit_rates,omitempty"`
	AuthModes   []string     `json:"auth_modes,omitempty"`
	EncModes    []string     `json:"enc_modes,omitempty"`
}

func NewResourceWlan(res *C.hd_res_t, resType ResourceType) (*ResourceWlan, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeWlan {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeWlan, resType)
	}

	wlan := C.hd_res_get_wlan(res)

	return &ResourceWlan{
		Type:        resType,
		Channels:    ReadStringList(wlan.channels),
		Frequencies: ReadStringList(wlan.frequencies),
		BitRates:    ReadStringList(wlan.bitrates),
		AuthModes:   ReadStringList(wlan.auth_modes),
		EncModes:    ReadStringList(wlan.enc_modes),
	}, nil
}

func (r ResourceWlan) ResourceType() ResourceType {
	return ResourceTypeWlan
}
