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

	rq := require.New(t)

	result, err := udev.NewUdev(map[string]string{
		"ID_BUS":       "acpi",
		"ID_INPUT":     "1",
		"ID_INPUT_KEY": "1",
	})
	rq.NoError(err, "NewUdev must accept ID_BUS=acpi without error")
	rq.NotNil(result)
	rq.NotNil(result.Input)
	rq.True(result.Input.IsKey)
}

// TestNewUdevUsbOldKeyNames reproduces issue #202: udev < 252 exports only the
// unprefixed ID_MODEL_ID / ID_VENDOR_ID / ID_REVISION keys (the ID_USB_*
// variants arrived in systemd 252), which previously aborted the whole scan.
// The env below is the reporter's verbatim /run/udev/data/+input:input4 from
// udev 249 (Razer keyboard).
func TestNewUdevUsbOldKeyNames(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	result, err := udev.NewUdevUsb(map[string]string{
		"ID_INPUT":             "1",
		"ID_INPUT_MOUSE":       "1",
		"ID_VENDOR":            "Razer",
		"ID_VENDOR_ENC":        "Razer",
		"ID_VENDOR_ID":         "1532",
		"ID_MODEL":             "DeathStalker_Ultimate",
		"ID_MODEL_ENC":         `DeathStalker\x20Ultimate`,
		"ID_MODEL_ID":          "0114",
		"ID_REVISION":          "0100",
		"ID_SERIAL":            "Razer_DeathStalker_Ultimate",
		"ID_TYPE":              "hid",
		"ID_BUS":               "usb",
		"ID_USB_INTERFACES":    ":030102:000000:030001:030101:fff000:",
		"ID_USB_INTERFACE_NUM": "00",
		"ID_USB_DRIVER":        "usbhid",
	})
	rq.NoError(err, "udev 249 key names must not abort the scan")
	rq.NotNil(result)
	rq.Equal(uint16(0x1532), result.VendorID)
	rq.Equal(uint16(0x0114), result.ModelID)
	rq.Equal(uint16(0x0100), result.Revision)
	rq.Equal("DeathStalker_Ultimate", result.Model)
	rq.Equal("Razer", result.Vendor)
	rq.Equal("hid", result.Type)
}

// New key names take precedence when both are present (systemd >= 252 exports both).
func TestNewUdevUsbPrefixedKeyNames(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	result, err := udev.NewUdevUsb(map[string]string{
		"ID_BUS":           "usb",
		"ID_USB_MODEL_ID":  "c52b",
		"ID_USB_VENDOR_ID": "046d",
		"ID_USB_REVISION":  "2411",
		// stale values the parser must ignore when prefixed keys are present
		"ID_MODEL_ID":   "ffff",
		"ID_VENDOR_ID":  "ffff",
		"ID_REVISION":   "ffff",
		"ID_USB_MODEL":  "USB_Receiver",
		"ID_USB_VENDOR": "Logitech",
	})
	rq.NoError(err)
	rq.Equal(uint16(0xc52b), result.ModelID)
	rq.Equal(uint16(0x046d), result.VendorID)
	rq.Equal(uint16(0x2411), result.Revision)
}

func TestParseVersion(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	// bare number, as udevadm and systemd-udevd print it
	version, err := udev.ParseVersion("260\n")
	rq.NoError(err)
	rq.Equal(uint64(260), version)

	// distro-decorated form
	version, err = udev.ParseVersion("systemd 249 (249.11-0ubuntu3)\n+PAM +AUDIT\n")
	rq.NoError(err)
	rq.Equal(uint64(249), version)

	_, err = udev.ParseVersion("no numbers here\n")
	rq.Error(err)
}
