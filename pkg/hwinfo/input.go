package hwinfo

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/numtide/nixos-facter/pkg/linux/input"
	"github.com/numtide/nixos-facter/pkg/udev"
)

// captureTouchpads scans the input devices and identifies touchpads, returning a slice of HardwareDevice structs or an
// error.
// It accepts a deviceIdx to ensure it continues on from the last device index generated by hwinfo.
func captureTouchpads(deviceIdx uint16) ([]HardwareDevice, error) {
	inputDevices, err := input.ReadDevices(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read input devices: %w", err)
	}

	var result []HardwareDevice //nolint:prealloc

	for _, inputDevice := range inputDevices {
		path := "/sys" + inputDevice.Sysfs

		udevData, err := udev.Read(path)
		if errors.Is(err, udev.ErrNotFound) {
			slog.Warn("udev data not found", "name", inputDevice.Name, "sysfs", inputDevice.Sysfs)
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to fetch udev data for device %q with udev data: %w", path, err)
		}

		if udevData.Input == nil {
			slog.Debug("udev data missing input", "name", inputDevice.Name, "sysfs", inputDevice.Sysfs)
			continue
		}

		if !udevData.Input.IsTouchpad {
			// currently, we are only interested in touchpads, as hwinfo does not capture them
			// eventually, we may take over more of the input processing that hwinfo performs
			continue
		}

		if len(inputDevice.Handlers) == 0 {
			// I believe this shouldn't be possible, and if it does occur, we should error
			return nil, fmt.Errorf("no handlers found for input device %s", inputDevice.Sysfs)
		}

		// create a hardware entry for the report
		hd := HardwareDevice{
			// todo AttachedTo: it's unclear how to work this out
			Class:     HardwareClassMouse,
			BaseClass: NewBaseClassID(BaseClassTouchpad),
			Vendor: &ID{
				Name:  udevData.Vendor,
				Value: inputDevice.Vendor,
			},
			Device: &ID{
				Name:  udevData.Model,
				Value: inputDevice.Product,
			},
			SysfsID: inputDevice.Sysfs,
		}

		switch inputDevice.Bus {
		case input.BusI2c:
			hd.BusType = NewBusID(BusSerial)
			hd.SubClass = &ID{
				Name:  SubClassMouseBus.String(),
				Value: uint16(SubClassMouseSer),
			}

		case input.BusUsb:
			hd.BusType = NewBusID(BusUsb)
			hd.SubClass = &ID{
				Name:  SubClassMouseUsb.String(),
				Value: uint16(SubClassMouseUsb),
			}
		case input.BusI8042:
			hd.BusType = NewBusID(BusPs2)
			hd.SubClass = &ID{
				Name:  SubClassMousePs2.String(),
				Value: uint16(SubClassMousePs2),
			}

		case input.BusRmi:
			// RMI is a protocol which runs over other physical buses, typically i2c, but it can also be over others
			// such as USB or SPI.
			// I'm not sure how to map this into hwinfo's bus classification, so for now we will use other for both the
			// bus and mouse subclass.
			hd.BusType = NewBusID(BusOther)
			hd.SubClass = &ID{
				Name:  SubClassMouseOther.String(),
				Value: uint16(SubClassMouseOther),
			}

		case input.BusPci,
			input.BusIsapnp,
			input.BusHil,
			input.BusBluetooth,
			input.BusVirtual,
			input.BusIsa,
			input.BusXtkbd,
			input.BusRs232,
			input.BusGameport,
			input.BusParport,
			input.BusAmiga,
			input.BusAdb,
			input.BusHost,
			input.BusGsc,
			input.BusAtari,
			input.BusSpi,
			input.BusCec,
			input.BusIntelIshtp,
			input.BusAmdSfh:
			// todo unsure if touchpads can be on any other bus type
			return nil, fmt.Errorf("unsupported bus type: %s", inputDevice.Bus)
		}

		// todo should we error if no event handler is found?
		if handler := inputDevice.EventHandler(); handler != "" {
			hd.UnixDeviceNames = append(hd.UnixDeviceNames, "/dev/input/"+handler)
		}

		if handler := inputDevice.MouseHandler(); handler != "" {
			hd.UnixDeviceNames = append(hd.UnixDeviceNames, "/dev/input/ + handler")
		}

		hd.Index = deviceIdx
		result = append(result, hd)

		deviceIdx++
	}

	return result, nil
}
