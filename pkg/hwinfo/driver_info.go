package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

enum driver_info_type driver_info_get_type(driver_info_t *info) { return info->any.type; }
driver_info_module_t driver_info_get_module(driver_info_t *info) { return info->module; }
driver_info_mouse_t driver_info_get_mouse(driver_info_t *info) { return info->mouse; }
driver_info_display_t driver_info_get_display(driver_info_t *info) { return info->display; }
driver_info_kbd_t driver_info_get_kbd(driver_info_t *info) { return info->kbd; }
driver_info_dsl_t driver_info_get_dsl(driver_info_t *info) { return info->dsl; }
driver_info_isdn_t driver_info_get_isdn(driver_info_t *info) { return info->isdn; }
driver_info_x11_t driver_info_get_x11(driver_info_t *info) { return info->x11; }
*/
import "C"

import (
	"errors"
)

//go:generate enumer -type=DriverInfoType -json -transform=snake -trimprefix DriverInfoType -output=./driver_info_enum_type.go
type DriverInfoType uint //nolint:recvcheck

const (
	DriverInfoTypeAny DriverInfoType = iota
	DriverInfoTypeDisplay
	DriverInfoTypeModule
	DriverInfoTypeMouse
	DriverInfoTypeX11
	DriverInfoTypeIsdn
	DriverInfoTypeKeyboard
	DriverInfoTypeDsl
)

type DriverInfo interface {
	DriverInfoType() DriverInfoType
}

//nolint:ireturn
func NewDriverInfo(info *C.driver_info_t) (DriverInfo, error) {
	if info == nil {
		return nil, errors.New("info is nil")
	}

	var (
		err    error
		result DriverInfo
	)

	switch DriverInfoType(C.driver_info_get_type(info)) {
	case DriverInfoTypeModule:
		result, err = NewDriverInfoModule(C.driver_info_get_module(info)), nil
	case DriverInfoTypeMouse:
		result, err = NewDriverInfoMouse(C.driver_info_get_mouse(info)), nil
	case DriverInfoTypeDisplay:
		result, err = NewDriverInfoDisplay(C.driver_info_get_display(info)), nil
	case DriverInfoTypeKeyboard:
		result, err = NewDriverInfoKeyboard(C.driver_info_get_kbd(info)), nil
	case DriverInfoTypeDsl:
		result, err = NewDriverInfoDsl(C.driver_info_get_dsl(info)), nil
	case DriverInfoTypeIsdn:
		result, err = NewDriverInfoIsdn(C.driver_info_get_isdn(info)), nil
	case DriverInfoTypeX11:
		result, err = NewDriverInfoX11(C.driver_info_get_x11(info)), nil
	case DriverInfoTypeAny:
		err = errors.New("unknown driver info type")
	}

	return result, err
}
