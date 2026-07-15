package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"slices"
	"unsafe"
)

func ReadStringList(list *C.str_list_t) []string {
	var result []string
	for entry := list; entry != nil; entry = entry.next {
		result = append(result, C.GoString(entry.str))
	}

	return result
}

// ReadUint64Array copies a C uint64_t array into a Go slice.
func ReadUint64Array(arr unsafe.Pointer, length int) []uint64 {
	if arr == nil || length <= 0 {
		return nil
	}

	return slices.Clone(unsafe.Slice((*uint64)(arr), length))
}

// ReadUintArray copies a C unsigned array into a Go slice.
func ReadUintArray(arr unsafe.Pointer, length int) []uint {
	if arr == nil || length <= 0 {
		return nil
	}

	result := make([]uint, length)
	for i, v := range unsafe.Slice((*C.uint)(arr), length) {
		result[i] = uint(v)
	}

	return result
}

// ReadIntArray copies a C int array into a Go slice.
func ReadIntArray(arr unsafe.Pointer, length int) []int {
	if arr == nil || length <= 0 {
		return nil
	}

	result := make([]int, length)
	for i, v := range unsafe.Slice((*C.int)(arr), length) {
		result[i] = int(v)
	}

	return result
}

// ReadByteArray copies a C unsigned char array into a Go slice.
func ReadByteArray(arr unsafe.Pointer, length int) []byte {
	if arr == nil || length <= 0 {
		return nil
	}

	return slices.Clone(unsafe.Slice((*byte)(arr), length))
}
