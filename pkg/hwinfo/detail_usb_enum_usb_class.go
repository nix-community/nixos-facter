// Code generated by "enumer -type=UsbClass -json -transform=snake -trimprefix UsbClass -output=./detail_usb_enum_usb_class.go"; DO NOT EDIT.

package hwinfo

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	_UsbClassName_0      = "per_interfaceaudiocommhid"
	_UsbClassLowerName_0 = "per_interfaceaudiocommhid"
	_UsbClassName_1      = "physicalimageprintermass_storagehubdatasmart_card"
	_UsbClassLowerName_1 = "physicalimageprintermass_storagehubdatasmart_card"
	_UsbClassName_2      = "content_securityvideopersonal_healthcareaudio_videobillboardusb_type_c_bridge"
	_UsbClassLowerName_2 = "content_securityvideopersonal_healthcareaudio_videobillboardusb_type_c_bridge"
	_UsbClassName_3      = "diagnostic_device"
	_UsbClassLowerName_3 = "diagnostic_device"
	_UsbClassName_4      = "wireless"
	_UsbClassLowerName_4 = "wireless"
	_UsbClassName_5      = "miscellaneous"
	_UsbClassLowerName_5 = "miscellaneous"
	_UsbClassName_6      = "applicationvendor_spec"
	_UsbClassLowerName_6 = "applicationvendor_spec"
)

var (
	_UsbClassIndex_0 = [...]uint8{0, 13, 18, 22, 25}
	_UsbClassIndex_1 = [...]uint8{0, 8, 13, 20, 32, 35, 39, 49}
	_UsbClassIndex_2 = [...]uint8{0, 16, 21, 40, 51, 60, 77}
	_UsbClassIndex_3 = [...]uint8{0, 17}
	_UsbClassIndex_4 = [...]uint8{0, 8}
	_UsbClassIndex_5 = [...]uint8{0, 13}
	_UsbClassIndex_6 = [...]uint8{0, 11, 22}
)

