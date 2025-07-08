package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_detail_type_t hd_detail_get_type(hd_detail_t *det) { return det->type; }
hd_detail_pci_t hd_detail_get_pci(hd_detail_t *det) { return det->pci; }
hd_detail_usb_t hd_detail_get_usb(hd_detail_t *det) { return det->usb; }
hd_detail_isapnp_t hd_detail_get_isapnp(hd_detail_t *det) { return det->isapnp; }
hd_detail_cpu_t hd_detail_get_cpu(hd_detail_t *det) { return det->cpu; }
hd_detail_monitor_t hd_detail_get_monitor(hd_detail_t *det) { return det->monitor; }
hd_detail_bios_t hd_detail_get_bios(hd_detail_t *det) { return det->bios; }
hd_detail_sys_t hd_detail_get_sys(hd_detail_t *det) { return det->sys; }

*/
import "C"

import (
	"encoding/hex"
	"errors"
	"fmt"
	"unsafe"
)

//go:generate enumer -type=DetailType -json -transform=snake -trimprefix DetailType -output=./detail_enum_type.go
type DetailType uint //nolint:recvcheck

const (
	DetailTypePci DetailType = iota
	DetailTypeUsb
	DetailTypeIsaPnp
	DetailTypeCdrom

	DetailTypeFloppy
	DetailTypeBios
	DetailTypeCpu
	DetailTypeProm

	DetailTypeMonitor
	DetailTypeSys
	DetailTypeScsi
	DetailTypeDevtree

	DetailTypeCcw
	DetailTypeJoystick
)

type Detail interface {
	DetailType() DetailType
}

//nolint:ireturn
func NewDetail(detail *C.hd_detail_t) (Detail, error) {
	if detail == nil {
		return nil, errors.New("detail is nil")
	}

	var (
		err    error
		result Detail
	)

	switch DetailType(C.hd_detail_get_type(detail)) {
	case DetailTypePci:
		result, err = NewDetailPci(C.hd_detail_get_pci(detail))
	case DetailTypeUsb:
		result, err = NewDetailUsb(C.hd_detail_get_usb(detail))
	case DetailTypeIsaPnp:
		result, err = NewDetailIsaPnpDevice(C.hd_detail_get_isapnp(detail))
	case DetailTypeCpu:
		result, err = NewDetailCPU(C.hd_detail_get_cpu(detail))
	case DetailTypeMonitor:
		result, err = NewDetailMonitor(C.hd_detail_get_monitor(detail))
	case DetailTypeBios:
		result, err = NewDetailBios(C.hd_detail_get_bios(detail))
	case DetailTypeSys:
		result, err = NewDetailSys(C.hd_detail_get_sys(detail))
	case DetailTypeCdrom, DetailTypeFloppy, DetailTypeProm, DetailTypeScsi, DetailTypeDevtree, DetailTypeCcw,
		DetailTypeJoystick:
		// do nothing for now

	default:
		err = fmt.Errorf("unknown detail type %d", DetailType(C.hd_detail_get_type(detail)))
	}

	return result, err
}

type MemoryRange struct {
	Start string `json:"start"`
	Size  string `json:"size"`
	Data  string `json:"-"` // hex encoded
}

func NewMemoryRange(mem C.memory_range_t) MemoryRange {
	return MemoryRange{
		Start: fmt.Sprintf("0x%x", uint(mem.start)),
		Size:  fmt.Sprintf("0x%x", uint(mem.size)),
		Data:  hex.EncodeToString(C.GoBytes(unsafe.Pointer(&mem.data), C.int(mem.size))), //nolint:gocritic
	}
}
