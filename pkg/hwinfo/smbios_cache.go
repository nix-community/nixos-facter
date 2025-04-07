package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosCache captures processor information.
type SmbiosCache struct {
	Type          SmbiosType `json:"-"`
	Handle        int        `json:"handle"`
	Socket        string     `json:"socket"`       // socket designation
	SizeMax       uint32     `json:"size_max"`     // max cache size in kbytes
	SizeCurrent   uint32     `json:"size_current"` // current size in kbytes
	Speed         uint32     `json:"speed"`        // cache speed in nanoseconds
	Mode          *ID        `json:"mode"`         // operational mode
	Enabled       bool       `json:"enabled"`
	Location      *ID        `json:"location"` // cache location
	Socketed      bool       `json:"socketed"`
	Level         uint       `json:"level"`         // cache level (0 = L1, 1 = L2, ...)
	ECC           *ID        `json:"ecc"`           // error correction type
	CacheType     *ID        `json:"cache_type"`    // logical cache type
	Associativity *ID        `json:"associativity"` // cache associativity
	SRAMType      []string   `json:"sram_type_current"`
	SRAMTypes     []string   `json:"sram_type_supported"`
}

func (s SmbiosCache) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosCache(info C.smbios_cache_t) (*SmbiosCache, error) {
	return &SmbiosCache{
		Type:          SmbiosTypeCache,
		Handle:        int(info.handle),
		Socket:        C.GoString(info.socket),
		SizeMax:       uint32(info.max_size),
		SizeCurrent:   uint32(info.current_size),
		Speed:         uint32(info.speed),
		Mode:          NewID(info.mode),
		Enabled:       uint(info.state) == 1,
		Location:      NewID(info.location),
		Socketed:      uint(info.socketed) == 1,
		Level:         uint(info.level),
		ECC:           NewID(info.ecc),
		CacheType:     NewID(info.cache_type),
		Associativity: NewID(info.assoc),
		SRAMType:      ReadStringList(info.sram.str),
		SRAMTypes:     ReadStringList(info.supp_sram.str),
	}, nil
}
