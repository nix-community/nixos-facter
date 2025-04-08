package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_framebuffer_t hd_res_get_framebuffer(hd_res_t *res) { return res->framebuffer; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceFrameBuffer struct {
	Type         ResourceType `json:"type"`
	Width        uint32       `json:"width"`
	Height       uint32       `json:"height"`
	BytesPerLine uint16       `json:"bytes_per_line"`
	ColorBits    uint16       `json:"color_bits"`
	Mode         uint16       `json:"mode"`
}

func (r ResourceFrameBuffer) ResourceType() ResourceType {
	return ResourceTypeFramebuffer
}

func NewResourceFrameBuffer(res *C.hd_res_t, resType ResourceType) (*ResourceFrameBuffer, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeFramebuffer {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeFramebuffer, resType)
	}

	fb := C.hd_res_get_framebuffer(res)

	return &ResourceFrameBuffer{
		Type:         resType,
		Width:        uint32(fb.width),
		Height:       uint32(fb.height),
		BytesPerLine: uint16(fb.bytes_p_line),
		ColorBits:    uint16(fb.colorbits),
		Mode:         uint16(fb.mode),
	}, nil
}
