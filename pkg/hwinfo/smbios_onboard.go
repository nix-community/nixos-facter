package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

char* smbios_onboard_get_name(smbios_onboard_t sm, int i) { return sm.dev[i].name; }
hd_id_t smbios_onboard_get_type(smbios_onboard_t sm, int i) { return sm.dev[i].type; }
unsigned smbios_onboard_get_status(smbios_onboard_t sm, int i) { return sm.dev[i].status; }
*/
import "C"

type OnboardDevice struct {
	Name    string `json:"name"`
	Type    *ID    `json:"type"`
	Enabled bool   `json:"enabled"`
}

// SmbiosOnboard captures overall system related information.
type SmbiosOnboard struct {
	Type    SmbiosType      `json:"-"`
	Handle  int             `json:"handle"`
	Devices []OnboardDevice `json:"devices,omitempty"`
}

func NewSmbiosOnboard(info C.smbios_onboard_t) (*SmbiosOnboard, error) {
	count := int(info.dev_len)
	devices := make([]OnboardDevice, 0, count)

	for i := range count {
		devices = append(devices, OnboardDevice{
			Name:    C.GoString(C.smbios_onboard_get_name(info, C.int(i))),
			Type:    NewID(C.smbios_onboard_get_type(info, C.int(i))),
			Enabled: uint(C.smbios_onboard_get_status(info, C.int(i))) == 1,
		})
	}

	return &SmbiosOnboard{
		Type:    SmbiosTypeOnboard,
		Handle:  int(info.handle),
		Devices: devices,
	}, nil
}

func (s SmbiosOnboard) SmbiosType() SmbiosType {
	return s.Type
}
