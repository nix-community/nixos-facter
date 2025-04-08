package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "fmt"

// SmbiosMemoryArrayMappedAddress captures physical memory array information (consists of several memory devices).
type SmbiosMemoryArrayMappedAddress struct {
	Type         SmbiosType `json:"-"`
	Handle       int        `json:"handle"`
	ArrayHandle  int        `json:"array_handle"`  // memory array this mapping belongs to
	StartAddress string     `json:"start_address"` // memory range start address
	EndAddress   string     `json:"end_address"`   // end address
	PartWidth    uint16     `json:"part_width"`    // number of memory devices
}

func (s SmbiosMemoryArrayMappedAddress) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemArrayMap(info C.smbios_memarraymap_t) (*SmbiosMemoryArrayMappedAddress, error) {
	return &SmbiosMemoryArrayMappedAddress{
		Type:         SmbiosTypeMemoryArrayMappedAddress,
		Handle:       int(info.handle),
		ArrayHandle:  int(info.array_handle),
		StartAddress: fmt.Sprintf("0x%x", uint64(info.start_addr)),
		EndAddress:   fmt.Sprintf("0x%x", uint64(info.end_addr)),
		PartWidth:    uint16(info.part_width),
	}, nil
}
