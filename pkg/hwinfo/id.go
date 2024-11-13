package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"slices"
)

//go:generate enumer -type=IDTag -json -transform=snake -trimprefix IDTag -output=./id_tag_enum.go
type IDTag byte

const (
	IDTagPci IDTag = iota + 1
	IDTagEisa
	IDTagUsb
	IDTagSpecial
	IDTagPcmcia
	IDTagSdio
)

type ID struct {
	Type  IDTag
	Value uint16
	// Name (if any)
	Name string
}

type idJSON struct {
	Hex   string `json:"hex,omitempty"`
	Name  string `json:"name,omitempty"`
	Value uint16 `json:"value"`
}

func (i ID) MarshalJSON() ([]byte, error) {
	switch i.Type {

	case IDTagSpecial:
		return json.Marshal(i.Name)

	case 0, IDTagPci, IDTagEisa, IDTagUsb, IDTagPcmcia, IDTagSdio:

		return json.Marshal(idJSON{
			Hex:   fmt.Sprintf("%04x", i.Value),
			Name:  i.Name,
			Value: i.Value,
		})

	default:
		return nil, fmt.Errorf("unknown id type %d", i.Type)
	}
}

func (i ID) IsEmpty() bool {
	return i.Type == 0 && i.Value == 0 && (i.Name == "" || i.Name == "None")
}

func (i ID) String() string {
	return fmt.Sprintf("%d:%s", i.Value, i.Name)
}

func (i ID) Is(ids ...uint16) bool {
	return slices.Contains(ids, i.Value)
}

func NewID(id C.hd_id_t) *ID {
	result := ID{
		/*
			 	ID is actually a combination of some tag to differentiate the various id types and the real id.
				We do the same thing as the ID_VALUE macro in hd.h to get the true value.
		*/
		Type:  IDTag((id.id >> 16) & 0xf),
		Value: uint16(id.id),
		Name:  C.GoString(id.name),
	}
	if result.IsEmpty() {
		return nil
	}
	return &result
}
