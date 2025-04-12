package facter

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

// Smbios captures comprehensive information about the system's hardware components
// as reported by the System Management BIOS (SMBIOS).
type Smbios struct {
	// Bios provides detailed information about the system's BIOS, including vendor, version, date, features, and ROM size.
	Bios *hwinfo.SmbiosBios `json:"bios,omitempty"`

	// Board holds motherboard information such as manufacturer, product, and version.
	Board *hwinfo.SmbiosBoard `json:"board,omitempty"`

	// Cache provides details about the system's cache components.
	Cache []hwinfo.Smbios `json:"cache,omitempty"`

	// Chassis holds information related to the system's chassis, including manufacturer, version, and lock presence.
	Chassis []hwinfo.Smbios `json:"chassis,omitempty"`

	// Config captures system configuration options.
	Config *hwinfo.SmbiosConfig `json:"config,omitempty"`

	// GroupAssociations lists associations between different hardware groups in the system.
	GroupAssociations []hwinfo.Smbios `json:"group_associations,omitempty"`

	// HardwareSecurity provides information on the system's hardware security configurations.
	HardwareSecurity []hwinfo.Smbios `json:"hardware_security,omitempty"`

	// Language contains language-related information, including supported and current languages.
	Language []hwinfo.Smbios `json:"language,omitempty"`

	// Memory64Error provides information on 64-bit memory errors.
	Memory64Error []hwinfo.Smbios `json:"memory_64_error,omitempty"`

	// MemoryArray details the physical memory arrays present in the system.
	MemoryArray []hwinfo.Smbios `json:"memory_array,omitempty"`

	// MemoryArrayMappedAddress provides the mapped addresses of memory arrays.
	MemoryArrayMappedAddress []hwinfo.Smbios `json:"memory_array_mapped_address,omitempty"`

	// MemoryDevice captures information about individual memory devices in the system.
	MemoryDevice []hwinfo.Smbios `json:"memory_device,omitempty"`

	// MemoryDeviceMappedAddress provides the mapped addresses of memory devices.
	MemoryDeviceMappedAddress []hwinfo.Smbios `json:"memory_device_mapped_address,omitempty"`

	// MemoryError provides information on memory errors detected in the system.
	MemoryError []hwinfo.Smbios `json:"memory_error,omitempty"`

	// Onboard lists the onboard devices present in the system.
	Onboard []hwinfo.Smbios `json:"onboard,omitempty"`

	// PointingDevice details the system's pointing devices.
	PointingDevice []hwinfo.Smbios `json:"pointing_device,omitempty"`

	// PortConnector lists the port connectors present on the system.
	PortConnector []hwinfo.Smbios `json:"port_connector,omitempty"`

	// PowerControls provides information on the power control mechanisms in the system.
	PowerControls []hwinfo.Smbios `json:"power_controls,omitempty"`

	// Processor captures details about the processors used in the system.
	Processor []hwinfo.Smbios `json:"processor,omitempty"`

	// Slot lists the expansion slots available in the system.
	Slot []hwinfo.Smbios `json:"slot,omitempty"`

	// System captures overall system-related information such as manufacturer, product, version, and UUID.
	System *hwinfo.SmbiosSystem `json:"system,omitempty"`
}

