// Code generated by "enumer -type=CPUArch -json -transform=snake -trimprefix CPUArch -output=./detail_enum_cpu_arch.go"; DO NOT EDIT.

package hwinfo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _CPUArchName = "unknownintelalphasparcsparc64ppcppc64cpi_arch68kia64s390s390xarmmipsx86_64aarch64loongarchriscv"

var _CPUArchIndex = [...]uint8{0, 7, 12, 17, 22, 29, 32, 37, 48, 52, 56, 61, 64, 68, 74, 81, 90, 95}

const _CPUArchLowerName = "unknownintelalphasparcsparc64ppcppc64cpi_arch68kia64s390s390xarmmipsx86_64aarch64loongarchriscv"

func (i CPUArch) String() string {
	if i >= CPUArch(len(_CPUArchIndex)-1) {
		return fmt.Sprintf("CPUArch(%d)", i)
	}
	return _CPUArchName[_CPUArchIndex[i]:_CPUArchIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CPUArchNoOp() {
	var x [1]struct{}
	_ = x[CPUArchUnknown-(0)]
	_ = x[CPUArchIntel-(1)]
	_ = x[CPUArchAlpha-(2)]
	_ = x[CPUArchSparc-(3)]
	_ = x[CPUArchSparc64-(4)]
	_ = x[CPUArchPpc-(5)]
	_ = x[CPUArchPpc64-(6)]
	_ = x[CpiArch68k-(7)]
	_ = x[CPUArchIa64-(8)]
	_ = x[CPUArchS390-(9)]
	_ = x[CPUArchS390x-(10)]
	_ = x[CPUArchArm-(11)]
	_ = x[CPUArchMips-(12)]
	_ = x[CPUArchX86_64-(13)]
	_ = x[CPUArchAarch64-(14)]
	_ = x[CPUArchLoongarch-(15)]
	_ = x[CPUArchRiscv-(16)]
}

var _CPUArchValues = []CPUArch{CPUArchUnknown, CPUArchIntel, CPUArchAlpha, CPUArchSparc, CPUArchSparc64, CPUArchPpc, CPUArchPpc64, CpiArch68k, CPUArchIa64, CPUArchS390, CPUArchS390x, CPUArchArm, CPUArchMips, CPUArchX86_64, CPUArchAarch64, CPUArchLoongarch, CPUArchRiscv}

var _CPUArchNameToValueMap = map[string]CPUArch{
	_CPUArchName[0:7]:        CPUArchUnknown,
	_CPUArchLowerName[0:7]:   CPUArchUnknown,
	_CPUArchName[7:12]:       CPUArchIntel,
	_CPUArchLowerName[7:12]:  CPUArchIntel,
	_CPUArchName[12:17]:      CPUArchAlpha,
	_CPUArchLowerName[12:17]: CPUArchAlpha,
	_CPUArchName[17:22]:      CPUArchSparc,
	_CPUArchLowerName[17:22]: CPUArchSparc,
	_CPUArchName[22:29]:      CPUArchSparc64,
	_CPUArchLowerName[22:29]: CPUArchSparc64,
	_CPUArchName[29:32]:      CPUArchPpc,
	_CPUArchLowerName[29:32]: CPUArchPpc,
	_CPUArchName[32:37]:      CPUArchPpc64,
	_CPUArchLowerName[32:37]: CPUArchPpc64,
	_CPUArchName[37:48]:      CpiArch68k,
	_CPUArchLowerName[37:48]: CpiArch68k,
	_CPUArchName[48:52]:      CPUArchIa64,
	_CPUArchLowerName[48:52]: CPUArchIa64,
	_CPUArchName[52:56]:      CPUArchS390,
	_CPUArchLowerName[52:56]: CPUArchS390,
	_CPUArchName[56:61]:      CPUArchS390x,
	_CPUArchLowerName[56:61]: CPUArchS390x,
	_CPUArchName[61:64]:      CPUArchArm,
	_CPUArchLowerName[61:64]: CPUArchArm,
	_CPUArchName[64:68]:      CPUArchMips,
	_CPUArchLowerName[64:68]: CPUArchMips,
	_CPUArchName[68:74]:      CPUArchX86_64,
	_CPUArchLowerName[68:74]: CPUArchX86_64,
	_CPUArchName[74:81]:      CPUArchAarch64,
	_CPUArchLowerName[74:81]: CPUArchAarch64,
	_CPUArchName[81:90]:      CPUArchLoongarch,
	_CPUArchLowerName[81:90]: CPUArchLoongarch,
	_CPUArchName[90:95]:      CPUArchRiscv,
	_CPUArchLowerName[90:95]: CPUArchRiscv,
}

var _CPUArchNames = []string{
	_CPUArchName[0:7],
	_CPUArchName[7:12],
	_CPUArchName[12:17],
	_CPUArchName[17:22],
	_CPUArchName[22:29],
	_CPUArchName[29:32],
	_CPUArchName[32:37],
	_CPUArchName[37:48],
	_CPUArchName[48:52],
	_CPUArchName[52:56],
	_CPUArchName[56:61],
	_CPUArchName[61:64],
	_CPUArchName[64:68],
	_CPUArchName[68:74],
	_CPUArchName[74:81],
	_CPUArchName[81:90],
	_CPUArchName[90:95],
}

// CPUArchString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CPUArchString(s string) (CPUArch, error) {
	if val, ok := _CPUArchNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CPUArchNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CPUArch values", s)
}

// CPUArchValues returns all values of the enum
func CPUArchValues() []CPUArch {
	return _CPUArchValues
}

// CPUArchStrings returns a slice of all String values of the enum
func CPUArchStrings() []string {
	strs := make([]string, len(_CPUArchNames))
	copy(strs, _CPUArchNames)
	return strs
}

// IsACPUArch returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CPUArch) IsACPUArch() bool {
	for _, v := range _CPUArchValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CPUArch
func (i CPUArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CPUArch
func (i *CPUArch) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CPUArch should be a string, got %s", data)
	}

	var err error
	*i, err = CPUArchString(s)
	return err
}