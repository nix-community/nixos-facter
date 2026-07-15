package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "unsafe"

// SmbiosGroupAssociations captures group associations.
type SmbiosGroupAssociations struct {
	Type    SmbiosType `json:"-"`
	Handle  int        `json:"handle"`
	Name    string     `json:"name"`              // group name
	Handles []int      `json:"handles,omitempty"` // array of item handles
}

func NewSmbiosGroup(info C.smbios_group_t) (*SmbiosGroupAssociations, error) {
	return &SmbiosGroupAssociations{
		Type:    SmbiosTypeGroupAssociations,
		Handle:  int(info.handle),
		Name:    C.GoString(info.name),
		Handles: ReadIntArray(unsafe.Pointer(info.item_handles), int(info.items_len)),
	}, nil
}

func (s SmbiosGroupAssociations) SmbiosType() SmbiosType {
	return s.Type
}
