package udev

/*
#cgo pkg-config: libudev
#include <libudev.h>
#include <stdlib.h>
#include <string.h>

struct kv_pair {
    char *key;
    char *value;
};

struct kv_pair *facter_udev_get_device_properties(struct udev_device *device, size_t *size) {
	*size = 0;
	size_t count = 0;
	struct kv_pair *results = NULL;
	struct udev_list_entry *entry = NULL;

	udev_list_entry_foreach(
		entry,
		udev_device_get_properties_list_entry(device)
	) {
		struct kv_pair *new_results = realloc(results, (count + 1) * sizeof(struct kv_pair));

		if (!new_results) {
      		// Handle allocation failure
			for (size_t i = 0; i < count; ++i) {
				free(results[i].key);
				free(results[i].value);
			}
			free(results);
			return NULL;
    	}

		results = new_results;
		results[count].key = strdup(udev_list_entry_get_name(entry));
		results[count].value = strdup(udev_list_entry_get_value(entry));

		count++;
	};

	*size = count;
	return results;
}
*/
import "C"

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"
)

//go:generate enumer -type=Type -json -text -trimprefix Type -output=./udev_type.go
type Type int //nolint:recvcheck

const (
	TypeDisk Type = iota
	TypeCD
	TypeFloppy
	TypeTape
	TypePartition
	TypeUsb
	TypeScsi
	TypePci
	TypeNetwork
	TypeMouse
	TypeKeyboard
	TypePrinter
	TypeAudio
	TypeVideo
	TypeGeneric
)

var ErrNotFound = errors.New("udev data not found")

type Input struct {
	IsAccelerometer       bool
	IsJoystick            bool
	IsJoystickIntegration bool
	IsKey                 bool
	IsKeyboard            bool
	IsMouse               bool
	IsPointingStick       bool
	IsSwitch              bool
	IsTablet              bool
	IsTabletPad           bool
	IsTouchpad            bool
	IsTouchpadIntegration bool
	IsTouchscreen         bool
	IsTrackball           bool
}

func NewUdevInput(env map[string]string) *Input {
	if env["ID_INPUT"] != "1" {
		return nil
	}

	return &Input{
		IsAccelerometer:       env["ID_INPUT_ACCELEROMETER"] == "1",
		IsJoystick:            env["ID_INPUT_JOYSTICK"] == "1",
		IsJoystickIntegration: env["ID_INPUT_JOYSTICK_INTEGRATION"] == "1",
		IsKey:                 env["ID_INPUT_KEY"] == "1",
		IsKeyboard:            env["ID_INPUT_KEYBOARD"] == "1",
		IsMouse:               env["ID_INPUT_MOUSE"] == "1",
		IsPointingStick:       env["ID_INPUT_POINTINGSTICK"] == "1",
		IsSwitch:              env["ID_INPUT_SWITCH"] == "1",
		IsTablet:              env["ID_INPUT_TABLET"] == "1",
		IsTabletPad:           env["ID_INPUT_TABLET_PAD"] == "1",
		IsTouchpad:            env["ID_INPUT_TOUCHPAD"] == "1",
		IsTouchpadIntegration: env["ID_INPUT_TOUCHPAD_INTEGRATION"] == "1",
		IsTouchscreen:         env["ID_INPUT_TOUCHSCREEN"] == "1",
		IsTrackball:           env["ID_INPUT_TRACKBALL"] == "1",
	}
}

type Usb struct {
	Model        string
	ModelID      uint16
	Vendor       string
	VendorID     uint16
	Revision     uint16
	Serial       string
	Type         string
	Interfaces   string
	InterfaceNum string
	Driver       string
}

// usbEnv returns the value of the ID_USB_-prefixed key, falling back to the
// unprefixed name. The prefixed variants were introduced in systemd 252
// (which exports both); older udev daemons (Ubuntu 22.04 ships 249) write only
// the unprefixed names to /run/udev/data.
func usbEnv(env map[string]string, key string) string {
	if v, ok := env["ID_USB_"+key]; ok {
		return v
	}

	return env["ID_"+key]
}

func NewUdevUsb(env map[string]string) (*Usb, error) {
	if bus := env["ID_BUS"]; bus != "usb" {
		return nil, fmt.Errorf("invalid bus: %s", bus)
	}

	result := &Usb{
		Model:        usbEnv(env, "MODEL"),
		Vendor:       usbEnv(env, "VENDOR"),
		Serial:       env["ID_SERIAL"],
		Type:         usbEnv(env, "TYPE"),
		Interfaces:   env["ID_USB_INTERFACES"],
		InterfaceNum: env["ID_USB_INTERFACE_NUM"],
		Driver:       env["ID_USB_DRIVER"],
	}

	modelID, err := strconv.ParseUint(usbEnv(env, "MODEL_ID"), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse model id: %w", err)
	}

	result.ModelID = uint16(modelID)

	vendorID, err := strconv.ParseUint(usbEnv(env, "VENDOR_ID"), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vendor id: %w", err)
	}

	result.VendorID = uint16(vendorID)

	revision, err := strconv.ParseUint(usbEnv(env, "REVISION"), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse revision: %w", err)
	}

	result.Revision = uint16(revision)

	return result, nil
}

type Pci struct {
	Class     string
	SubClass  string
	Interface string
}

