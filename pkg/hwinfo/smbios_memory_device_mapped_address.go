package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "fmt"

// SmbiosMemoryDeviceMappedAddress captures physical memory array information (consists of several memory devices).
//
//nolint:lll
type SmbiosMemoryDeviceMappedAddress struct {
	Type               SmbiosType `json:"-"`
	Handle             int        `json:"handle"`
	MemoryDeviceHandle int        `json:"memory_device_handle"`
	ArrayMapHandle     int        `json:"array_map_handle"`
	StartAddress       string     `json:"start_address"`
	EndAddress         string     `json:"end_address"`
	RowPosition        uint16     `json:"row_position"`        // position of the referenced memory device in a row of the address partition
	InterleavePosition uint16     `json:"interleave_position"` // dto, in an interleave
	InterleaveDepth    uint16     `json:"interleave_depth"`    // number of consecutive rows
}

func (s SmbiosMemoryDeviceMappedAddress) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemDeviceMap(info C.smbios_memdevicemap_t) (*SmbiosMemoryDeviceMappedAddress, error) {
	return &SmbiosMemoryDeviceMappedAddress{
		Type:               SmbiosTypeMemoryDeviceMappedAddress,
		Handle:             int(info.handle),
		MemoryDeviceHandle: int(info.memdevice_handle),
		ArrayMapHandle:     int(info.arraymap_handle),
		StartAddress:       fmt.Sprintf("0x%x", uint64(info.start_addr)),
		EndAddress:         fmt.Sprintf("0x%x", uint64(info.end_addr)),
		RowPosition:        uint16(info.row_pos),
		InterleavePosition: uint16(info.interleave_pos),
		InterleaveDepth:    uint16(info.interleave_depth),
	}, nil
}
