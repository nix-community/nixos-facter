package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "unsafe"

func ReadStringList(list *C.str_list_t) []string {
	var result []string
	for entry := list; entry != nil; entry = entry.next {
		result = append(result, C.GoString(entry.str))
	}
	return result
}

func ReadUint64Array(arr unsafe.Pointer, length int) []uint64 {
	start := uintptr(arr)
	result := make([]uint64, length)
	for i := range result {
		next := start + uintptr(i*C.sizeof_uint64_t)
		result[i] = *((*uint64)(unsafe.Pointer(next))) //nolint:govet
	}
	return result
}

func ReadUintArray(arr unsafe.Pointer, length int) []uint {
	start := uintptr(arr)
	result := make([]uint, length)
	for i := range result {
		next := start + uintptr(i*C.sizeof_uint)
		result[i] = *((*uint)(unsafe.Pointer(next))) //nolint:govet
	}
	return result
}

func ReadIntArray(arr unsafe.Pointer, length int) []int {
	// TODO see if we can use generics to combine some of these methods
	start := uintptr(arr)
	result := make([]int, length)
	for i := range result {
		next := start + uintptr(i*C.sizeof_uint)
		result[i] = *((*int)(unsafe.Pointer(next))) //nolint:govet
	}
	return result
}

func ReadByteArray(arr unsafe.Pointer, length int) []byte {
	start := uintptr(arr)
	result := make([]byte, length)
	for i := range result {
		next := start + uintptr(i*C.sizeof_uint)
		result[i] = *((*byte)(unsafe.Pointer(next))) //nolint:govet
	}
	return result
}
