package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// ProbeFeature is a type that specifies various hardware probing features.
//
//go:generate enumer -type=ProbeFeature -json -transform=snake -trimprefix ProbeFeature -output=./hardware_enum_probe_feature.go
type ProbeFeature uint //nolint:recvcheck

//nolint:revive,stylecheck
const (
	ProbeFeatureMemory ProbeFeature = iota + 1
	ProbeFeaturePci
	ProbeFeatureIsapnp
	ProbeFeatureNet
	ProbeFeatureFloppy
	ProbeFeatureMisc

	ProbeFeatureMiscSerial
	ProbeFeatureMiscPar
	ProbeFeatureMiscFloppy
	ProbeFeatureSerial
	ProbeFeatureCpu
	ProbeFeatureBios

	ProbeFeatureMonitor
	ProbeFeatureMouse
	ProbeFeatureScsi
	ProbeFeatureUsb
	ProbeFeatureUsbMods
	ProbeFeatureAdb
	ProbeFeatureModem

	ProbeFeatureModemUsb
	ProbeFeatureParallel
	ProbeFeatureParallelLp
	ProbeFeatureParallelZip
	ProbeFeatureIsa

	ProbeFeatureIsaIsdn
	ProbeFeatureIsdn
	ProbeFeatureKbd
	ProbeFeatureProm
	ProbeFeatureSbus
	ProbeFeatureInt
	ProbeFeatureBraille

	ProbeFeatureBrailleAlva
	ProbeFeatureBrailleFhp
	ProbeFeatureBrailleHt
	ProbeFeatureIgnx11
	ProbeFeatureSys

	ProbeFeatureBiosVbe
	ProbeFeatureIsapnpOld
	ProbeFeatureIsapnpNew
	ProbeFeatureIsapnpMod
	ProbeFeatureBrailleBaum

	ProbeFeatureManual
	ProbeFeatureFb
	ProbeFeatureVeth
	ProbeFeaturePppoe
	ProbeFeatureScan
	ProbeFeaturePcmcia
	ProbeFeatureFork

	ProbeFeatureParallelImm
	ProbeFeatureS390
	ProbeFeatureCpuemu
	ProbeFeatureSysfs
	ProbeFeatureS390disks
	ProbeFeatureUdev

	ProbeFeatureBlock
	ProbeFeatureBlockCdrom
	ProbeFeatureBlockPart
	ProbeFeatureEdd
	ProbeFeatureEddMod
	ProbeFeatureBiosDdc

	ProbeFeatureBiosFb
	ProbeFeatureBiosMode
	ProbeFeatureInput
	ProbeFeatureBlockMods
	ProbeFeatureBiosVesa

	ProbeFeatureCpuemuDebug
	ProbeFeatureScsiNoserial
	ProbeFeatureWlan
	ProbeFeatureBiosCrc
	ProbeFeatureHal

	ProbeFeatureBiosVram
	ProbeFeatureBiosAcpi
	ProbeFeatureBiosDdcPorts
	ProbeFeatureModulesPata

	ProbeFeatureNetEeprom
	ProbeFeatureX86emu

	ProbeFeatureMax
	ProbeFeatureLxrc
	ProbeFeatureDefault

	ProbeFeatureAll
)

// HardwareClass represents the classification of different hardware components.
//
//go:generate enumer -type=HardwareClass -json -transform=snake -trimprefix HardwareClass -output=./hardware_enum_hardware_class.go
type HardwareClass uint //nolint:recvcheck

