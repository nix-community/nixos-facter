package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool driver_info_x11_supports_3d(driver_info_x11_t info) { return info.x3d; }
char driver_info_x11_colors_all(driver_info_x11_t info) { return info.colors.all; }
char driver_info_x11_colors_c8(driver_info_x11_t info) { return info.colors.c8; }
char driver_info_x11_colors_c15(driver_info_x11_t info) { return info.colors.c15; }
char driver_info_x11_colors_c16(driver_info_x11_t info) { return info.colors.c16; }
char driver_info_x11_colors_c24(driver_info_x11_t info) { return info.colors.c24; }
char driver_info_x11_colors_c32(driver_info_x11_t info) { return info.colors.c32; }
*/
import "C"

type DriverInfoX11 struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DBEntry0 []string `json:"db_entry_0,omitempty"`
	DBEntry1 []string `json:"db_entry_1,omitempty"`

	Server      string `json:"server,omitempty"`       // the server/module name
	XF86Version string `json:"xf86_version,omitempty"` // XFree86 version (3 or 4)
	Supports3D  bool   `json:"supports_3d"`            // has 3D support
	Colors      struct {
		// the next 5 entries combined
		All byte `json:"all"`
		C8  byte `json:"c8"`
		C15 byte `json:"c15"`
		C16 byte `json:"c16"`
		C24 byte `json:"c24"`
		C32 byte `json:"c32"`
	} `json:",omitempty"`
	DacSpeed   uint32   `json:"dac_speed"`            // max. ramdac clock
	Extensions []string `json:"extensions,omitempty"` // additional X extensions to load ('Module' section)
	Options    []string `json:"options,omitempty"`    // special server options
	Raw        []string `json:"raw,omitempty"`        // extra info to add to XF86Config
	Script     string   `json:"script"`               // 3d script to run
}

func (d DriverInfoX11) DriverInfoType() DriverInfoType {
	return DriverInfoTypeX11
}

func NewDriverInfoX11(info C.driver_info_x11_t) DriverInfoX11 {
	result := DriverInfoX11{
		Type:        DriverInfoTypeX11,
		DBEntry0:    ReadStringList(info.hddb0),
		DBEntry1:    ReadStringList(info.hddb1),
		Server:      C.GoString(info.server),
		XF86Version: C.GoString(info.xf86_ver),
		Supports3D:  bool(C.driver_info_x11_supports_3d(info)),
		DacSpeed:    uint32(info.dacspeed),
		Extensions:  ReadStringList(info.extensions),
		Options:     ReadStringList(info.options),
		Raw:         ReadStringList(info.raw),
		Script:      C.GoString(info.script),
	}

	result.Colors.All = byte(C.driver_info_x11_colors_all(info))
	result.Colors.C8 = byte(C.driver_info_x11_colors_c8(info))
	result.Colors.C15 = byte(C.driver_info_x11_colors_c15(info))
	result.Colors.C16 = byte(C.driver_info_x11_colors_c16(info))
	result.Colors.C24 = byte(C.driver_info_x11_colors_c24(info))
	result.Colors.C32 = byte(C.driver_info_x11_colors_c32(info))

	return result
}
