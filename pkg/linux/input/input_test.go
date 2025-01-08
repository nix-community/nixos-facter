//nolint:lll
package input_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/numtide/nixos-facter/pkg/linux/input"
	"github.com/stretchr/testify/require"
)

const (
	devices = `I: Bus=0003 Vendor=1038 Product=1634 Version=0111
N: Name="SteelSeries SteelSeries Apex 9 TKL"
P: Phys=usb-0000:08:00.3-2.4.2/input2
S: Sysfs=/devices/pci0000:00/0000:00:01.2/0000:02:00.0/0000:03:08.0/0000:08:00.3/usb3/3-2/3-2.4/3-2.4.2/3-2.4.2:1.2/0003:1038:1634.000B/input/input11
U: Uniq=
H: Handlers=sysrq kbd leds event4 
B: PROP=0
B: EV=120013
B: KEY=1000000000007 ff9f207ac14057ff febeffdfffefffff fffffffffffffffe
B: MSC=10
B: LED=7

I: Bus=0003 Vendor=046d Product=407b Version=0111
N: Name="Logitech MX Vertical"
P: Phys=usb-0000:08:00.3-2.2/input2:1
S: Sysfs=/devices/pci0000:00/0000:00:01.2/0000:02:00.0/0000:03:08.0/0000:08:00.3/usb3/3-2/3-2.2/3-2.2:1.2/0003:046D:C52B.0004/0003:046D:407B.0005/input/input8
U: Uniq=5c-c6-5d-d4
H: Handlers=sysrq kbd leds event1 mouse0
B: PROP=0
B: EV=12001f
B: KEY=3f00033fff 0 0 483ffff17aff32d bfd4444600000000 ffff0001 130ff38b17d007 ffff7bfad9415fff ffbeffdfffefffff fffffffffffffffe
B: REL=1943
B: ABS=100000000
B: MSC=10
B: LED=1f

I: Bus=0003 Vendor=b58e Product=9e84 Version=0100
N: Name="Blue Microphones Yeti Stereo Microphone Consumer Control"
P: Phys=usb-0000:0f:00.3-4/input3
S: Sysfs=/devices/pci0000:00/0000:00:08.1/0000:0f:00.3/usb5/5-4/5-4:1.3/0003:B58E:9E84.0001/input/input0
U: Uniq=797_2020/06/26_02565
H: Handlers=kbd event0
B: PROP=0
B: EV=1b
B: KEY=1 0 7800000000 e000000000000 0
B: ABS=10000000000
B: MSC=10

`
)

func TestReadDevices(t *testing.T) {
	as := require.New(t)
	r := io.NopCloser(bytes.NewReader([]byte(devices)))

	devices, err := input.ReadDevices(r)
	as.NoError(err)

	expected := []*input.Device{
		{
			Bus:      input.BusUsb,
			Vendor:   uint16(0x1038),
			Product:  uint16(0x1634),
			Version:  uint16(0x0111),
			Name:     "SteelSeries SteelSeries Apex 9 TKL", //nolint:dupword
			Phys:     "usb-0000:08:00.3-2.4.2/input2",
			Sysfs:    "/devices/pci0000:00/0000:00:01.2/0000:02:00.0/0000:03:08.0/0000:08:00.3/usb3/3-2/3-2.4/3-2.4.2/3-2.4.2:1.2/0003:1038:1634.000B/input/input11",
			Handlers: []string{"sysrq", "kbd", "leds", "event4"},
			Capabilities: map[string]string{
				"PROP": "0",
				"EV":   "120013",
				"KEY":  "1000000000007 ff9f207ac14057ff febeffdfffefffff fffffffffffffffe",
				"MSC":  "10",
				"LED":  "7",
			},
		},
		{
			Bus:      input.BusUsb,
			Vendor:   uint16(0x046d),
			Product:  uint16(0x407b),
			Version:  uint16(0x0111),
			Name:     "Logitech MX Vertical",
			Phys:     "usb-0000:08:00.3-2.2/input2:1",
			Sysfs:    "/devices/pci0000:00/0000:00:01.2/0000:02:00.0/0000:03:08.0/0000:08:00.3/usb3/3-2/3-2.2/3-2.2:1.2/0003:046D:C52B.0004/0003:046D:407B.0005/input/input8",
			Handlers: []string{"sysrq", "kbd", "leds", "event1", "mouse0"},
			Capabilities: map[string]string{
				"PROP": "0",
				"EV":   "12001f",
				"KEY":  "3f00033fff 0 0 483ffff17aff32d bfd4444600000000 ffff0001 130ff38b17d007 ffff7bfad9415fff ffbeffdfffefffff fffffffffffffffe",
				"REL":  "1943",
				"ABS":  "100000000",
				"MSC":  "10",
				"LED":  "1f",
			},
		},
		{
			Bus:      input.BusUsb,
			Vendor:   uint16(0xb58e),
			Product:  uint16(0x9e84),
			Version:  uint16(0x0100),
			Name:     "Blue Microphones Yeti Stereo Microphone Consumer Control",
			Phys:     "usb-0000:0f:00.3-4/input3",
			Sysfs:    "/devices/pci0000:00/0000:00:08.1/0000:0f:00.3/usb5/5-4/5-4:1.3/0003:B58E:9E84.0001/input/input0",
			Handlers: []string{"kbd", "event0"},
			Capabilities: map[string]string{
				"PROP": "0",
				"EV":   "1b",
				"KEY":  "1 0 7800000000 e000000000000 0",
				"ABS":  "10000000000",
				"MSC":  "10",
			},
		},
	}

	as.EqualValues(expected, devices)
}