//nolint:revive,stylecheck
const (
	HardwareClassNone HardwareClass = iota
	HardwareClassSystem
	HardwareClassCpu
	HardwareClassKeyboard
	HardwareClassBraille
	HardwareClassMouse

	HardwareClassJoystick
	HardwareClassPrinter
	HardwareClassScanner
	HardwareClassChipCard
	HardwareClassMonitor
	HardwareClassTvCard

	HardwareClassGraphicsCard
	HardwareClassFrameBuffer
	HardwareClassCamera
	HardwareClassSound
	HardwareClassStorageController

	HardwareClassNetworkController
	HardwareClassIsdnAdapter
	HardwareClassModem
	HardwareClassNetworkInterface
	HardwareClassDisk
	HardwareClassPartition

	HardwareClassCdrom
	HardwareClassFloppy
	HardwareClassManual
	HardwareClassUsbController
	HardwareClassUsb
	HardwareClassBios
	HardwareClassPci

	HardwareClassIsapnp
	HardwareClassBridge
	HardwareClassHub
	HardwareClassScsi
	HardwareClassIde
	HardwareClassMemory
	HardwareClassDvbCard

	HardwareClassPcmcia
	HardwareClassPcmciaController
	HardwareClassFirewire
	HardwareClassFirewireController
	HardwareClassHotplug

	HardwareClassHotplugController
	HardwareClassZip
	HardwareClassPppoe
	HardwareClassWlanCard
	HardwareClassRedasd
	HardwareClassDslAdapter
	HardwareClassBlockDevice

	HardwareClassTape
	HardwareClassVesaBios
	HardwareClassBluetooth
	HardwareClassFingerprint
	HardwareClassMmcController
	HardwareClassNvme

	HardwareClassUnknown
	HardwareClassAll
)

// BaseClass values (superset of PCI classes)
//
//go:generate enumer -type=BaseClass -json -transform=snake -trimprefix BaseClass -output=./hardware_enum_base_class.go
type BaseClass uint16 //nolint:recvcheck

// these *must* match standard PCI class numbers
const (
	BaseClassNone BaseClass = iota
	BaseClassStorage
	BaseClassNetwork
	BaseClassDisplay
	BaseClassMultimedia

	BaseClassMemory
	BaseClassBridge
	BaseClassComm
	BaseClassSystem
	BaseClassInput
	BaseClassDocking

	BaseClassProcessor
	BaseClassSerial
	BaseClassWireless
	BaseClassI2o
	BaseClassOther BaseClass = 0xff
)

// add our own base classes here (starting at 0x100 as PCI values are 8 bit)
const (
	BaseClassMonitor BaseClass = iota + 0x100
	BaseClassInternal
	BaseClassModem
	BaseClassIsdn
	BaseClassPs2
	BaseClassMouse

	BaseClassStorageDevice
	BaseClassNetworkInterface
	BaseClassKeyboard
	BaseClassPrinter

	BaseClassHub
	BaseClassBraille
	BaseClassScanner
	BaseClassJoystick
	BaseClassChipcard
	BaseClassCamera

	BaseClassFramebuffer
	BaseClassDvb
	BaseClassTv
	BaseClassPartition
	BaseClassDsl
	BaseClassBluetooth
	BaseClassFingerprint

	BaseClassMmcController
	BaseClassTouchpad
)

// SubClassKeyboard values of BaseClassKeyboard
//
//go:generate enumer -type=SubClassKeyboard -json -transform=snake -trimprefix SubClassKeyboard -output=./hardware_enum_sub_class_keyboard.go
type SubClassKeyboard uint16 //nolint:recvcheck

const (
	SubClassKeyboardKbd SubClassKeyboard = iota
	SubClassKeyboardConsole
)

// SubClassMouse values of BaseClassMouse
//
//go:generate enumer -type=SubClassMouse -json -transform=snake -trimprefix SubClassMouse -output=./hardware_enum_sub_class_mouse.go
type SubClassMouse uint16 //nolint:recvcheck

const (
	SubClassMousePs2 SubClassMouse = iota
	SubClassMouseSer
	SubClassMouseBus
	SubClassMouseUsb
	SubClassMouseSun
	SubClassMouseOther SubClassMouse = 0x80
)

// Slot represents a bus and slot number.
// Bits 0-7: slot number, 8-31 bus number
type Slot uint

func (s *Slot) Slot() byte {
	return byte(*s & 0xFF)
}

func (s *Slot) Bus() uint {
	return uint((*s & 0xFFFFFF00) >> 8)
}

func (s *Slot) String() string {
	return fmt.Sprintf("%d:%d", s.Bus(), s.Slot())
}

func (s *Slot) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(map[string]any{
		"bus":    s.Bus(),
		"number": s.Slot(),
	})
	if err != nil {
		err = fmt.Errorf("failed to marshal Slot: %w", err)
	}

	return b, err
}

