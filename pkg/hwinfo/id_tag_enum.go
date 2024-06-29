// Code generated by "enumer -type=IdTag -json -transform=snake -trimprefix IdTag -output=./id_tag_enum.go"; DO NOT EDIT.

package hwinfo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _IdTagName = "pcieisausbspecialpcmciasdio"

var _IdTagIndex = [...]uint8{0, 3, 7, 10, 17, 23, 27}

const _IdTagLowerName = "pcieisausbspecialpcmciasdio"

func (i IdTag) String() string {
	i -= 1
	if i >= IdTag(len(_IdTagIndex)-1) {
		return fmt.Sprintf("IdTag(%d)", i+1)
	}
	return _IdTagName[_IdTagIndex[i]:_IdTagIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _IdTagNoOp() {
	var x [1]struct{}
	_ = x[IdTagPci-(1)]
	_ = x[IdTagEisa-(2)]
	_ = x[IdTagUsb-(3)]
	_ = x[IdTagSpecial-(4)]
	_ = x[IdTagPcmcia-(5)]
	_ = x[IdTagSdio-(6)]
}

var _IdTagValues = []IdTag{IdTagPci, IdTagEisa, IdTagUsb, IdTagSpecial, IdTagPcmcia, IdTagSdio}

var _IdTagNameToValueMap = map[string]IdTag{
	_IdTagName[0:3]:        IdTagPci,
	_IdTagLowerName[0:3]:   IdTagPci,
	_IdTagName[3:7]:        IdTagEisa,
	_IdTagLowerName[3:7]:   IdTagEisa,
	_IdTagName[7:10]:       IdTagUsb,
	_IdTagLowerName[7:10]:  IdTagUsb,
	_IdTagName[10:17]:      IdTagSpecial,
	_IdTagLowerName[10:17]: IdTagSpecial,
	_IdTagName[17:23]:      IdTagPcmcia,
	_IdTagLowerName[17:23]: IdTagPcmcia,
	_IdTagName[23:27]:      IdTagSdio,
	_IdTagLowerName[23:27]: IdTagSdio,
}

var _IdTagNames = []string{
	_IdTagName[0:3],
	_IdTagName[3:7],
	_IdTagName[7:10],
	_IdTagName[10:17],
	_IdTagName[17:23],
	_IdTagName[23:27],
}

// IdTagString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func IdTagString(s string) (IdTag, error) {
	if val, ok := _IdTagNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _IdTagNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to IdTag values", s)
}

// IdTagValues returns all values of the enum
func IdTagValues() []IdTag {
	return _IdTagValues
}

// IdTagStrings returns a slice of all String values of the enum
func IdTagStrings() []string {
	strs := make([]string, len(_IdTagNames))
	copy(strs, _IdTagNames)
	return strs
}

// IsAIdTag returns "true" if the value is listed in the enum definition. "false" otherwise
func (i IdTag) IsAIdTag() bool {
	for _, v := range _IdTagValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for IdTag
func (i IdTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for IdTag
func (i *IdTag) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IdTag should be a string, got %s", data)
	}

	var err error
	*i, err = IdTagString(s)
	return err
}
