package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

bool cpu_info_fpu(cpu_info_t *info) { return info->fpu; }
bool cpu_info_fpu_exception(cpu_info_t *info) { return info->fpu_exception; }
bool cpu_info_write_protect(cpu_info_t *info) { return info->write_protect; }
*/
import "C"

import (
	"fmt"
	"os"
	"regexp"
)

//go:generate enumer -type=CPUArch -json -transform=snake -trimprefix CPUArch -output=./detail_enum_cpu_arch.go
type CPUArch uint //nolint:recvcheck

const (
	CPUArchUnknown CPUArch = iota
	CPUArchIntel
	CPUArchAlpha
	CPUArchSparc
	CPUArchSparc64
	CPUArchPpc
	CPUArchPpc64
	CpiArch68k
	CPUArchIa64
	CPUArchS390
	CPUArchS390x
	CPUArchArm
	CPUArchMips
	CPUArchX86_64
	CPUArchAarch64
	CPUArchLoongarch
	CPUArchRiscv
)

type AddressSizes struct {
	Physical string `json:"physical,omitempty"`
	Virtual  string `json:"virtual,omitempty"`
}

type DetailCPU struct {
	Type DetailType `json:"-"`

	Architecture CPUArch `json:"architecture"`

	VendorName string `json:"vendor_name,omitempty"`
	ModelName  string `json:"model_name,omitempty"`

	Family   uint16 `json:"family"`
	Model    uint16 `json:"model"`
	Stepping uint32 `json:"stepping"`

	Platform string `json:"platform,omitempty"`

	Features        []string `json:"features,omitempty"`
	Bugs            []string `json:"bugs,omitempty"`
	PowerManagement []string `json:"power_management,omitempty"`

	Bogo     uint32 `json:"bogo"`
	Cache    uint32 `json:"cache,omitempty"`
	Units    uint32 `json:"units,omitempty"`
	Clock    uint   `json:"-"`
	PageSize uint32 `json:"page_size"`

	// x86 only fields
	PhysicalID     uint16       `json:"physical_id"`
	Siblings       uint16       `json:"siblings,omitempty"`
	Cores          uint16       `json:"cores,omitempty"`
	CoreID         uint16       `json:"-"`
	Fpu            bool         `json:"fpu"`
	FpuException   bool         `json:"fpu_exception"`
	CpuidLevel     uint8        `json:"cpuid_level,omitempty"`
	WriteProtect   bool         `json:"write_protect"`
	TlbSize        uint16       `json:"tlb_size,omitempty"`
	ClflushSize    uint16       `json:"clflush_size,omitempty"`
	CacheAlignment int          `json:"cache_alignment,omitempty"`
	AddressSizes   AddressSizes `json:"address_sizes,omitempty"`
	Apicid         uint         `json:"-"`
	ApicidInitial  uint         `json:"-"`
}

var matchCPUFreq = regexp.MustCompile(`, \d+ MHz$`)

func stripCPUFreq(s string) string {
	// strip frequency of the model name as it is not stable.
	return matchCPUFreq.ReplaceAllString(s, "")
}

func NewDetailCPU(cpu C.hd_detail_cpu_t) (*DetailCPU, error) {
	data := cpu.data
	if data == nil {
		// Not an error: hwinfo can return detail structures with NULL data pointers
		// on certain systems (e.g., older VMware ESXi). See hdp.c:1082 for hwinfo's
		// own handling: if(!(ct = hd->detail->cpu.data)) return;
		return nil, nil
	}

	return &DetailCPU{
		Type: DetailTypeCpu,

		Architecture: CPUArch(data.architecture),
		VendorName:   C.GoString(data.vend_name),
		ModelName:    stripCPUFreq(C.GoString(data.model_name)),

		Family:   uint16(data.family),
		Model:    uint16(data.model),
		Stepping: uint32(data.stepping),

		Platform: C.GoString(data.platform),

		Features:        ReadStringList(data.features),
		Bugs:            ReadStringList(data.bugs),
		PowerManagement: ReadStringList(data.power_management),

		Clock: uint(data.clock),

		// Bogo is reported as a float, but that value isn't stable so we truncate to an int.
		Bogo:     uint32(data.bogo),
		Cache:    uint32(data.cache),
		Units:    uint32(data.units),
		PageSize: uint32(os.Getpagesize()),

		PhysicalID:     uint16(data.physical_id),
		Siblings:       uint16(data.siblings),
		Cores:          uint16(data.cores),
		CoreID:         uint16(data.core_id),
		Apicid:         uint(data.apicid),
		ApicidInitial:  uint(data.apicid_initial),
		Fpu:            bool(C.cpu_info_fpu(data)),
		FpuException:   bool(C.cpu_info_fpu_exception(data)),
		CpuidLevel:     uint8(data.cpuid_level),
		WriteProtect:   bool(C.cpu_info_write_protect(data)),
		TlbSize:        uint16(data.tlb_size),
		ClflushSize:    uint16(data.clflush_size),
		CacheAlignment: int(data.cache_alignment),
		AddressSizes: AddressSizes{
			Physical: fmt.Sprintf("0x%x", uint(data.address_size_physical)),
			Virtual:  fmt.Sprintf("0x%x", uint(data.address_size_virtual)),
		},
	}, nil
}

func (d DetailCPU) DetailType() DetailType {
	return DetailTypeCpu
}
