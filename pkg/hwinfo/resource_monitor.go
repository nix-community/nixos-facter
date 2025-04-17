package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_monitor_get_interlaced(res_monitor_t res) { return res.interlaced; }

// CGO cannot access union type fields, so we do this as a workaround
res_monitor_t hd_res_get_monitor(hd_res_t *res) { return res->monitor; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceMonitor struct {
	Type              ResourceType `json:"type"`
	Width             uint32       `json:"width"`
	Height            uint32       `json:"height"`
	VerticalFrequency uint16       `json:"vertical_frequency"`
	Interlaced        bool         `json:"interlaced"`
}

func (r ResourceMonitor) ResourceType() ResourceType {
	return ResourceTypeMonitor
}

func NewResourceMonitor(res *C.hd_res_t, resType ResourceType) (*ResourceMonitor, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeMonitor {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeMonitor, resType)
	}

	monitor := C.hd_res_get_monitor(res)

	return &ResourceMonitor{
		Type:              resType,
		Width:             uint32(monitor.width),
		Height:            uint32(monitor.height),
		VerticalFrequency: uint16(monitor.vfreq),
		Interlaced:        bool(C.hd_res_monitor_get_interlaced(monitor)),
	}, nil
}
