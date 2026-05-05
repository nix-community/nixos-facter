package udev_test

import (
	"testing"

	"github.com/numtide/nixos-facter/pkg/udev"
	"github.com/stretchr/testify/require"
)

// TestNewUdevAcpiBus reproduces issue #554: systemd >= 260 tags ACPI input
// devices (e.g. LNXPWRBN, LNXVIDEO) with ID_BUS=acpi, which previously aborted
// the whole hardware scan because the udev parser tried to coerce ID_BUS into
// the kernel input.Bus enum that has no acpi member.
func TestNewUdevAcpiBus(t *testing.T) {
	t.Parallel()

	result, err := udev.NewUdev(map[string]string{
		"ID_BUS":       "acpi",
		"ID_INPUT":     "1",
		"ID_INPUT_KEY": "1",
	})
	require.NoError(t, err, "NewUdev must accept ID_BUS=acpi without error")
	require.NotNil(t, result)
	require.NotNil(t, result.Input)
	require.True(t, result.Input.IsKey)
}
