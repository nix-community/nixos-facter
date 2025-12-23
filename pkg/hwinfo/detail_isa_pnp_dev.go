package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_isapnp_card_get_broken(isapnp_card_t *card) { return card->broken; }
*/
import "C"

import (
	"encoding/hex"
	"errors"
	"unsafe"
)

type IsaPnpResource struct {
	Length int    `json:"length"`
	Type   int    `json:"type"`
	Data   string `json:"-"` // hex encoded
}

func NewIsaPnpResource(res *C.isapnp_res_t) *IsaPnpResource {
	if res == nil {
		return nil
	}

	return &IsaPnpResource{
		Length: int(res.len),
		Type:   int(res._type),
		Data:   hex.EncodeToString(C.GoBytes(unsafe.Pointer(&res.data), res.len)), //nolint:gocritic
	}
}

type IsaPnpCard struct {
	Csn      int             `json:"csn"`
	LogDevs  int             `json:"log_devs"` // todo full name?
	Serial   string          `json:"-"`
	CardRegs string          `json:"card_regs"` // todo full name?
	LdevRegs string          `json:"ldev_regs"` // todo full name? hex encoded
	ResLen   int             `json:"res_len"`   // todo full name?
	Broken   bool            `json:"broken"`    // mark a broken card
	Resource *IsaPnpResource `json:"resource"`
}

func NewIsaPnpCard(card *C.isapnp_card_t) (*IsaPnpCard, error) {
	if card == nil {
		return nil, errors.New("card is nil")
	}

	return &IsaPnpCard{
		Csn:     int(card.csn),
		LogDevs: int(card.log_devs),
		// Serial:   C.GoString(card.serial),	todo
		// CardRegs: C.GoString(card.card_regs), todo
		LdevRegs: hex.EncodeToString(C.GoBytes(unsafe.Pointer(&card.ldev_regs), C.int(0xd0))), //nolint:gocritic
		ResLen:   int(card.res_len),
		Broken:   bool(C.hd_isapnp_card_get_broken(card)),
		Resource: NewIsaPnpResource(card.res),
	}, nil
}

type DetailIsaPnpDevice struct {
	Type   DetailType  `json:"-"`
	Card   *IsaPnpCard `json:"card"`
	Device uint32      `json:"device"`
	Flags  uint32      `json:"flags"`
}

func (d DetailIsaPnpDevice) DetailType() DetailType {
	return DetailTypeIsaPnp
}

func NewDetailIsaPnpDevice(pnp C.hd_detail_isapnp_t) (*DetailIsaPnpDevice, error) {
	data := pnp.data
	if data == nil {
		// Not an error: hwinfo can return detail structures with NULL data pointers.
		return nil, nil
	}

	card, err := NewIsaPnpCard(data.card)
	if err != nil {
		return nil, err
	}

	return &DetailIsaPnpDevice{
		Type:   DetailTypeIsaPnp,
		Card:   card,
		Device: uint32(data.dev),
		Flags:  uint32(data.flags),
	}, nil
}
