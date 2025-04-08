package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_cache_t hd_res_get_cache(hd_res_t *res) { return res->cache; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceCache struct {
	Type ResourceType `json:"type"`
	Size string       `json:"size"`
}

func (r ResourceCache) ResourceType() ResourceType {
	return ResourceTypeCache
}

func NewResourceCache(res *C.hd_res_t, resType ResourceType) (*ResourceCache, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeCache {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeCache, resType)
	}

	cache := C.hd_res_get_cache(res)

	return &ResourceCache{
		Type: resType,
		Size: fmt.Sprintf("0x%x", uint(cache.size)),
	}, nil
}