// DeviceNumber represents a Unix device number, a unique identifier for devices in the system.
type DeviceNumber struct {
	// Type indicates if the device is a character or a block device.
	Type int `json:"type"`
	// Major identifies the driver for a device.
	Major uint `json:"major"`
	// Minor is used by the device driver to distinguish between different devices it controls, or different instances
	// of the same device.
	Minor uint `json:"minor"`
	Range uint `json:"range"`
}

// IsEmpty checks if the DeviceNumber has all zero values for its fields.
func (d DeviceNumber) IsEmpty() bool {
	return d.Type == 0 && d.Major == 0 && d.Minor == 0 && d.Range == 0
}

// NewDeviceNumber creates a new DeviceNumber instance from the given C.hd_dev_num_t and returns a pointer to it.
// Returns nil if the newly created DeviceNumber is empty.
func NewDeviceNumber(num C.hd_dev_num_t) *DeviceNumber {
	result := DeviceNumber{
		Type:  int(num._type),
		Major: uint(num.major),
		Minor: uint(num.minor),
		Range: uint(num._range),
	}
	if result.IsEmpty() {
		return nil
	}
	return &result
}

// Hotplug defines types of hotplug devices.
//
//go:generate enumer -type=Hotplug -json -transform=snake -trimprefix Hotplug -output=./hardware_enum_hotplug.go
type Hotplug int //nolint:recvcheck

const (
	HotplugNone Hotplug = iota
	HotplugPcmcia
	HotplugCardbus
	HotplugPci
	HotplugUsb
	HotplugFirewire
)

// Bus types similar to PCI bridge subclasses
//
//go:generate enumer -type=Bus -json -transform=snake -trimprefix Bus -output=./hardware_enum_bus.go
type Bus uint16 //nolint:recvcheck

const (
	BusNone Bus = iota
	BusIsa
	BusEisa
	BusMc
	BusPci
	BusPcmcia
	BusNubus

	BusCardbus
	BusOther
)

const (
	/** outside the range of the PCI values **/

	BusPs2 Bus = iota + 0x80
	BusSerial
	BusParallel
	BusFloppy
	BusScsi
	BusIde
	BusUsb

	BusAdb
	BusRaid
	BusSbus
	BusI2o
	BusVio
	BusCcw
	BusIucv
	BusPs3SystemBus

	BusVirtio
	BusIbmebus
	BusGameport
	BusUisvirtpci
	BusMmc
	BusSdio
	BusNd

	BusNvme
)

