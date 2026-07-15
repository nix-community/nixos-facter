package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosLanguage captures language information.
type SmbiosLanguage struct {
	Type            SmbiosType `json:"-"`
	Handle          int        `json:"handle"`
	Languages       []string   `json:"languages,omitempty"`
	CurrentLanguage string     `json:"-"`
}

func NewSmbiosLang(info C.smbios_lang_t) (*SmbiosLanguage, error) {
	return &SmbiosLanguage{
		Type:            SmbiosTypeLanguage,
		Handle:          int(info.handle),
		Languages:       ReadStringList(info.strings),
		CurrentLanguage: C.GoString(info.current),
	}, nil
}

func (s SmbiosLanguage) SmbiosType() SmbiosType {
	return s.Type
}