func (s *Smbios) add(item hwinfo.Smbios) error {
	slog.Debug("smbios.add", "type", item.SmbiosType())

	switch item.SmbiosType() {
	case hwinfo.SmbiosTypeBios:
		if s.Bios != nil {
			return errors.New("bios field is already set")
		} else if bios, ok := item.(*hwinfo.SmbiosBios); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosBios, found %T", item)
		} else {
			s.Bios = bios
		}
	case hwinfo.SmbiosTypeBoard:
		if s.Board != nil {
			return errors.New("board field is already set")
		} else if board, ok := item.(*hwinfo.SmbiosBoard); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosBoard, found %T", item)
		} else {
			s.Board = board
		}
	case hwinfo.SmbiosTypeCache:
		s.Cache = append(s.Cache, item)
	case hwinfo.SmbiosTypeConfig:
		if s.Config != nil {
			return errors.New("config field is already set")
		} else if config, ok := item.(*hwinfo.SmbiosConfig); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosConfig, found %T", item)
		} else {
			s.Config = config
		}
	case hwinfo.SmbiosTypeChassis:
		s.Chassis = append(s.Chassis, item)
	case hwinfo.SmbiosTypeGroupAssociations:
		s.GroupAssociations = append(s.GroupAssociations, item)
	case hwinfo.SmbiosTypeHardwareSecurity:
		s.GroupAssociations = append(s.GroupAssociations, item)
	case hwinfo.SmbiosTypeLanguage:
		s.Language = append(s.Language, item)
	case hwinfo.SmbiosTypeMemory64Error:
		s.Memory64Error = append(s.Memory64Error, item)
	case hwinfo.SmbiosTypeMemoryArray:
		s.MemoryArray = append(s.MemoryArray, item)
	case hwinfo.SmbiosTypeMemoryArrayMappedAddress:
		s.MemoryArrayMappedAddress = append(s.MemoryArrayMappedAddress, item)
	case hwinfo.SmbiosTypeMemoryDevice:
		s.MemoryDevice = append(s.MemoryDevice, item)
	case hwinfo.SmbiosTypeMemoryDeviceMappedAddress:
		s.MemoryDeviceMappedAddress = append(s.MemoryDeviceMappedAddress, item)
	case hwinfo.SmbiosTypeMemoryError:
		s.MemoryError = append(s.MemoryError, item)
	case hwinfo.SmbiosTypeOnboard:
		s.Onboard = append(s.Onboard, item)
	case hwinfo.SmbiosTypePointingDevice:
		s.PointingDevice = append(s.PointingDevice, item)
	case hwinfo.SmbiosTypePortConnector:
		s.PortConnector = append(s.PortConnector, item)
	case hwinfo.SmbiosTypePowerControls:
		s.PowerControls = append(s.PowerControls, item)
	case hwinfo.SmbiosTypeProcessor:
		s.Processor = append(s.Processor, item)
	case hwinfo.SmbiosTypeSlot:
		s.Slot = append(s.Slot, item)
	case hwinfo.SmbiosTypeSystem:
		if s.System != nil {
			return errors.New("system field is already set")
		} else if system, ok := item.(*hwinfo.SmbiosSystem); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosSystem, found %T", item)
		} else {
			s.System = system
		}

	case hwinfo.SmbiosTypeMemoryController, hwinfo.SmbiosTypeMemoryModule, hwinfo.SmbiosTypeOEMStrings,
		hwinfo.SmbiosTypeEventLog, hwinfo.SmbiosTypeBattery, hwinfo.SmbiosTypeSystemReset, hwinfo.SmbiosTypeVoltage,
		hwinfo.SmbiosTypeCoolingDevice, hwinfo.SmbiosTypeTemperature, hwinfo.SmbiosTypeCurrent,
		hwinfo.SmbiosTypeOutOfBandRemoteAccess, hwinfo.SmbiosTypeBootIntegrityServices, hwinfo.SmbiosTypeSystemBoot,
		hwinfo.SmbiosTypeManagementDevice, hwinfo.SmbiosTypeManDeviceComponent, hwinfo.SmbiosTypeManDeviceThreshold,
		hwinfo.SmbiosTypeMemoryChannel, hwinfo.SmbiosTypeIPMIDevice, hwinfo.SmbiosTypeSystemPowerSupply,
		hwinfo.SmbiosTypeAdditionalInfo, hwinfo.SmbiosTypeOnboardExtended,
		hwinfo.SmbiosTypeManagementControllerHostInterface, hwinfo.SmbiosTypeTPM, hwinfo.SmbiosTypeProcessorAdditional,
		hwinfo.SmbiosTypeFirmwareInventory, hwinfo.SmbiosTypeInactive, hwinfo.SmbiosTypeEndOfTable:
		// currently not supported
	default:
		return fmt.Errorf("unknown smbios type %d", item.SmbiosType())
	}

	return nil
}
