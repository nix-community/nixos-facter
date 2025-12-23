package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DetailSys struct {
	Type       DetailType `json:"-"`
	SystemType string     `json:"system_type,omitempty"`
	Generation string     `json:"generation,omitempty"`
	Vendor     string     `json:"vendor,omitempty"`
	Model      string     `json:"model,omitempty"`
	Serial     string     `json:"-"`
	Language   string     `json:"language,omitempty"`
	FormFactor string     `json:"form_factor,omitempty"`
}

func (d DetailSys) DetailType() DetailType {
	return DetailTypeSys
}

func NewDetailSys(sys C.hd_detail_sys_t) (*DetailSys, error) {
	data := sys.data
	if data == nil {
		return nil, nil
	}

	return &DetailSys{
		Type:       DetailTypeSys,
		SystemType: C.GoString(data.system_type),
		Generation: C.GoString(data.generation),
		Vendor:     C.GoString(data.vendor),
		Model:      C.GoString(data.model),
		Serial:     C.GoString(data.serial),
		Language:   C.GoString(data.lang),
		FormFactor: C.GoString(data.formfactor),
	}, nil
}
