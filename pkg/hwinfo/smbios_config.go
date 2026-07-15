package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosConfig captures system config information.
type SmbiosConfig struct {
	Type    SmbiosType `json:"-"`
	Handle  int        `json:"handle"`
	Options []string   `json:"options,omitempty"`
}

func NewSmbiosConfig(info C.smbios_config_t) (*SmbiosConfig, error) {
	return &SmbiosConfig{
		Type:    SmbiosTypeConfig,
		Handle:  int(info.handle),
		Options: ReadStringList(info.options),
	}, nil
}

func (s SmbiosConfig) SmbiosType() SmbiosType {
	return s.Type
}
