package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_smbios_type_t hd_smbios_get_type(hd_smbios_t *sm) { return sm->any.type; }
smbios_biosinfo_t hd_smbios_get_biosinfo(hd_smbios_t *sm) { return sm->biosinfo; }
smbios_sysinfo_t hd_smbios_get_sysinfo(hd_smbios_t *sm) { return sm->sysinfo; }
smbios_boardinfo_t hd_smbios_get_boardinfo(hd_smbios_t *sm) { return sm->boardinfo; }
smbios_chassis_t hd_smbios_get_chassis(hd_smbios_t *sm) { return sm->chassis; }
smbios_processor_t hd_smbios_get_processor(hd_smbios_t *sm) { return sm->processor; }
smbios_cache_t hd_smbios_get_cache(hd_smbios_t *sm) { return sm->cache; }
smbios_connect_t hd_smbios_get_connect(hd_smbios_t *sm) { return sm->connect; }
smbios_slot_t hd_smbios_get_slot(hd_smbios_t *sm) { return sm->slot; }
smbios_onboard_t hd_smbios_get_onboard(hd_smbios_t *sm) { return sm->onboard; }
smbios_oem_t hd_smbios_get_oem(hd_smbios_t *sm) { return sm->oem; }
smbios_config_t hd_smbios_get_config(hd_smbios_t *sm) { return sm->config; }
smbios_lang_t hd_smbios_get_lang(hd_smbios_t *sm) { return sm->lang; }
smbios_group_t hd_smbios_get_group(hd_smbios_t *sm) { return sm->group; }
smbios_memarray_t hd_smbios_get_memarray(hd_smbios_t *sm) { return sm->memarray; }
smbios_memarraymap_t hd_smbios_get_memarraymap(hd_smbios_t *sm) { return sm->memarraymap; }
smbios_memdevice_t hd_smbios_get_memdevice(hd_smbios_t *sm) { return sm->memdevice; }
smbios_memerror_t hd_smbios_get_memerror(hd_smbios_t *sm) { return sm->memerror; }
smbios_mem64error_t hd_smbios_get_mem64error(hd_smbios_t *sm) { return sm->mem64error; }
smbios_memdevicemap_t hd_smbios_get_memdevicemap(hd_smbios_t *sm) { return sm->memdevicemap; }
smbios_mouse_t hd_smbios_get_mouse(hd_smbios_t *sm) { return sm->mouse; }
smbios_secure_t hd_smbios_get_secure(hd_smbios_t *sm) { return sm->secure; }
smbios_power_t hd_smbios_get_power(hd_smbios_t *sm) { return sm->power; }
smbios_any_t hd_smbios_get_any(hd_smbios_t *sm) { return sm->any; }
*/
import "C"
import "errors"

//go:generate enumer -type=SmbiosType -json -transform=snake -trimprefix SmbiosType -output=./smbios_enum_type.go
type SmbiosType uint //nolint:recvcheck

// For a full list of types see https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.6.0.pdf.
// hwinfo doesn't provide structs for all of these, but we've ensured we at least have their ids so they format
// well in the json output when used with the Any type.
const (
	SmbiosTypeBios SmbiosType = iota
	SmbiosTypeSystem
	SmbiosTypeBoard
	SmbiosTypeChassis

	SmbiosTypeProcessor
	SmbiosTypeMemoryController
	SmbiosTypeMemoryModule
	SmbiosTypeCache

	SmbiosTypePortConnector
	SmbiosTypeSlot
	SmbiosTypeOnboard
	SmbiosTypeOEMStrings

	SmbiosTypeConfig
	SmbiosTypeLanguage
	SmbiosTypeGroupAssociations
	SmbiosTypeEventLog

	SmbiosTypeMemoryArray
	SmbiosTypeMemoryDevice
	SmbiosTypeMemoryError
	SmbiosTypeMemoryArrayMappedAddress

	SmbiosTypeMemoryDeviceMappedAddress
	SmbiosTypePointingDevice
	SmbiosTypeBattery
	SmbiosTypeSystemReset

	SmbiosTypeHardwareSecurity
	SmbiosTypePowerControls
	SmbiosTypeVoltage
	SmbiosTypeCoolingDevice

	SmbiosTypeTemperature
	SmbiosTypeCurrent
	SmbiosTypeOutOfBandRemoteAccess
	SmbiosTypeBootIntegrityServices

	SmbiosTypeSystemBoot
	SmbiosTypeMemory64Error
	SmbiosTypeManagementDevice
	SmbiosTypeManDeviceComponent
	SmbiosTypeManDeviceThreshold
	SmbiosTypeMemoryChannel
	SmbiosTypeIPMIDevice

	SmbiosTypeSystemPowerSupply
	SmbiosTypeAdditionalInfo
	SmbiosTypeOnboardExtended
	SmbiosTypeManagementControllerHostInterface
	SmbiosTypeTPM
	SmbiosTypeProcessorAdditional
	SmbiosTypeFirmwareInventory

	SmbiosTypeInactive   SmbiosType = 126
	SmbiosTypeEndOfTable SmbiosType = 127
)

