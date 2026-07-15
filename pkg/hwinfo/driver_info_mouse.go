package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoMouse struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DBEntry0 []string `json:"db_entry_0,omitempty"`
	DBEntry1 []string `json:"db_entry_1,omitempty"`

	XF86    string `json:"xf86,omitempty"`    // XF86 protocol name
	GPM     string `json:"gpm,omitempty"`     // dto, gpm
	Buttons int    `json:"buttons,omitempty"` // number of buttons, -1 -> unknown
	Wheels  int    `json:"wheels,omitempty"`  // dto, wheels
}

func NewDriverInfoMouse(info C.driver_info_mouse_t) DriverInfoMouse {
	return DriverInfoMouse{
		Type:     DriverInfoTypeMouse,
		DBEntry0: ReadStringList(info.hddb0),
		DBEntry1: ReadStringList(info.hddb1),
		XF86:     C.GoString(info.xf86),
		GPM:      C.GoString(info.gpm),
		Buttons:  int(info.buttons),
		Wheels:   int(info.wheels),
	}
}

func (d DriverInfoMouse) DriverInfoType() DriverInfoType {
	return DriverInfoTypeMouse
}
