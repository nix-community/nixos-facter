package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"encoding/hex"
	"slices"
	"unsafe"
)

//go:generate enumer -type=PciFlag -json -transform=snake -trimprefix PciFlag -output=./detail_enum_pci_flag.go
type PciFlag uint //nolint:recvcheck

const (
	PciFlagOk PciFlag = iota
	PciFlagPm
	PciFlagAgp
)

func ParsePciFlags(flags uint) []PciFlag {
	var result []PciFlag

	for _, flag := range PciFlagValues() {
		if (flag & (1 << flags)) == 1 {
			result = append(result, flag)
		}
	}
	// ensure stable output
	slices.Sort(result)

	return result
}

type DetailPci struct {
	Type DetailType `json:"-"`

	Flags    []PciFlag `json:"flags,omitempty"` //
	Function uint32    `json:"function"`

	// todo map pci constants from pci.h?
	Command      uint32 `json:"command"`       // PCI_COMMAND
	HeaderType   uint32 `json:"header_type"`   // PCI_HEADER_TYPE
	SecondaryBus uint32 `json:"secondary_bus"` // > 0 for PCI & CB bridges

	Irq uint16 `json:"-"` // used irq if any
	// Programming Interface Byte: a read-only register that specifies a register-level programming interface for the
	// device.
	ProgIf uint16 `json:"prog_if"`

	// already included in the parent model, so we omit from JSON output
	Bus  uint `json:"-"`
	Slot uint `json:"-"`

	BaseClass uint `json:"-"`
	SubClass  uint `json:"-"`

	Device    uint `json:"-"`
	Vendor    uint `json:"-"`
	SubDevice uint `json:"-"`
	SubVendor uint `json:"-"`
	Revision  uint `json:"-"`

	BaseAddress  [7]uint64 `json:"-"` // I/O or memory base
	BaseLength   [7]uint64 `json:"-"` // I/O or memory ranges
	AddressFlags [7]uint   `json:"-"` // I/O or memory address flags

	RomBaseAddress uint64 `json:"-"` // memory base for card ROM
	RomBaseLength  uint64 `json:"-"` // memory range for card ROM

	SysfsID     string `json:"-"` // sysfs path
	SysfsBusID  string `json:"-"` // sysfs bus id
	ModuleAlias string `json:"-"` // module alias
	Label       string `json:"-"` // Consistent Device Name (CDN), pci firmware 3.1, chapter 4.6.7

	// Omit from JSON output
	Data          string `json:"-"` // the PCI data, hex encoded
	DataLength    uint   `json:"-"` // holds the actual length of Data
	DataExtLength uint   `json:"-"` // max. accessed config byte
	Log           string `json:"-"` // log messages
}

func (p DetailPci) DetailType() DetailType {
	return DetailTypePci
}

func NewDetailPci(pci C.hd_detail_pci_t) (*DetailPci, error) {
	data := pci.data

	return &DetailPci{
		Type:           DetailTypePci,
		Data:           hex.EncodeToString(C.GoBytes(unsafe.Pointer(&data.data), 256)), //nolint:gocritic
		DataLength:     uint(data.data_len),
		DataExtLength:  uint(data.data_ext_len),
		Log:            C.GoString(data.log),
		Flags:          ParsePciFlags(uint(data.flags)),
		Command:        uint32(data.cmd),
		HeaderType:     uint32(data.hdr_type),
		SecondaryBus:   uint32(data.secondary_bus),
		Bus:            uint(data.bus),
		Slot:           uint(data.slot),
		Function:       uint32(data._func),
		BaseClass:      uint(data.base_class),
		SubClass:       uint(data.sub_class),
		ProgIf:         uint16(data.prog_if),
		Device:         uint(data.dev),
		Vendor:         uint(data.vend),
		SubDevice:      uint(data.sub_dev),
		SubVendor:      uint(data.sub_vend),
		Revision:       uint(data.rev),
		Irq:            uint16(data.irq),
		BaseAddress:    [7]uint64(ReadUint64Array(unsafe.Pointer(&data.base_addr), 7)),
		BaseLength:     [7]uint64(ReadUint64Array(unsafe.Pointer(&data.base_len), 7)),
		AddressFlags:   [7]uint(ReadUintArray(unsafe.Pointer(&data.addr_flags), 7)),
		RomBaseAddress: uint64(data.rom_base_addr),
		RomBaseLength:  uint64(data.rom_base_len),
		SysfsID:        C.GoString(data.sysfs_id),
		SysfsBusID:     C.GoString(data.sysfs_bus_id),
		ModuleAlias:    C.GoString(data.modalias),
		Label:          C.GoString(data.label),
	}, nil
}
