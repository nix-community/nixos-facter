package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_disk_geo_t hd_res_get_disk_geo(hd_res_t *res) { return res->disk_geo; }
*/
import "C"

import (
	"errors"
	"fmt"
)

//go:generate enumer -type=GeoType -json -transform=snake -trimprefix GeoType -output=./resource_enum_geo_type.go
type GeoType uint //nolint:recvcheck

const (
	GeoTypePhysical GeoType = iota
	GeoTypeLogical
	GeoTypeBiosEdd
	GeoTypeBiosLegacy
)

type ResourceDiskGeo struct {
	Type      ResourceType `json:"type"`
	Cylinders uint32       `json:"cylinders"`
	Heads     uint8        `json:"heads"`
	Sectors   uint32       `json:"sectors"`
	Size      string       `json:"size"`
	GeoType   GeoType      `json:"geo_type"`
}

func (r ResourceDiskGeo) ResourceType() ResourceType {
	return ResourceTypeDiskGeo
}

func NewResourceDiskGeo(res *C.hd_res_t, resType ResourceType) (*ResourceDiskGeo, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeDiskGeo {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeDiskGeo, resType)
	}

	disk := C.hd_res_get_disk_geo(res)

	return &ResourceDiskGeo{
		Type:      resType,
		Cylinders: uint32(disk.cyls),
		Heads:     uint8(disk.heads),
		Sectors:   uint32(disk.sectors),
		Size:      fmt.Sprintf("0x%x", uint64(disk.size)),
		GeoType:   GeoType(disk.geotype),
	}, nil
}