type Smbios interface {
	SmbiosType() SmbiosType
}

//nolint:ireturn
func NewSmbios(smbios *C.hd_smbios_t) (Smbios, error) {
	if smbios == nil {
		return nil, errors.New("smbios is nil")
	}

	var (
		err    error
		result Smbios
	)

	//nolint:exhaustive
	switch SmbiosType(C.hd_smbios_get_type(smbios)) {
	case SmbiosTypeBios:
		result, err = NewSmbiosBiosInfo(C.hd_smbios_get_biosinfo(smbios))
	case SmbiosTypeBoard:
		result, err = NewSmbiosBoardInfo(C.hd_smbios_get_boardinfo(smbios))
	case SmbiosTypeCache:
		result, err = NewSmbiosCache(C.hd_smbios_get_cache(smbios))
	case SmbiosTypeChassis:
		result, err = NewSmbiosChassis(C.hd_smbios_get_chassis(smbios))
	case SmbiosTypeConfig:
		result, err = NewSmbiosConfig(C.hd_smbios_get_config(smbios))
	case SmbiosTypeGroupAssociations:
		result, err = NewSmbiosGroup(C.hd_smbios_get_group(smbios))
	case SmbiosTypeHardwareSecurity:
		result, err = NewSmbiosSecure(C.hd_smbios_get_secure(smbios))
	case SmbiosTypeLanguage:
		result, err = NewSmbiosLang(C.hd_smbios_get_lang(smbios))
	case SmbiosTypeMemory64Error:
		result, err = NewSmbiosMem64Error(C.hd_smbios_get_mem64error(smbios))
	case SmbiosTypeMemoryArray:
		result, err = NewSmbiosMemArray(C.hd_smbios_get_memarray(smbios))
	case SmbiosTypeMemoryArrayMappedAddress:
		result, err = NewSmbiosMemArrayMap(C.hd_smbios_get_memarraymap(smbios))
	case SmbiosTypeMemoryDevice:
		result, err = NewSmbiosMemDevice(C.hd_smbios_get_memdevice(smbios))
	case SmbiosTypeMemoryDeviceMappedAddress:
		result, err = NewSmbiosMemDeviceMap(C.hd_smbios_get_memdevicemap(smbios))
	case SmbiosTypeMemoryError:
		result, err = NewSmbiosMemError(C.hd_smbios_get_memerror(smbios))
	case SmbiosTypeOEMStrings:
		// At least for framework, this contains asset_tags. Since it's unstructured information, we skip it for now
	case SmbiosTypeOnboard:
		result, err = NewSmbiosOnboard(C.hd_smbios_get_onboard(smbios))
	case SmbiosTypePointingDevice:
		result, err = NewSmbiosMouse(C.hd_smbios_get_mouse(smbios))
	case SmbiosTypePortConnector:
		result, err = NewSmbiosConnect(C.hd_smbios_get_connect(smbios))
	case SmbiosTypePowerControls:
		result, err = NewSmbiosPower(C.hd_smbios_get_power(smbios))
	case SmbiosTypeProcessor:
		result, err = NewSmbiosProcessor(C.hd_smbios_get_processor(smbios))
	case SmbiosTypeSlot:
		result, err = NewSmbiosSlot(C.hd_smbios_get_slot(smbios))
	case SmbiosTypeSystem:
		result, err = NewSmbiosSysInfo(C.hd_smbios_get_sysinfo(smbios))

		// We could return Any by default, but it's just noise in the report.
		// As we support new types, users can run it again.
	default:
	}

	return result, err
}
