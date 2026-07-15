//nolint:gosec
package hwinfo_test

import (
	"testing"
	"unsafe"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/stretchr/testify/require"
)

// These tests lay out buffers exactly as the C structs do (unsigned char *,
// int *, unsigned *, uint64_t *) and verify the readers decode them
// element-for-element.

func TestReadUint64Array(t *testing.T) {
	t.Parallel()

	src := []uint64{0xdeadbeefcafebabe, 1, 0, 42}
	got := hwinfo.ReadUint64Array(unsafe.Pointer(&src[0]), len(src))
	require.Equal(t, src, got)

	require.Nil(t, hwinfo.ReadUint64Array(nil, 4))
	require.Nil(t, hwinfo.ReadUint64Array(unsafe.Pointer(&src[0]), 0))
}

func TestReadUintArray(t *testing.T) {
	t.Parallel()

	// C `unsigned` is 32-bit on all supported platforms.
	src := []uint32{0xffffffff, 0, 1, 7}
	got := hwinfo.ReadUintArray(unsafe.Pointer(&src[0]), len(src))
	require.Equal(t, []uint{0xffffffff, 0, 1, 7}, got)

	require.Nil(t, hwinfo.ReadUintArray(nil, 4))
	require.Nil(t, hwinfo.ReadUintArray(unsafe.Pointer(&src[0]), -1))
}

func TestReadIntArray(t *testing.T) {
	t.Parallel()

	// C `int` is 32-bit; adjacent non-zero elements catch 8-byte over-reads.
	src := []int32{1, 2, 3, -4}
	got := hwinfo.ReadIntArray(unsafe.Pointer(&src[0]), len(src))
	require.Equal(t, []int{1, 2, 3, -4}, got)

	require.Nil(t, hwinfo.ReadIntArray(nil, 4))
}

func TestReadByteArray(t *testing.T) {
	t.Parallel()

	// C `unsigned char` is 1 byte; distinct values catch 4-byte strides.
	src := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	got := hwinfo.ReadByteArray(unsafe.Pointer(&src[0]), len(src))
	require.Equal(t, src, got)

	require.Nil(t, hwinfo.ReadByteArray(nil, 4))
}
