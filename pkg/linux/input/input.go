package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:generate enumer -type=Bus -json -text -trimprefix Bus -output=./input_bus.go
type Bus uint16 //nolint:recvcheck

// Codes taken from
// https://github.com/torvalds/linux/blob/cfaaa7d010d1fc58f9717fcc8591201e741d2d49/include/uapi/linux/input.h#L254
const (
	BusPci Bus = iota + 1
	BusIsapnp
	BusUsb
	BusHil
	BusBluetooth
	BusVirtual

	BusIsa Bus = iota + 10
	BusI8042
	BusXtkbd
	BusRs232
	BusGameport
	BusParport
	BusAmiga
	BusAdb
	BusI2c
	BusHost
	BusGsc
	BusAtari
	BusSpi
	BusRmi
	BusCec
	BusIntelIshtp
	BusAmdSfh
)

const (
	devicesPath = "/proc/bus/input/devices"
)

var (
	capRegex   = regexp.MustCompile(`^B: (\w+)=(.*)$`)
	nameRegex  = regexp.MustCompile(`^N: Name="(.*)"$`)
	physRegex  = regexp.MustCompile(`^P: Phys=(.*)$`)
	sysfsRegex = regexp.MustCompile(`^S: Sysfs=(.*)$`)
	basicRegex = regexp.MustCompile(
		`^I: Bus=([0-9abcdef]{4}) Vendor=([0-9abcdef]{4}) Product=([0-9abcdef]{4}) Version=([0-9abcdef]{4})$`,
	)

	handlersRegex     = regexp.MustCompile(`^H: Handlers=(.*)`)
	mouseHandlerRegex = regexp.MustCompile(`mouse\d+`)
	eventHandlerRegex = regexp.MustCompile(`event\d+`)
)

type Device struct {
	Bus          Bus
	Vendor       uint16
	Product      uint16
	Version      uint16
	Name         string
	Handlers     []string
	Sysfs        string
	Phys         string
	Capabilities map[string]string
}

func (d *Device) Path() string {
	return "/dev/input/" + d.Name
}

func (d *Device) MouseHandler() string {
	for _, handler := range d.Handlers {
		if mouseHandlerRegex.MatchString(handler) {
			return handler
		}
	}

	return ""
}

func (d *Device) EventHandler() string {
	for _, handler := range d.Handlers {
		if eventHandlerRegex.MatchString(handler) {
			return handler
		}
	}

	return ""
}

func ReadDevices(r io.ReadCloser) ([]*Device, error) {
	var err error

	if r == nil {
		r, err = os.Open(devicesPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open %s: %w", devicesPath, err)
		}
		defer r.Close() //nolint:errcheck
	}

	var (
		device  *Device
		devices []*Device
	)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			devices = append(devices, device)
			device = nil

			continue
		}

		if device == nil {
			device = &Device{
				Capabilities: make(map[string]string),
			}
		}

		switch line[:3] {
		case "I: ":
			if err := readBasicInfo(line, device); err != nil {
				return nil, fmt.Errorf("failed to read basic info: %w", err)
			}
		case "N: ":
			matches := nameRegex.FindStringSubmatch(line)
			if len(matches) != 2 {
				return nil, fmt.Errorf("invalid name: %s", line)
			}

			device.Name = matches[1]

		case "P: ":
			matches := physRegex.FindStringSubmatch(line)
			if len(matches) != 2 {
				return nil, fmt.Errorf("invalid phys: %s", line)
			}

			device.Phys = matches[1]

		case "S: ":
			matches := sysfsRegex.FindStringSubmatch(line)
			if len(matches) != 2 {
				return nil, fmt.Errorf("invalid sysfs: %s", line)
			}

			device.Sysfs = matches[1]

		case "B: ":
			matches := capRegex.FindStringSubmatch(line)
			if len(matches) != 3 {
				return nil, fmt.Errorf("invalid capability: %s", line)
			}

			device.Capabilities[matches[1]] = matches[2]

		case "H: ":
			matches := handlersRegex.FindStringSubmatch(line)
			if len(matches) != 2 {
				return nil, fmt.Errorf("invalid handlers: %s", line)
			}

			device.Handlers = strings.Split(
				strings.Trim(matches[1], " "),
				" ",
			)

		default: // do nothing
		}
	}

	return devices, nil
}

func readBasicInfo(line string, device *Device) error {
	matches := basicRegex.FindStringSubmatch(line)

	if len(matches) != 5 {
		return fmt.Errorf("invalid basic info: %s", line)
	}

	bus, err := strconv.ParseUint(matches[1], 16, 16)
	if err != nil {
		return fmt.Errorf("invalid bus: %s", matches[1])
	}

	device.Bus = Bus(bus)

	vendor, err := strconv.ParseUint(matches[2], 16, 16)
	if err != nil {
		return fmt.Errorf("invalid vendor: %s", matches[2])
	}

	device.Vendor = uint16(vendor)

	product, err := strconv.ParseUint(matches[3], 16, 16)
	if err != nil {
		return fmt.Errorf("invalid product: %s", matches[3])
	}

	device.Product = uint16(product)

	version, err := strconv.ParseUint(matches[4], 16, 16)
	if err != nil {
		return fmt.Errorf("invalid version: %s", matches[4])
	}

	device.Version = uint16(version)

	return nil
}
