package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_res_t *hd_res_next(hd_res_t *res) { return res->next; }
hd_resource_types_t hd_res_get_type(hd_res_t *res) { return res->any.type; }
res_any_t hd_res_get_any(hd_res_t *res) { return res->any; }
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"slices"
)

//go:generate enumer -type=ResourceType -json -transform=snake -trimprefix ResourceType -output=./resource_enum_type.go
type ResourceType uint //nolint:recvcheck

const (
	ResourceTypeAny ResourceType = iota
	ResourceTypePhysMem
	ResourceTypeMem
	ResourceTypeIo
	ResourceTypeIrq
	ResourceTypeDma
	ResourceTypeMonitor

	ResourceTypeSize
	ResourceTypeDiskGeo
	ResourceTypeCache
	ResourceTypeBaud
	ResourceTypeInitStrings
	ResourceTypePppdOption

	ResourceTypeFramebuffer
	ResourceTypeHwaddr
	ResourceTypeLink
	ResourceTypeWlan
	ResourceTypeFc
	ResourceTypePhwaddr
)

//nolint:ireturn
func NewResource(res *C.hd_res_t) (Resource, error) {
	if res == nil {
		return nil, errors.New("resource is nil")
	}

	var (
		err    error
		result Resource
	)

	resourceType := ResourceType(C.hd_res_get_type(res))

	switch resourceType {
	case ResourceTypeFc:
		result, err = NewResourceFc(res, resourceType)
	case ResourceTypePhysMem:
		result, err = NewResourcePhysicalMemory(res, resourceType)
	case ResourceTypeMem:
		result, err = NewResourceMemory(res, resourceType)
	case ResourceTypeIo:
		result, err = NewResourceIO(res, resourceType)
	case ResourceTypeIrq:
		result, err = NewResourceIrq(res, resourceType)
	case ResourceTypeDma:
		result, err = NewResourceDma(res, resourceType)
	case ResourceTypeMonitor:
		result, err = NewResourceMonitor(res, resourceType)
	case ResourceTypeSize:
		result, err = NewResourceSize(res, resourceType)
	case ResourceTypeDiskGeo:
		result, err = NewResourceDiskGeo(res, resourceType)
	case ResourceTypeCache:
		result, err = NewResourceCache(res, resourceType)
	case ResourceTypeBaud:
		result, err = NewResourceBaud(res, resourceType)
	case ResourceTypeInitStrings:
		result, err = NewResourceInitStrings(res, resourceType)
	case ResourceTypePppdOption:
		result, err = NewResourcePppdOption(res, resourceType)
	case ResourceTypeFramebuffer:
		result, err = NewResourceFrameBuffer(res, resourceType)
	case ResourceTypeHwaddr, ResourceTypePhwaddr:
		result, err = NewResourceHardwareAddress(res, resourceType)
	case ResourceTypeLink:
		// this is the link status of a network interface and can change when we plug/unplug a cable
	case ResourceTypeWlan:
		result, err = NewResourceWlan(res, resourceType)
	default:
		err = fmt.Errorf("unexpected resource type: %v", resourceType)
	}

	return result, err
}

func NewResources(hd *C.hd_t) ([]Resource, error) {
	var result []Resource
	for res := hd.res; res != nil; res = C.hd_res_next(res) {
		resource, err := NewResource(res)

		if err != nil {
			return nil, err
		}
		if resource == nil {
			continue
		}

		switch resource.ResourceType() {
		// these resources are not stable, so we filter them out
		case ResourceTypeMem, ResourceTypeIrq:
			continue

		default:
			result = append(result, resource)
		}

	}

	slices.SortFunc(result, func(a, b Resource) int {
		// We don't really care about a proper ordering for resources, just a stable sort that is reasonably quick.
		var err error
		jsonA, err := json.Marshal(a)
		if err != nil {
			log.Panicf("failed to marshal resource: %s", err)
		}
		jsonB, err := json.Marshal(b)
		if err != nil {
			log.Panicf("failed to marshal resource: %s", err)
		}

		return bytes.Compare(jsonA, jsonB)
	})

	return result, nil
}