// HardwareDevice represents a hardware component detected in the system.
type HardwareDevice struct {
	// Index is a unique index provided by hwinfo, starting at 1
	Index uint `json:"index"`

	// AttachedTo is the index of the hardware device this is attached to
	AttachedTo uint `json:"attached_to"`

	// Class represents the type of the hardware component.
	Class HardwareClass `json:"-"`

	// Class represents a list of hardware types which matches the component.
	ClassList []HardwareClass `json:"class_list,omitempty"`

	// BusType represents the type of bus to which the hardware device is connected.
	BusType *ID `json:"bus_type,omitempty"`

	// Slot represents a bus and slot number for the hardware device.
	Slot *Slot `json:"slot,omitempty"`

	// BaseClass specifies the base classification of the hardware device.
	BaseClass *ID `json:"base_class,omitempty"`

	// SubClass represents the specific subclass of the hardware component, providing more granular identification
	// within its class.
	SubClass *ID `json:"sub_class,omitempty"`

	// PciInterface specifies the PCI interface identifier of the hardware device.
	PciInterface *ID `json:"pci_interface,omitempty"`

	// Vendor represents the vendor ID of the hardware device.
	Vendor *ID `json:"vendor,omitempty"`

	// SubVendor represents the ID of the subsystem vendor.
	SubVendor *ID `json:"sub_vendor,omitempty"`

	// Device represents the unique identifier for the hardware device.
	Device *ID `json:"device,omitempty"`

	// SubDevice represents the identifier of a sub-device in the hardware
	SubDevice *ID `json:"sub_device,omitempty"`

	// Revision specifies the hardware revision identifier.
	Revision *ID `json:"revision,omitempty"`

	Serial string `json:"serial,omitempty"`

	// CompatVendor is a vendor id and name of some compatible hardware.
	// Used mainly for ISA-PnP devices.
	CompatVendor *ID `json:"compat_vendor,omitempty"`

	// CompatDevice is a device id and name of some compatible hardware.
	// Used mainly for ISA-PnP devices.
	CompatDevice *ID `json:"compat_device,omitempty"`

	// Model is a combination of vendor and device names. Some heuristics are used to make it more presentable.
	Model string `json:"model,omitempty"`

	// SysfsID is a sysfs entry for this hardware, if any.
	SysfsID string `json:"sysfs_id,omitempty"`

	// SysfsBusID is a sysfs bus entry for this hardware, if any.
	SysfsBusID string `json:"sysfs_bus_id,omitempty"`

	// SysfsDeviceLink is the string path to the system file system (sysfs) link for this hardware device.
	SysfsDeviceLink string `json:"sysfs_device_link,omitempty"`

	// SysfsIOMMUGroupID represents the IOMMU group ID associated with the hardware device in sysfs, if any.
	SysfsIOMMUGroupID *IOMMUGroup `json:"sysfs_iommu_group_id,omitempty"`

	// UnixDeviceName is a path to a device file used to access this hardware.
	// Normally something below /dev.
	// For network interfaces, this is the interface name.
	UnixDeviceName string `json:"unix_device_name,omitempty"`

	// UnixDeviceNumber represents the device type and number according to sysfs.
	UnixDeviceNumber *DeviceNumber `json:"unix_device_number,omitempty"`

	// UnixDeviceNames is a list of device names which can be used to access this hardware.
	// Normally something below /dev.
	// They should all be equivalent.
	// The preferred name is UnixDeviceName.
	UnixDeviceNames []string `json:"unix_device_names,omitempty"`

	// UnixDeviceName2 is a path to a device file used to access this hardware.
	// Most hardware only has one device name, stored in UnixDeviceName.
	// In some cases, there's an alternative name.
	UnixDeviceName2 string `json:"unix_device_name2,omitempty"`

	// UnixDeviceNumber2 is an alternative device type and number according to sysfs.
	UnixDeviceNumber2 *DeviceNumber `json:"unix_device_number2,omitempty"`

	// RomID represents a BIOS/PROM id.
	// Where appropriate, this is a special BIOS/PROM id (e.g. "0x80" for the first harddisk on Intel-PCs).
	// CHPID for s390.
	RomID string `json:"rom_id,omitempty"`

	// Udi is a HAL unique device identifier.
	Udi string `json:"udi,omitempty"`

	// ParentUdi is the udi of a parent device, if any.
	ParentUdi string `json:"parent_udi,omitempty"`

	// Resources is a list of device resources.
	Resources []Resource `json:"resources,omitempty"`

	// Detail is specific information associated with this hardware.
	Detail Detail `json:"detail,omitempty"`

	// Hotplug indicates the type of hotplug controller associated with this device, if any.
	Hotplug *Hotplug `json:"hotplug,omitempty"`

	// HotplugSlot indicates the slot this device is connected to, if any (e.g. PCMCIA socket).
	// Counts are 1-based (0: no info available).
	HotplugSlot uint `json:"hotplug_slot,omitempty"`

	// Driver is the currently active driver, if any.
	Driver string `json:"driver,omitempty"`

	// DriverModule is the currently active driver module, if any.
	DriverModule string `json:"driver_module,omitempty"`

	// Drivers is a list of currently active drivers.
	Drivers []string `json:"drivers,omitempty"`

	// DriverModules is a list of currently active driver modules.
	DriverModules []string `json:"driver_modules,omitempty"`

	// DriverInfo is available driver information for the currently active driver, if any.
	DriverInfo DriverInfo `json:"driver_info,omitempty"` // device driver info

	// UsbGUID is a USB Global Unique Identifier.
	// Available for USB devices.
	// This may be set even if the bus type is not USB (e.g. USB storage devices will have bus set to SCSI due to SCSI
	// emulation)
	UsbGUID string `json:"usb_guid,omitempty"` // USB Global Unique Identifier.

	// todo hal_prop
	// todo persistent_prop

	// ModuleAlias for matching and loading kernel modules for this hardware.
	ModuleAlias string `json:"module_alias,omitempty"`

	// Label is a Consistent Device Name (CDN), as per PCI firmware spec 3.1, chapter 4.6.7
	Label string `json:"label,omitempty"`
}

