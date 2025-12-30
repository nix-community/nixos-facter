package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_fc_get_wwpn_ok(res_fc_t res) { return res.wwpn_ok; }
bool hd_res_fc_get_fcp_lun_ok(res_fc_t res) { return res.fcp_lun_ok; }
bool hd_res_fc_get_port_id_ok(res_fc_t res) { return res.port_id_ok; }

// CGO cannot access union type fields, so we do this as a workaround
res_fc_t hd_res_get_fc(hd_res_t *res) { return res->fc; }
*/
import "C"

import (
	"errors"
	"fmt"
)

// todo what is FC?
type ResourceFc struct {
	Type         ResourceType `json:"type"`
	WwpnOk       bool         `json:"wwpn_ok"`
	FcpLunOk     bool         `json:"fcp_lun_ok"`
	PortIDOk     bool         `json:"port_id_ok"`
	Wwpn         uint64       `json:"wwpn"`
	FcpLun       uint64       `json:"fcp_lun"`
	PortID       uint32       `json:"port_id"`
	ControllerID byte         `json:"controller_id"`
}

func (r ResourceFc) ResourceType() ResourceType {
	return ResourceTypeFc
}

func NewResourceFc(res *C.hd_res_t, resType ResourceType) (*ResourceFc, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeFc {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeFc, resType)
	}

	fc := C.hd_res_get_fc(res)
	var controllerID byte
	if fc.controller_id != nil {
		controllerID = byte(*fc.controller_id)
	}

	return &ResourceFc{
		WwpnOk:       bool(C.hd_res_fc_get_wwpn_ok(fc)),
		FcpLunOk:     bool(C.hd_res_fc_get_fcp_lun_ok(fc)),
		PortIDOk:     bool(C.hd_res_fc_get_port_id_ok(fc)),
		Wwpn:         uint64(fc.wwpn),
		FcpLun:       uint64(fc.fcp_lun),
		PortID:       uint32(fc.port_id),
		ControllerID: controllerID,
	}, nil
}
