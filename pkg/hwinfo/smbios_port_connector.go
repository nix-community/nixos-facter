package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPortConnector captures port connector information.
type SmbiosPortConnector struct {
	Type                        SmbiosType `json:"-"`
	Handle                      int        `json:"handle"`
	PortType                    *ID        `json:"port_type"`
	InternalConnectorType       *ID        `json:"internal_connector_type,omitempty"`
	InternalReferenceDesignator string     `json:"internal_reference_designator,omitempty"`
	ExternalConnectorType       *ID        `json:"external_connector_type,omitempty"`
	ExternalReferenceDesignator string     `json:"external_reference_designator,omitempty"`
}

func NewSmbiosConnect(info C.smbios_connect_t) (*SmbiosPortConnector, error) {
	return &SmbiosPortConnector{
		Type:                        SmbiosTypePortConnector,
		Handle:                      int(info.handle),
		PortType:                    NewID(info.port_type),
		InternalConnectorType:       NewID(info.i_type),
		InternalReferenceDesignator: C.GoString(info.i_des),
		ExternalConnectorType:       NewID(info.x_type),
		ExternalReferenceDesignator: C.GoString(info.x_des),
	}, nil
}

func (s SmbiosPortConnector) SmbiosType() SmbiosType {
	return s.Type
}