func NewUdevPci(env map[string]string) (*Pci, error) {
	if bus := env["ID_BUS"]; bus != "pci" {
		return nil, fmt.Errorf("invalid bus: %s", bus)
	}

	result := &Pci{
		Class:     env["ID_PCI_CLASS"],
		SubClass:  env["ID_PCI_SUBCLASS"],
		Interface: env["ID_PCI_INTERFACE"],
	}

	return result, nil
}

type Udev struct {
	Type        Type
	Model       string
	ModelID     uint16
	Vendor      string
	VendorID    uint16
	Revision    uint16
	Serial      string
	SerialShort string

	Usb   *Usb
	Pci   *Pci
	Input *Input
}

func NewUdev(env map[string]string) (*Udev, error) {
	result := &Udev{
		Model:       env["ID_MODEL"],
		Vendor:      env["ID_VENDOR"],
		Serial:      env["ID_SERIAL"],
		SerialShort: env["ID_SERIAL_SHORT"],
		Input:       NewUdevInput(env),
	}

	if str, ok := env["ID_MODEL_ID"]; ok {
		modelID, err := strconv.ParseUint(str, 16, 16)
		if err != nil {
			return nil, fmt.Errorf("failed to parse model id: %w", err)
		}

		result.ModelID = uint16(modelID)
	}

	if str, ok := env["ID_VENDOR_ID"]; ok {
		vendorID, err := strconv.ParseUint(str, 16, 16)
		if err != nil {
			return nil, fmt.Errorf("failed to parse vendor id: %w", err)
		}

		result.VendorID = uint16(vendorID)
	}

	if str, ok := env["ID_REVISION"]; ok {
		revision, err := strconv.ParseUint(str, 16, 16)
		if err != nil {
			return nil, fmt.Errorf("failed to parse revision: %w", err)
		}

		result.Revision = uint16(revision)
	}

	// ID_BUS is set by systemd-udev rules and has an open-ended value space
	// (usb, pci, scsi, ata, acpi, platform, virtio, ...); only dispatch on
	// the buses we have parsers for. See issue #554.
	switch env["ID_BUS"] {
	case "usb":
		usb, err := NewUdevUsb(env)
		if err != nil {
			return nil, fmt.Errorf("failed to parse usb: %w", err)
		}

		result.Usb = usb
	case "pci":
		pci, err := NewUdevPci(env)
		if err != nil {
			return nil, fmt.Errorf("failed to parse pci: %w", err)
		}

		result.Pci = pci
	}

	return result, nil
}

func Read(sysPath string) (*Udev, error) {
	udev := C.udev_new()
	if udev == nil {
		return nil, errors.New("failed to create udev")
	}

	defer C.udev_unref(udev)

	device := C.udev_device_new_from_syspath(udev, C.CString(sysPath))
	if device == nil {
		return nil, ErrNotFound
	}

	defer C.udev_device_unref(device)

	count := C.size_t(0)

	propsArray := C.facter_udev_get_device_properties(device, &count)
	if propsArray == nil {
		return nil, errors.New("failed to get device properties")
	}

	defer C.free(unsafe.Pointer(propsArray))

	env := make(map[string]string)
	propsSlice := unsafe.Slice(propsArray, count)

	for idx := range propsSlice {
		kv := (*C.struct_kv_pair)(unsafe.Pointer(&propsSlice[idx]))

		key := C.GoString(kv.key)
		value := C.GoString(kv.value)

		C.free(unsafe.Pointer(kv.key))
		C.free(unsafe.Pointer(kv.value))

		env[key] = value
	}

	return NewUdev(env)
}

// Version retrieves the major version of the running systemd-udevd daemon.
// It asks systemd for the daemon's main PID (assuming systemctl is on PATH)
// and executes the daemon's own binary with --version, so the result reflects
// the running daemon rather than whatever udevadm happens to be installed.
func Version() (uint64, error) {
	out, err := exec.CommandContext(
		context.Background(),
		"systemctl", "show", "--property=MainPID", "--value", "systemd-udevd.service",
	).Output()
	if err != nil {
		return 0, fmt.Errorf("failed to query systemd-udevd.service via systemctl: %w", err)
	}

	pidStr := strings.TrimSpace(string(out))

	pid, err := strconv.ParseUint(pidStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse systemd-udevd MainPID from %q: %w", pidStr, err)
	}

	if pid == 0 {
		return 0, errors.New("systemd-udevd.service is not running")
	}

	exe := filepath.Join("/proc", strconv.FormatUint(pid, 10), "exe")

	output, err := exec.CommandContext(context.Background(), exe, "--version").Output()
	if err != nil {
		return 0, fmt.Errorf("failed to run %s --version: %w", exe, err)
	}

	return ParseVersion(string(output))
}

// ParseVersion extracts the numeric udev version from --version output.
// udevadm prints a bare number (e.g. "260"); systemd-udevd prints
// "systemd 249 (249.11-0ubuntu3)", so take the first numeric field of the
// first line.
func ParseVersion(output string) (uint64, error) {
	line, _, _ := strings.Cut(output, "\n")

	for field := range strings.FieldsSeq(line) {
		if version, err := strconv.ParseUint(field, 10, 16); err == nil {
			return version, nil
		}
	}

	return 0, fmt.Errorf("failed to parse udev version from %q", line)
}
