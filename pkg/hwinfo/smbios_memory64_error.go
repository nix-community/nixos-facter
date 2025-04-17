package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "fmt"

// SmbiosMemory64Error captures 32-bit memory error information.
type SmbiosMemory64Error struct {
	Type          SmbiosType `json:"-"`
	Handle        int        `json:"handle"`
	ErrorType     *ID        `json:"error_type"`     // error type memory
	Granularity   *ID        `json:"granularity"`    // memory array or memory partition
	Operation     *ID        `json:"operation"`      // mem operation causing the rror
	Syndrome      uint32     `json:"syndrome"`       // vendor-specific ECC syndrome; 0: unknown
	ArrayAddress  string     `json:"array_address"`  // fault address relative to mem array; 0x80000000: unknown
	DeviceAddress string     `json:"device_address"` // fault address relative to mem array; 0x80000000: unknown
	Range         uint32     `json:"range"`          // range within which the error can be determined; 0x80000000: unknown
}

func (s SmbiosMemory64Error) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMem64Error(info C.smbios_mem64error_t) (*SmbiosMemory64Error, error) {
	return &SmbiosMemory64Error{
		Type:          SmbiosTypeMemory64Error,
		Handle:        int(info.handle),
		ErrorType:     NewID(info.err_type),
		Granularity:   NewID(info.granularity),
		Operation:     NewID(info.operation),
		Syndrome:      uint32(info.syndrome),
		ArrayAddress:  fmt.Sprintf("0x%x", uint(info.array_addr)),
		DeviceAddress: fmt.Sprintf("0x%x", uint(info.device_addr)),
		Range:         uint32(info._range),
	}, nil
}
