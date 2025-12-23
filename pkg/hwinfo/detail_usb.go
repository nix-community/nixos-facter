package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

//go:generate enumer -type=UsbClass -json -transform=snake -trimprefix UsbClass -output=./detail_usb_enum_usb_class.go
type UsbClass uint16 //nolint:recvcheck

const (
	UsbClassPerInterface       UsbClass = 0x00
	UsbClassAudio              UsbClass = 0x01
	UsbClassComm               UsbClass = 0x02
	UsbClassHID                UsbClass = 0x03
	UsbClassPhysical           UsbClass = 0x05
	UsbClassImage              UsbClass = 0x06
	UsbClassPTP                UsbClass = UsbClassImage // legacy name for image
	UsbClassPrinter            UsbClass = 0x07
	UsbClassMassStorage        UsbClass = 0x08
	UsbClassHub                UsbClass = 0x09
	UsbClassData               UsbClass = 0x0a
	UsbClassSmartCard          UsbClass = 0x0b
	UsbClassContentSecurity    UsbClass = 0x0d
	UsbClassVideo              UsbClass = 0x0e
	UsbClassPersonalHealthcare UsbClass = 0x0f
	UsbClassAudioVideo         UsbClass = 0x10
	UsbClassBillboard          UsbClass = 0x11
	UsbClassUSBTypeCBridge     UsbClass = 0x12
	UsbClassDiagnosticDevice   UsbClass = 0xdc
	UsbClassWireless           UsbClass = 0xe0
	UsbClassMiscellaneous      UsbClass = 0xef
	UsbClassApplication        UsbClass = 0xfe
	UsbClassVendorSpec         UsbClass = 0xff
)

type DetailUsb struct {
	Type DetailType `json:"-"`

	DeviceClass    ID  `json:"device_class"`
	DeviceSubclass ID  `json:"device_subclass"`
	DeviceProtocol int `json:"device_protocol"`

	InterfaceClass            ID  `json:"interface_class"`
	InterfaceSubclass         ID  `json:"interface_subclass"`
	InterfaceProtocol         int `json:"interface_protocol"`
	InterfaceNumber           int `json:"interface_number"`
	InterfaceAlternateSetting int `json:"interface_alternate_setting"`

	InterfaceAssociation *DetailUsbInterfaceAssociation `json:"interface_association,omitempty"`
}

type DetailUsbInterfaceAssociation struct {
	FunctionClass    ID  `json:"function_class"`
	FunctionSubclass ID  `json:"function_subclass"`
	FunctionProtocol int `json:"function_protocol"`
	InterfaceCount   int `json:"interface_count"`
	FirstInterface   int `json:"first_interface"`
}

func (d DetailUsb) DetailType() DetailType {
	return DetailTypeUsb
}

func NewDetailUsb(usb C.hd_detail_usb_t) (*DetailUsb, error) {
	data := usb.data
	if data == nil {
		return nil, nil
	}

	if data.next != nil {
		println("usb next is not nil")
	}

	detail := &DetailUsb{
		Type: DetailTypeUsb,
		DeviceClass: ID{
			Type:  IDTagUsb,
			Value: uint16(data.d_cls),
			Name:  UsbClass(data.d_cls).String(),
		},
		DeviceSubclass: ID{
			Type:  IDTagUsb,
			Value: uint16(data.d_sub),
			Name:  UsbClass(data.d_sub).String(),
		},
		DeviceProtocol: int(data.d_prot),
		InterfaceClass: ID{
			Type:  IDTagUsb,
			Value: uint16(data.i_cls),
			Name:  UsbClass(data.i_cls).String(),
		},
		InterfaceSubclass: ID{
			Type:  IDTagUsb,
			Value: uint16(data.i_sub),
			Name:  UsbClass(data.i_sub).String(),
		},
		InterfaceProtocol:         int(data.i_prot),
		InterfaceNumber:           int(data.ifdescr),
		InterfaceAlternateSetting: int(data.i_alt),
	}

	// The Interface Association Descriptor groups multiple interfaces that are part of a single functional device.
	// For instance, a USB webcam with an integrated microphone would use an IAD to group the video input interface and
	// the audio input interface together.
	if data.iad_i_count > 0 {
		detail.InterfaceAssociation = &DetailUsbInterfaceAssociation{
			FunctionClass: ID{
				Type:  IDTagUsb,
				Value: uint16(data.iad_f_cls),
				Name:  UsbClass(data.iad_f_cls).String(),
			},
			FunctionSubclass: ID{
				Type:  IDTagUsb,
				Value: uint16(data.iad_f_sub),
				Name:  UsbClass(data.iad_f_sub).String(),
			},
			FunctionProtocol: int(data.iad_f_prot),
			FirstInterface:   int(data.iad_i_first),
			InterfaceCount:   int(data.iad_i_count),
		}
	}

	return detail, nil
}
