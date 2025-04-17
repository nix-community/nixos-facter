package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoDisplay struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DBEntry0 []string `json:"db_entry_0,omitempty"`
	DBEntry1 []string `json:"db_entry_1,omitempty"`

	Width                 uint32      `json:"width"`
	Height                uint32      `json:"height"`
	VerticalSync          SyncRange   `json:"vertical_sync"`
	HorizontalSync        SyncRange   `json:"horizontal_sync"`
	Bandwidth             uint32      `json:"bandwidth"`
	HorizontalSyncTimings SyncTimings `json:"horizontal_sync_timings"`
	VerticalSyncTimings   SyncTimings `json:"vertical_sync_timings"`
	HorizontalFlag        byte        `json:"horizontal_flag"`
	VerticalFlag          byte        `json:"vertical_flag"`
}

func (d DriverInfoDisplay) DriverInfoType() DriverInfoType {
	return DriverInfoTypeDisplay
}

func NewDriverInfoDisplay(info C.driver_info_display_t) DriverInfoDisplay {
	return DriverInfoDisplay{
		Type:     DriverInfoTypeDisplay,
		DBEntry0: ReadStringList(info.hddb0),
		DBEntry1: ReadStringList(info.hddb1),
		Width:    uint32(info.width),
		Height:   uint32(info.height),
		VerticalSync: SyncRange{
			Min: uint16(info.min_vsync),
			Max: uint16(info.max_vsync),
		},
		HorizontalSync: SyncRange{
			Min: uint16(info.min_hsync),
			Max: uint16(info.max_hsync),
		},
		HorizontalSyncTimings: SyncTimings{
			Disp:      uint16(info.hdisp),
			SyncStart: uint16(info.hsyncstart),
			SyncEnd:   uint16(info.hsyncend),
			Total:     uint32(info.htotal),
		},
		VerticalSyncTimings: SyncTimings{
			Disp:      uint16(info.vdisp),
			SyncStart: uint16(info.vsyncstart),
			SyncEnd:   uint16(info.vsyncend),
			Total:     uint32(info.vtotal),
		},
		HorizontalFlag: byte(info.hflag),
		VerticalFlag:   byte(info.vflag),
	}
}