func NewHardwareDevice(hd *C.hd_t, ephemeral bool) (*HardwareDevice, error) {
	if hd == nil {
		return nil, errors.New("hd is nil")
	}

	var (
		err        error
		resources  []Resource
		detail     Detail
		driverInfo DriverInfo
	)

	resources, err = NewResources(hd)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	if hd.detail != nil {
		detail, err = NewDetail(hd.detail)
		if err != nil {
			return nil, fmt.Errorf("failed to read detail: %w", err)
		}
	}

	if hd.driver_info != nil {
		driverInfo, err = NewDriverInfo(hd.driver_info)
		if err != nil {
			return nil, fmt.Errorf("failed to read driver info: %w", err)
		}
	}

	model := C.GoString(hd.model)

	hwClass := HardwareClass(hd.hw_class)
	if hwClass == HardwareClassCpu {
		model = stripCPUFreq(model)
	}

	var hwClassList []HardwareClass
	for i := HardwareClassSystem; i < HardwareClassAll; i++ {
		if C.hd_is_hw_class(hd, C.hd_hw_item_t(i)) == 1 {
			hwClassList = append(hwClassList, i)
		}
	}

	result := &HardwareDevice{
		Index:             uint(hd.idx),
		AttachedTo:        uint(hd.attached_to),
		BusType:           NewID(hd.bus),
		BaseClass:         NewID(hd.base_class),
		SubClass:          NewID(hd.sub_class),
		PciInterface:      NewID(hd.prog_if),
		Vendor:            NewID(hd.vendor),
		SubVendor:         NewID(hd.sub_vendor),
		Device:            NewID(hd.device),
		SubDevice:         NewID(hd.sub_device),
		Revision:          NewID(hd.revision),
		Serial:            C.GoString(hd.serial),
		CompatVendor:      NewID(hd.compat_vendor),
		CompatDevice:      NewID(hd.compat_device),
		Class:             hwClass,
		ClassList:         hwClassList,
		Model:             model,
		SysfsID:           C.GoString(hd.sysfs_id),
		SysfsBusID:        C.GoString(hd.sysfs_bus_id),
		SysfsDeviceLink:   C.GoString(hd.sysfs_device_link),
		UnixDeviceName:    C.GoString(hd.unix_dev_name),
		UnixDeviceNumber:  NewDeviceNumber(hd.unix_dev_num),
		UnixDeviceName2:   C.GoString(hd.unix_dev_name2),
		UnixDeviceNames:   ReadStringList(hd.unix_dev_names),
		UnixDeviceNumber2: NewDeviceNumber(hd.unix_dev_num2),
		RomID:             C.GoString(hd.rom_id),
		Udi:               C.GoString(hd.udi),
		ParentUdi:         C.GoString(hd.parent_udi),
		Resources:         resources,
		Detail:            detail,
		Driver:            C.GoString(hd.driver),
		DriverModule:      C.GoString(hd.driver_module),
		Drivers:           ReadStringList(hd.drivers),
		DriverModules:     ReadStringList(hd.driver_modules),
		UsbGUID:           C.GoString(hd.usb_guid),
		DriverInfo:        driverInfo,
		ModuleAlias:       C.GoString(hd.modalias),
		Label:             C.GoString(hd.label),
	}

	if !ephemeral {
		// /dev/disk/by-diskseq is not stable, so we remove them
		result.UnixDeviceNames = slices.DeleteFunc(result.UnixDeviceNames, func(name string) bool {
			return strings.HasPrefix(name, "/dev/disk/by-diskseq/")
		})
	}

	// only set the slot information if the bus type has been set
	if result.BusType != nil {
		slot := Slot(hd.slot)
		result.Slot = &slot
	}

	// only set hotplug info if it's available
	hotplug := Hotplug(hd.hotplug)
	if hotplug != HotplugNone {
		result.Hotplug = &hotplug
		result.HotplugSlot = uint(hd.hotplug_slot)
	}

	// sort some fields to ensure stable report output
	slices.Sort(result.UnixDeviceNames)
	slices.Sort(result.Drivers)
	slices.Sort(result.DriverModules)

	return result, nil
}
