package facter

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

// Hardware represents various hardware components and their details.
// The fields are in alphabetical order, and list types are sorted to ensure a consistent and stable report output.
type Hardware struct {
	// Bios holds the BIOS details of the hardware.
	Bios *hwinfo.DetailBios `json:"bios,omitempty"`

	// BlockDevice holds the list of block devices in the hardware.
	BlockDevice []hwinfo.HardwareDevice `json:"block_device,omitempty"`

	// Bluetooth holds the list of Bluetooth devices in the hardware.
	Bluetooth []hwinfo.HardwareDevice `json:"bluetooth,omitempty"`

	// Braille holds the list of Braille devices in the hardware.
	Braille []hwinfo.HardwareDevice `json:"braille,omitempty"`

	// Bridge holds the list of bridge devices in the hardware.
	Bridge []hwinfo.HardwareDevice `json:"bridge,omitempty"`

	// Camera holds the list of camera devices in the hardware.
	Camera []hwinfo.HardwareDevice `json:"camera,omitempty"`

	// Cdrom holds the list of CD-ROM devices in the hardware.
	Cdrom []hwinfo.HardwareDevice `json:"cdrom,omitempty"`

	// ChipCard holds the list of chip card devices in the hardware.
	ChipCard []hwinfo.HardwareDevice `json:"chip_card,omitempty"`

	// CPU holds the list of CPU details in the hardware.
	// There is one entry per physical id.
	CPU []*hwinfo.DetailCPU `json:"cpu,omitempty"`

	// Disk holds the list of disk devices in the hardware.
	Disk []hwinfo.HardwareDevice `json:"disk,omitempty"`

	// DslAdapter holds the list of DSL adapter devices in the hardware.
	DslAdapter []hwinfo.HardwareDevice `json:"dsl_adapter,omitempty"`

	// DvbCard holds the list of DVB card devices in the hardware.
	DvbCard []hwinfo.HardwareDevice `json:"dvb_card,omitempty"`

	// Fingerprint holds the list of fingerprint devices in the hardware.
	Fingerprint []hwinfo.HardwareDevice `json:"fingerprint,omitempty"`

	// Firewire holds the list of Firewire devices in the hardware.
	Firewire []hwinfo.HardwareDevice `json:"firewire,omitempty"`

	// FirewireController holds the list of Firewire controllers in the hardware.
	FirewireController []hwinfo.HardwareDevice `json:"firewire_controller,omitempty"`

	// Floppy holds the list of floppy disk devices in the hardware.
	Floppy []hwinfo.HardwareDevice `json:"floppy,omitempty"`

	// FrameBuffer holds the list of frame buffer devices in the hardware.
	FrameBuffer []hwinfo.HardwareDevice `json:"frame_buffer,omitempty"`

	// GraphicsCard holds the list of graphics card devices in the hardware.
	GraphicsCard []hwinfo.HardwareDevice `json:"graphics_card,omitempty"`

	// Hotplug holds the list of hotplug devices in the hardware.
	Hotplug []hwinfo.HardwareDevice `json:"hotplug,omitempty"`

	// HotplugController holds the list of hotplug controllers in the hardware.
	HotplugController []hwinfo.HardwareDevice `json:"hotplug_controller,omitempty"`

	// Hub holds the list of hub devices in the hardware.
	Hub []hwinfo.HardwareDevice `json:"hub,omitempty"`

	// Ide holds the list of IDE devices in the hardware.
	Ide []hwinfo.HardwareDevice `json:"ide,omitempty"`

	// Isapnp holds the list of ISAPNP devices in the hardware.
	Isapnp []hwinfo.HardwareDevice `json:"isapnp,omitempty"`

	// IsdnAdapter holds the list of ISDN adapter devices in the hardware.
	IsdnAdapter []hwinfo.HardwareDevice `json:"isdn_adapter,omitempty"`

	// Joystick holds the list of joystick devices in the hardware.
	Joystick []hwinfo.HardwareDevice `json:"joystick,omitempty"`

	// Keyboard holds the list of keyboard devices in the hardware.
	Keyboard []hwinfo.HardwareDevice `json:"keyboard,omitempty"`

	// Manual holds the list of manually configured devices in the hardware.
	Manual []hwinfo.HardwareDevice `json:"manual,omitempty"`

	// Memory holds the list of memory devices in the hardware.
	Memory []hwinfo.HardwareDevice `json:"memory,omitempty"`

	// MmcController holds the list of MMC controller devices in the hardware.
	MmcController []hwinfo.HardwareDevice `json:"mmc_controller,omitempty"`

	// Modem holds the list of modem devices in the hardware.
	Modem []hwinfo.HardwareDevice `json:"modem,omitempty"`

	// Monitor holds the list of monitor devices in the hardware.
	Monitor []hwinfo.HardwareDevice `json:"monitor,omitempty"`

	// Mouse holds the list of mouse devices in the hardware.
	Mouse []hwinfo.HardwareDevice `json:"mouse,omitempty"`

	// NetworkController holds the list of network controller devices in the hardware.
	NetworkController []hwinfo.HardwareDevice `json:"network_controller,omitempty"`

	// NetworkInterface holds the list of network interface devices in the hardware.
	NetworkInterface []hwinfo.HardwareDevice `json:"network_interface,omitempty"`

	// None holds the list of devices in the hardware with no specific class.
	None []hwinfo.HardwareDevice `json:"none,omitempty"`

	// Nvme holds the list of NVMe devices in the hardware.
	Nvme []hwinfo.HardwareDevice `json:"nvme,omitempty"`

	// Partition holds the list of partition devices in the hardware.
	Partition []hwinfo.HardwareDevice `json:"partition,omitempty"`

	// Pci holds the list of PCI devices in the hardware.
	Pci []hwinfo.HardwareDevice `json:"pci,omitempty"`

	// Pcmcia holds the list of PCMCIA devices in the hardware.
	Pcmcia []hwinfo.HardwareDevice `json:"pcmcia,omitempty"`

	// PcmciaController holds the list of PCMCIA controllers in the hardware.
	PcmciaController []hwinfo.HardwareDevice `json:"pcmcia_controller,omitempty"`

	// Pppoe holds the list of PPPoE devices in the hardware.
	Pppoe []hwinfo.HardwareDevice `json:"pppoe,omitempty"`

	// Printer holds the list of printer devices in the hardware.
	Printer []hwinfo.HardwareDevice `json:"printer,omitempty"`

	// Redasd holds the list of REDASD devices in the hardware.
	Redasd []hwinfo.HardwareDevice `json:"redasd,omitempty"`

	// Scanner holds the list of scanner devices in the hardware.
	Scanner []hwinfo.HardwareDevice `json:"scanner,omitempty"`

	// Scsi holds the list of SCSI devices in the hardware.
	Scsi []hwinfo.HardwareDevice `json:"scsi,omitempty"`

	// Sound holds the list of sound devices in the hardware.
	Sound []hwinfo.HardwareDevice `json:"sound,omitempty"`

	// StorageController holds the list of storage controller devices in the hardware.
	StorageController []hwinfo.HardwareDevice `json:"storage_controller,omitempty"`

	// System holds the system details of the hardware.
	System *hwinfo.DetailSys `json:"system,omitempty"`

	// Tape holds the list of tape devices in the hardware.
	Tape []hwinfo.HardwareDevice `json:"tape,omitempty"`

	// TvCard holds the list of TV card devices in the hardware.
	TvCard []hwinfo.HardwareDevice `json:"tv_card,omitempty"`

	// Unknown holds the list of unknown devices in the hardware.
	Unknown []hwinfo.HardwareDevice `json:"unknown,omitempty"`

	// Usb holds the list of USB devices in the hardware.
	Usb []hwinfo.HardwareDevice `json:"usb,omitempty"`

	// UsbController holds the list of USB controller devices in the hardware.
	UsbController []hwinfo.HardwareDevice `json:"usb_controller,omitempty"`

	// VesaBios holds the list of VESA BIOS devices in the hardware.
	VesaBios []hwinfo.HardwareDevice `json:"vesa_bios,omitempty"`

	// WlanCard holds the list of WLAN card devices in the hardware.
	WlanCard []hwinfo.HardwareDevice `json:"wlan_card,omitempty"`

	// Zip holds the list of ZIP drive devices in the hardware.
	Zip []hwinfo.HardwareDevice `json:"zip,omitempty"`
}