func (i UsbClass) String() string {
	switch {
	case 0 <= i && i <= 3:
		return _UsbClassName_0[_UsbClassIndex_0[i]:_UsbClassIndex_0[i+1]]
	case 5 <= i && i <= 11:
		i -= 5
		return _UsbClassName_1[_UsbClassIndex_1[i]:_UsbClassIndex_1[i+1]]
	case 13 <= i && i <= 18:
		i -= 13
		return _UsbClassName_2[_UsbClassIndex_2[i]:_UsbClassIndex_2[i+1]]
	case i == 220:
		return _UsbClassName_3
	case i == 224:
		return _UsbClassName_4
	case i == 239:
		return _UsbClassName_5
	case 254 <= i && i <= 255:
		i -= 254
		return _UsbClassName_6[_UsbClassIndex_6[i]:_UsbClassIndex_6[i+1]]
	default:
		return fmt.Sprintf("UsbClass(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UsbClassNoOp() {
	var x [1]struct{}
	_ = x[UsbClassPerInterface-(0)]
	_ = x[UsbClassAudio-(1)]
	_ = x[UsbClassComm-(2)]
	_ = x[UsbClassHID-(3)]
	_ = x[UsbClassPhysical-(5)]
	_ = x[UsbClassImage-(6)]
	_ = x[UsbClassPrinter-(7)]
	_ = x[UsbClassMassStorage-(8)]
	_ = x[UsbClassHub-(9)]
	_ = x[UsbClassData-(10)]
	_ = x[UsbClassSmartCard-(11)]
	_ = x[UsbClassContentSecurity-(13)]
	_ = x[UsbClassVideo-(14)]
	_ = x[UsbClassPersonalHealthcare-(15)]
	_ = x[UsbClassAudioVideo-(16)]
	_ = x[UsbClassBillboard-(17)]
	_ = x[UsbClassUSBTypeCBridge-(18)]
	_ = x[UsbClassDiagnosticDevice-(220)]
	_ = x[UsbClassWireless-(224)]
	_ = x[UsbClassMiscellaneous-(239)]
	_ = x[UsbClassApplication-(254)]
	_ = x[UsbClassVendorSpec-(255)]
}

var _UsbClassValues = []UsbClass{UsbClassPerInterface, UsbClassAudio, UsbClassComm, UsbClassHID, UsbClassPhysical, UsbClassImage, UsbClassPrinter, UsbClassMassStorage, UsbClassHub, UsbClassData, UsbClassSmartCard, UsbClassContentSecurity, UsbClassVideo, UsbClassPersonalHealthcare, UsbClassAudioVideo, UsbClassBillboard, UsbClassUSBTypeCBridge, UsbClassDiagnosticDevice, UsbClassWireless, UsbClassMiscellaneous, UsbClassApplication, UsbClassVendorSpec}

var _UsbClassNameToValueMap = map[string]UsbClass{
	_UsbClassName_0[0:13]:       UsbClassPerInterface,
	_UsbClassLowerName_0[0:13]:  UsbClassPerInterface,
	_UsbClassName_0[13:18]:      UsbClassAudio,
	_UsbClassLowerName_0[13:18]: UsbClassAudio,
	_UsbClassName_0[18:22]:      UsbClassComm,
	_UsbClassLowerName_0[18:22]: UsbClassComm,
	_UsbClassName_0[22:25]:      UsbClassHID,
	_UsbClassLowerName_0[22:25]: UsbClassHID,
	_UsbClassName_1[0:8]:        UsbClassPhysical,
	_UsbClassLowerName_1[0:8]:   UsbClassPhysical,
	_UsbClassName_1[8:13]:       UsbClassImage,
	_UsbClassLowerName_1[8:13]:  UsbClassImage,
	_UsbClassName_1[13:20]:      UsbClassPrinter,
	_UsbClassLowerName_1[13:20]: UsbClassPrinter,
	_UsbClassName_1[20:32]:      UsbClassMassStorage,
	_UsbClassLowerName_1[20:32]: UsbClassMassStorage,
	_UsbClassName_1[32:35]:      UsbClassHub,
	_UsbClassLowerName_1[32:35]: UsbClassHub,
	_UsbClassName_1[35:39]:      UsbClassData,
	_UsbClassLowerName_1[35:39]: UsbClassData,
	_UsbClassName_1[39:49]:      UsbClassSmartCard,
	_UsbClassLowerName_1[39:49]: UsbClassSmartCard,
	_UsbClassName_2[0:16]:       UsbClassContentSecurity,
	_UsbClassLowerName_2[0:16]:  UsbClassContentSecurity,
	_UsbClassName_2[16:21]:      UsbClassVideo,
	_UsbClassLowerName_2[16:21]: UsbClassVideo,
	_UsbClassName_2[21:40]:      UsbClassPersonalHealthcare,
	_UsbClassLowerName_2[21:40]: UsbClassPersonalHealthcare,
	_UsbClassName_2[40:51]:      UsbClassAudioVideo,
	_UsbClassLowerName_2[40:51]: UsbClassAudioVideo,
	_UsbClassName_2[51:60]:      UsbClassBillboard,
	_UsbClassLowerName_2[51:60]: UsbClassBillboard,
	_UsbClassName_2[60:77]:      UsbClassUSBTypeCBridge,
	_UsbClassLowerName_2[60:77]: UsbClassUSBTypeCBridge,
	_UsbClassName_3[0:17]:       UsbClassDiagnosticDevice,
	_UsbClassLowerName_3[0:17]:  UsbClassDiagnosticDevice,
	_UsbClassName_4[0:8]:        UsbClassWireless,
	_UsbClassLowerName_4[0:8]:   UsbClassWireless,
	_UsbClassName_5[0:13]:       UsbClassMiscellaneous,
	_UsbClassLowerName_5[0:13]:  UsbClassMiscellaneous,
	_UsbClassName_6[0:11]:       UsbClassApplication,
	_UsbClassLowerName_6[0:11]:  UsbClassApplication,
	_UsbClassName_6[11:22]:      UsbClassVendorSpec,
	_UsbClassLowerName_6[11:22]: UsbClassVendorSpec,
}

var _UsbClassNames = []string{
	_UsbClassName_0[0:13],
	_UsbClassName_0[13:18],
	_UsbClassName_0[18:22],
	_UsbClassName_0[22:25],
	_UsbClassName_1[0:8],
	_UsbClassName_1[8:13],
	_UsbClassName_1[13:20],
	_UsbClassName_1[20:32],
	_UsbClassName_1[32:35],
	_UsbClassName_1[35:39],
	_UsbClassName_1[39:49],
	_UsbClassName_2[0:16],
	_UsbClassName_2[16:21],
	_UsbClassName_2[21:40],
	_UsbClassName_2[40:51],
	_UsbClassName_2[51:60],
	_UsbClassName_2[60:77],
	_UsbClassName_3[0:17],
	_UsbClassName_4[0:8],
	_UsbClassName_5[0:13],
	_UsbClassName_6[0:11],
	_UsbClassName_6[11:22],
}

// UsbClassString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UsbClassString(s string) (UsbClass, error) {
	if val, ok := _UsbClassNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _UsbClassNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UsbClass values", s)
}

// UsbClassValues returns all values of the enum
func UsbClassValues() []UsbClass {
	return _UsbClassValues
}

// UsbClassStrings returns a slice of all String values of the enum
func UsbClassStrings() []string {
	strs := make([]string, len(_UsbClassNames))
	copy(strs, _UsbClassNames)
	return strs
}

// IsAUsbClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UsbClass) IsAUsbClass() bool {
	for _, v := range _UsbClassValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for UsbClass
func (i UsbClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsbClass
func (i *UsbClass) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("UsbClass should be a string, got %s", data)
	}

	var err error
	*i, err = UsbClassString(s)
	return err
}
