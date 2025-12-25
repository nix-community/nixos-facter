package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type SyncRange struct {
	Min uint16 `json:"min"`
	Max uint16 `json:"max"`
}

type SyncTimings struct {
	Disp      uint16 `json:"disp"` // todo what's the proper name for this?
	SyncStart uint16 `json:"sync_start"`
	SyncEnd   uint16 `json:"sync_end"`
	Total     uint32 `json:"total"` // todo what's a better name for this?
}

type DetailMonitor struct {
	Type                  DetailType  `json:"-"`
	ManufactureYear       uint16      `json:"manufacture_year"`
	ManufactureWeek       uint8       `json:"manufacture_week"`
	VerticalSync          SyncRange   `json:"vertical_sync"`
	HorizontalSync        SyncRange   `json:"horizontal_sync"`
	HorizontalSyncTimings SyncTimings `json:"horizontal_sync_timings"`
	VerticalSyncTimings   SyncTimings `json:"vertical_sync_timings"`
	Clock                 uint32      `json:"clock"`
	Width                 uint32      `json:"width"`
	Height                uint32      `json:"height"`
	WidthMillimetres      uint32      `json:"width_millimetres"`
	HeightMillimetres     uint32      `json:"height_millimetres"`
	HorizontalFlag        byte        `json:"horizontal_flag"`
	VerticalFlag          byte        `json:"vertical_flag"`
	Vendor                string      `json:"vendor"`
	Name                  string      `json:"name"`

	Serial string `json:"-"`
}

func (d DetailMonitor) DetailType() DetailType {
	return DetailTypeMonitor
}

func NewDetailMonitor(mon C.hd_detail_monitor_t) (*DetailMonitor, error) {
	data := mon.data
	if data == nil {
		// Not an error: hwinfo can return detail structures with NULL data pointers.
		return nil, nil
	}

	return &DetailMonitor{
		Type:            DetailTypeMonitor,
		ManufactureYear: uint16(data.manu_year),
		ManufactureWeek: uint8(data.manu_week),
		VerticalSync: SyncRange{
			Min: uint16(data.min_vsync),
			Max: uint16(data.max_vsync),
		},
		HorizontalSync: SyncRange{
			Min: uint16(data.min_hsync),
			Max: uint16(data.max_hsync),
		},
		Clock:             uint32(data.clock),
		Width:             uint32(data.width),
		Height:            uint32(data.height),
		WidthMillimetres:  uint32(data.width_mm),
		HeightMillimetres: uint32(data.height_mm),
		HorizontalSyncTimings: SyncTimings{
			Disp:      uint16(data.hdisp),
			SyncStart: uint16(data.hsyncstart),
			SyncEnd:   uint16(data.hsyncend),
			Total:     uint32(data.htotal),
		},
		VerticalSyncTimings: SyncTimings{
			Disp:      uint16(data.vdisp),
			SyncStart: uint16(data.vsyncstart),
			SyncEnd:   uint16(data.vsyncend),
			Total:     uint32(data.vtotal),
		},
		HorizontalFlag: byte(data.hflag),
		VerticalFlag:   byte(data.vflag),
		Vendor:         C.GoString(data.vendor),
		Name:           C.GoString(data.name),
		Serial:         C.GoString(data.serial),
	}, nil
}
