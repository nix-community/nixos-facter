//nolint:testpackage
package hwinfo

import (
	"testing"

	"github.com/numtide/nixos-facter/pkg/udev"
)

// Regression test for https://github.com/nix-community/nixos-facter/issues/481:
// i2c HID multitouch devices with INPUT_PROP_DIRECT are classified by udev as
// touchscreens, not touchpads, and were previously dropped from the report.
func TestPointerBaseClass(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *udev.Input
		want  BaseClass
		ok    bool
	}{
		{name: "nil input", input: nil},
		{name: "touchpad", input: &udev.Input{IsTouchpad: true}, want: BaseClassTouchpad, ok: true},
		{name: "touchscreen", input: &udev.Input{IsTouchscreen: true}, want: BaseClassTouchscreen, ok: true},
		// e.g. HAILUCK 258a combos present as a plain mouse emulator and hwinfo already captures those.
		{name: "plain mouse", input: &udev.Input{IsMouse: true}},
		{name: "keyboard", input: &udev.Input{IsKeyboard: true}},
		// a touchpad that also claims touchscreen resolves to touchpad
		{
			name:  "touchpad wins over touchscreen",
			input: &udev.Input{IsTouchpad: true, IsTouchscreen: true},
			want:  BaseClassTouchpad,
			ok:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, ok := pointerBaseClass(tc.input)
			if ok != tc.ok {
				t.Fatalf("ok: expected %v, got %v", tc.ok, ok)
			}

			if ok && got != tc.want {
				t.Fatalf("base class: expected %s, got %s", tc.want, got)
			}
		})
	}
}
