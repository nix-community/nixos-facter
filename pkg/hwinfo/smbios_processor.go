package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosProcessor captures processor information.
type SmbiosProcessor struct {
	Type            SmbiosType `json:"-"`
	Handle          int        `json:"handle"`
	Socket          string     `json:"socket"`
	SocketType      *ID        `json:"socket_type"`
	SocketPopulated bool       `json:"socket_populated"` // true: populated, false: empty
	Manufacturer    string     `json:"manufacturer"`
	Version         string     `json:"version"`
	Serial          string     `json:"-"`    // omit from json output
	AssetTag        string     `json:"-"`    // asset tag
	Part            string     `json:"part"` // part number
	ProcessorType   *ID        `json:"processor_type"`
	ProcessorFamily *ID        `json:"processor_family"`
	ProcessorID     uint64     `json:"-"` // omit from json
	ProcessorStatus *ID        `json:"processor_status"`
	Voltage         uint       `json:"-"`
	ClockExt        uint32     `json:"clock_ext"`       // MHz
	ClockMax        uint32     `json:"clock_max"`       // MHz
	ClockCurrent    uint       `json:"-"`               // MHz
	CacheHandleL1   int        `json:"cache_handle_l1"` // handle of L1 cache
	CacheHandleL2   int        `json:"cache_handle_l2"` // handle of L2 cache
	CacheHandleL3   int        `json:"cache_handle_l3"` // handle of L3 cache
}

func (s SmbiosProcessor) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosProcessor(info C.smbios_processor_t) (*SmbiosProcessor, error) {
	return &SmbiosProcessor{
		Type:            SmbiosTypeProcessor,
		Handle:          int(info.handle),
		Socket:          C.GoString(info.socket),
		SocketType:      NewID(info.upgrade),
		SocketPopulated: uint(info.sock_status) == 1,
		Manufacturer:    C.GoString(info.manuf),
		Version:         C.GoString(info.version),
		Serial:          C.GoString(info.serial),
		AssetTag:        C.GoString(info.asset),
		Part:            C.GoString(info.part),
		ProcessorType:   NewID(info.pr_type),
		ProcessorFamily: NewID(info.family),
		ProcessorID:     uint64(info.cpu_id),
		ProcessorStatus: NewID(info.cpu_status),
		Voltage:         uint(info.voltage),
		ClockExt:        uint32(info.ext_clock),
		ClockMax:        uint32(info.max_speed),
		ClockCurrent:    uint(info.current_speed),
		CacheHandleL1:   int(info.l1_cache),
		CacheHandleL2:   int(info.l2_cache),
		CacheHandleL3:   int(info.l3_cache),
	}, nil
}