func compareDevice(a hwinfo.HardwareDevice, b hwinfo.HardwareDevice) int {
	return int(a.Index - b.Index) //nolint:gosec
}

func (h *Hardware) add(device hwinfo.HardwareDevice) error {
	slog.Debug(
		"hardware.add",
		"index", device.Index,
		"attached_to", device.AttachedTo,
		"class", device.Class,
		"class_list", device.ClassList,
		"bus", device.BusType,
		"sysfs_id", device.SysfsID,
	)

	// start with the overall device class
	class := device.Class

	// attempt to fall back to the class list if it's unknown
	if class == hwinfo.HardwareClassUnknown && len(device.ClassList) > 0 {
		class = device.ClassList[0]
	}

	switch class {
	case hwinfo.HardwareClassBios:
		if h.Bios != nil {
			return errors.New("bios field is already set")
		} else if bios, ok := device.Detail.(*hwinfo.DetailBios); !ok {
			return fmt.Errorf("expected hwinfo.DetailBios, found %T", device.Detail)
		} else {
			h.Bios = bios
		}
	case hwinfo.HardwareClassBlockDevice:
		h.BlockDevice = append(h.BlockDevice, device)
		slices.SortFunc(h.BlockDevice, compareDevice)
	case hwinfo.HardwareClassBluetooth:
		h.Bluetooth = append(h.Bluetooth, device)
		slices.SortFunc(h.Bluetooth, compareDevice)
	case hwinfo.HardwareClassBraille:
		h.Braille = append(h.Braille, device)
		slices.SortFunc(h.Braille, compareDevice)
	case hwinfo.HardwareClassBridge:
		h.Bridge = append(h.Bridge, device)
		slices.SortFunc(h.Bridge, compareDevice)
	case hwinfo.HardwareClassCamera:
		h.Camera = append(h.Camera, device)
		slices.SortFunc(h.Camera, compareDevice)
	case hwinfo.HardwareClassCdrom:
		h.Cdrom = append(h.Cdrom, device)
		slices.SortFunc(h.Cdrom, compareDevice)
	case hwinfo.HardwareClassChipCard:
		h.ChipCard = append(h.ChipCard, device)
		slices.SortFunc(h.ChipCard, compareDevice)
	case hwinfo.HardwareClassCpu:
		cpu, ok := device.Detail.(*hwinfo.DetailCPU)
		if !ok {
			return fmt.Errorf("expected hwinfo.DetailCPU, found %T", device.Detail)
		}

		// We insert by physical id, as we only want one entry per core.
		requiredSize := int(cpu.PhysicalID) + 1 //nolint:gosec

		if len(h.CPU) < requiredSize {
			newItems := make([]*hwinfo.DetailCPU, requiredSize-len(h.CPU))
			h.CPU = append(h.CPU, newItems...)
		}

		// add our new entry
		h.CPU[cpu.PhysicalID] = cpu

		// Sort in ascending order to ensure a stable output.
		// It's possible that the CPU list contains nil pointers as we are not guaranteed to read the CPU's in order of
		// physical id.
		slices.SortFunc(h.CPU, func(a, b *hwinfo.DetailCPU) int {
			switch {
			case a == nil && b == nil:
				return 0
			case a == nil && b != nil:
				return -1
			case a != nil && b == nil:
				return 1
			default:
				return int(a.PhysicalID - b.PhysicalID) //nolint:gosec
			}
		})

	case hwinfo.HardwareClassDisk:
		h.Disk = append(h.Disk, device)
		slices.SortFunc(h.Disk, compareDevice)
	case hwinfo.HardwareClassDslAdapter:
		h.DslAdapter = append(h.DslAdapter, device)
		slices.SortFunc(h.DslAdapter, compareDevice)
	case hwinfo.HardwareClassDvbCard:
		h.DvbCard = append(h.DvbCard, device)
		slices.SortFunc(h.DvbCard, compareDevice)
	case hwinfo.HardwareClassFingerprint:
		h.Fingerprint = append(h.Fingerprint, device)
		slices.SortFunc(h.Fingerprint, compareDevice)
	case hwinfo.HardwareClassFirewire:
		h.Firewire = append(h.Firewire, device)
		slices.SortFunc(h.Firewire, compareDevice)
	case hwinfo.HardwareClassFirewireController:
		h.FirewireController = append(h.FirewireController, device)
		slices.SortFunc(h.FirewireController, compareDevice)
	case hwinfo.HardwareClassFloppy:
		h.Floppy = append(h.Floppy, device)
		slices.SortFunc(h.Floppy, compareDevice)
	case hwinfo.HardwareClassFrameBuffer:
		h.FrameBuffer = append(h.FrameBuffer, device)
		slices.SortFunc(h.FrameBuffer, compareDevice)
	case hwinfo.HardwareClassGraphicsCard:
		h.GraphicsCard = append(h.GraphicsCard, device)
		slices.SortFunc(h.GraphicsCard, compareDevice)
	case hwinfo.HardwareClassHotplug:
		h.Hotplug = append(h.Hotplug, device)
		slices.SortFunc(h.Hotplug, compareDevice)
	case hwinfo.HardwareClassHotplugController:
		h.HotplugController = append(h.HotplugController, device)
		slices.SortFunc(h.HotplugController, compareDevice)
	case hwinfo.HardwareClassHub:
		h.Hub = append(h.Hub, device)
		slices.SortFunc(h.Hub, compareDevice)
	case hwinfo.HardwareClassIde:
		h.Ide = append(h.Ide, device)
		slices.SortFunc(h.Ide, compareDevice)
	case hwinfo.HardwareClassIsapnp:
		h.Isapnp = append(h.Isapnp, device)
		slices.SortFunc(h.Isapnp, compareDevice)
	case hwinfo.HardwareClassIsdnAdapter:
		h.IsdnAdapter = append(h.IsdnAdapter, device)
		slices.SortFunc(h.IsdnAdapter, compareDevice)
	case hwinfo.HardwareClassJoystick:
		h.Joystick = append(h.Joystick, device)
		slices.SortFunc(h.Joystick, compareDevice)
	case hwinfo.HardwareClassKeyboard:
		h.Keyboard = append(h.Keyboard, device)
		slices.SortFunc(h.Keyboard, compareDevice)
	case hwinfo.HardwareClassManual:
		h.Manual = append(h.Manual, device)
		slices.SortFunc(h.Manual, compareDevice)
	case hwinfo.HardwareClassMemory:
		h.Memory = append(h.Memory, device)
		slices.SortFunc(h.Memory, compareDevice)
	case hwinfo.HardwareClassMmcController:
		h.MmcController = append(h.MmcController, device)
		slices.SortFunc(h.MmcController, compareDevice)
	case hwinfo.HardwareClassModem:
		h.Modem = append(h.Modem, device)
		slices.SortFunc(h.Modem, compareDevice)
	case hwinfo.HardwareClassMonitor:
		h.Monitor = append(h.Monitor, device)
		slices.SortFunc(h.Monitor, compareDevice)
	case hwinfo.HardwareClassMouse:
		h.Mouse = append(h.Mouse, device)
		slices.SortFunc(h.Mouse, compareDevice)
	case hwinfo.HardwareClassNetworkController:
		h.NetworkController = append(h.NetworkController, device)
		slices.SortFunc(h.NetworkController, compareDevice)
	case hwinfo.HardwareClassNetworkInterface:
		h.NetworkInterface = append(h.NetworkInterface, device)
		slices.SortFunc(h.NetworkInterface, compareDevice)
	case hwinfo.HardwareClassNone:
		h.None = append(h.None, device)
		slices.SortFunc(h.None, compareDevice)
	case hwinfo.HardwareClassNvme:
		h.Nvme = append(h.Nvme, device)
		slices.SortFunc(h.Nvme, compareDevice)
	case hwinfo.HardwareClassPartition:
		h.Partition = append(h.Partition, device)
		slices.SortFunc(h.Partition, compareDevice)
	case hwinfo.HardwareClassPci:
		h.Pci = append(h.Pci, device)
		slices.SortFunc(h.Pci, compareDevice)
	case hwinfo.HardwareClassPcmcia:
		h.Pcmcia = append(h.Pcmcia, device)
		slices.SortFunc(h.Pcmcia, compareDevice)
	case hwinfo.HardwareClassPcmciaController:
		h.PcmciaController = append(h.PcmciaController, device)
		slices.SortFunc(h.PcmciaController, compareDevice)
	case hwinfo.HardwareClassPppoe:
		h.Pppoe = append(h.Pppoe, device)
		slices.SortFunc(h.Pppoe, compareDevice)
	case hwinfo.HardwareClassPrinter:
		h.Printer = append(h.Printer, device)
		slices.SortFunc(h.Printer, compareDevice)
	case hwinfo.HardwareClassRedasd:
		h.Redasd = append(h.Redasd, device)
		slices.SortFunc(h.Redasd, compareDevice)
	case hwinfo.HardwareClassScanner:
		h.Scanner = append(h.Scanner, device)
		slices.SortFunc(h.Scanner, compareDevice)
	case hwinfo.HardwareClassScsi:
		h.Scsi = append(h.Scsi, device)
		slices.SortFunc(h.Scsi, compareDevice)
	case hwinfo.HardwareClassSound:
		h.Sound = append(h.Sound, device)
		slices.SortFunc(h.Sound, compareDevice)
	case hwinfo.HardwareClassStorageController:
		h.StorageController = append(h.StorageController, device)
		slices.SortFunc(h.StorageController, compareDevice)
	case hwinfo.HardwareClassSystem:
		if h.System != nil {
			return errors.New("system field is already set")
		} else if system, ok := device.Detail.(*hwinfo.DetailSys); !ok {
			return fmt.Errorf("expected hwinfo.DetailSys, found %T", device.Detail)
		} else {
			h.System = system
		}
	case hwinfo.HardwareClassTape:
		h.Tape = append(h.Tape, device)
		slices.SortFunc(h.Tape, compareDevice)
	case hwinfo.HardwareClassTvCard:
		h.TvCard = append(h.TvCard, device)
		slices.SortFunc(h.TvCard, compareDevice)
	case hwinfo.HardwareClassUnknown:
		h.Unknown = append(h.Unknown, device)
		slices.SortFunc(h.Unknown, compareDevice)
	case hwinfo.HardwareClassUsb:
		h.Usb = append(h.Usb, device)
		slices.SortFunc(h.Usb, compareDevice)
	case hwinfo.HardwareClassUsbController:
		h.UsbController = append(h.UsbController, device)
		slices.SortFunc(h.UsbController, compareDevice)
	case hwinfo.HardwareClassVesaBios:
		h.VesaBios = append(h.VesaBios, device)
		slices.SortFunc(h.VesaBios, compareDevice)
	case hwinfo.HardwareClassWlanCard:
		h.WlanCard = append(h.WlanCard, device)
		slices.SortFunc(h.WlanCard, compareDevice)
	case hwinfo.HardwareClassZip:
		h.Zip = append(h.Zip, device)
		slices.SortFunc(h.Zip, compareDevice)
	case hwinfo.HardwareClassAll: // Do nothing, this is used by the hwinfo cli exclusively.
	}

	return nil
}
