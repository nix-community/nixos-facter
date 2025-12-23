package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

bool bios_info_is_apm_supported(bios_info_t *info) { return info->apm_supported; }
bool bios_info_is_apm_enabled(bios_info_t *info) { return info->apm_enabled; }
bool bios_info_is_pnp_bios(bios_info_t *info) { return info->is_pnp_bios; }
bool bios_info_has_lba_support(bios_info_t *info) { return info->lba_support; }
*/
import "C"

type ApmInfo struct {
	Supported  bool   `json:"supported"`
	Enabled    bool   `json:"enabled"`
	Version    uint8  `json:"version"`
	SubVersion uint8  `json:"sub_version"`
	BiosFlags  uint32 `json:"bios_flags"`
}

type VbeInfo struct {
	Version     uint16 `json:"version"`
	VideoMemory uint32 `json:"video_memory"`
}

type DetailBios struct {
	Type    DetailType `json:"-"`
	ApmInfo ApmInfo    `json:"apm_info"`
	VbeInfo VbeInfo    `json:"vbe_info"`

	// todo par and ser ports
	PnP           bool   `json:"pnp"`
	PnPId         uint   `json:"pnp_id"` // it is still in big endian format
	LbaSupport    bool   `json:"lba_support"`
	LowMemorySize uint32 `json:"low_memory_size"`
	// todo smp info
	// todo vbe info

	SmbiosVersion uint32 `json:"smbios_version"`

	// todo lcd
	// todo mouse
	// todo led
	// todo bios32
}

func (d DetailBios) DetailType() DetailType {
	return DetailTypeBios
}

func NewDetailBios(dev C.hd_detail_bios_t) (*DetailBios, error) {
	data := dev.data
	if data == nil {
		return nil, nil
	}

	return &DetailBios{
		Type: DetailTypeBios,
		ApmInfo: ApmInfo{
			Supported:  bool(C.bios_info_is_apm_supported(data)),
			Enabled:    bool(C.bios_info_is_apm_enabled(data)),
			Version:    uint8(data.apm_ver),
			SubVersion: uint8(data.apm_subver),
			BiosFlags:  uint32(data.apm_bios_flags),
		},
		VbeInfo: VbeInfo{
			Version:     uint16(data.vbe_ver),
			VideoMemory: uint32(data.vbe_video_mem),
		},
		PnP:           bool(C.bios_info_is_pnp_bios(data)),
		PnPId:         uint(data.pnp_id),
		LbaSupport:    bool(C.bios_info_has_lba_support(data)),
		LowMemorySize: uint32(data.low_mem_size),
		SmbiosVersion: uint32(data.smbios_ver),
	}, nil
}
