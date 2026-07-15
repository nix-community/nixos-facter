package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoIsdn struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DBEntry0 []string `json:"db_entry_0,omitempty"`
	DBEntry1 []string `json:"db_entry_1,omitempty"`

	I4lType    int    `json:"i4l_type"`
	I4lSubtype int    `json:"i4l_sub_type"`
	I4lName    string `json:"i4l_name,omitempty"`
	// todo isdn params
}

func NewDriverInfoIsdn(info C.driver_info_isdn_t) DriverInfoIsdn {
	return DriverInfoIsdn{
		Type:       DriverInfoTypeIsdn,
		DBEntry0:   ReadStringList(info.hddb0),
		DBEntry1:   ReadStringList(info.hddb1),
		I4lType:    int(info.i4l_type),
		I4lSubtype: int(info.i4l_subtype),
		I4lName:    C.GoString(info.i4l_name),
	}
}

func (d DriverInfoIsdn) DriverInfoType() DriverInfoType {
	return DriverInfoTypeIsdn
}
