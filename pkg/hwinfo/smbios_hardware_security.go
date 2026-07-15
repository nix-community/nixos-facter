package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosHardwareSecurity captures hardware security information.
type SmbiosHardwareSecurity struct {
	Type     SmbiosType `json:"-"`
	Handle   int        `json:"handle"`
	Power    *ID        `json:"power"`    // power-on password status
	Keyboard *ID        `json:"keyboard"` // keyboard password status
	Admin    *ID        `json:"admin"`    // admin password status
	Reset    *ID        `json:"reset"`    // front panel reset status
}

func NewSmbiosSecure(info C.smbios_secure_t) (*SmbiosHardwareSecurity, error) {
	return &SmbiosHardwareSecurity{
		Type:     SmbiosTypeHardwareSecurity,
		Handle:   int(info.handle),
		Power:    NewID(info.power),
		Keyboard: NewID(info.keyboard),
		Admin:    NewID(info.admin),
		Reset:    NewID(info.reset),
	}, nil
}

func (s SmbiosHardwareSecurity) SmbiosType() SmbiosType {
	return s.Type
}
