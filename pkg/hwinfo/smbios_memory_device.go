package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryDevice captures system slot information.
//
//nolint:lll
type SmbiosMemoryDevice struct {
	Type              SmbiosType `json:"-"`
	Handle            int        `json:"handle"`
	Location          string     `json:"location"`      // device location
	BankLocation      string     `json:"bank_location"` // bank location
	Manufacturer      string     `json:"manufacturer"`
	Serial            string     `json:"-"` // omit from json
	AssetTag          string     `json:"-"`
	PartNumber        string     `json:"part_number"`
	ArrayHandle       int        `json:"array_handle"` // memory array this device belongs to
	ErrorHandle       int        `json:"error_handle"` // points to error info record; 0xfffe: not supported, 0xffff: no error
	Width             uint16     `json:"width"`        // data width in bits
	ECCBits           uint8      `json:"ecc_bits"`     // ecc bits
	Size              uint       `json:"size"`         // kB
	FormFactor        *ID        `json:"form_factor"`
	Set               uint8      `json:"set"` // 0: does not belong to a set; 1-0xfe: set number; 0xff: unknown
	MemoryType        *ID        `json:"memory_type"`
	MemoryTypeDetails []string   `json:"memory_type_details"`
	Speed             uint32     `json:"speed"` // MHz
}

func (s SmbiosMemoryDevice) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemDevice(info C.smbios_memdevice_t) (*SmbiosMemoryDevice, error) {
	return &SmbiosMemoryDevice{
		Type:              SmbiosTypeMemoryDevice,
		Handle:            int(info.handle),
		Location:          C.GoString(info.location),
		BankLocation:      C.GoString(info.bank),
		Manufacturer:      C.GoString(info.manuf),
		Serial:            C.GoString(info.serial),
		AssetTag:          C.GoString(info.asset),
		PartNumber:        C.GoString(info.part),
		ArrayHandle:       int(info.array_handle),
		ErrorHandle:       int(info.error_handle),
		Width:             uint16(info.width),
		ECCBits:           uint8(info.eccbits),
		Size:              uint(info.size),
		FormFactor:        NewID(info.form),
		Set:               uint8(info.set),
		MemoryType:        NewID(info.mem_type),
		MemoryTypeDetails: ReadStringList(info.type_detail.str),
		Speed:             uint32(info.speed),
	}, nil
}
