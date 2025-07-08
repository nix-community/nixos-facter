// Package virt contains utilities for detecting virtualized environments.
// It has been adapted from [systemd].
//
// [systemd]: https://github.com/systemd/systemd/blob/main/src/basic/virt.c
package virt

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

// Type represents various virtualisation and container types.
//
//go:generate enumer -type=Type -json -text -transform=snake -trimprefix Type -output=./virt_enum_type.go
type Type int //nolint:recvcheck

const (
	TypeNone Type = iota
	TypeKvm
	TypeAmazon
	TypeQemu
	TypeBochs
	TypeXen
	TypeUml
	TypeVmware
	TypeOracle
	TypeMicrosoft
	TypeZvm
	TypeParallels
	TypeBhyve
	TypeQnx
	TypeAcrn
	TypePowervm
	TypeApple
	TypeSre
	TypeGoogle
	TypeVmOther
	TypeSystemdNspawn
	TypeLxcLibvirt
	TypeLxc
	TypeOpenvz
	TypeDocker
	TypePodman
	TypeRkt
	TypeWsl
	TypeProot
	TypePouch
	TypeContainerOther
)

// Detect identifies the virtualisation type of the current system.
// Returns the detected Type and an error if detection fails.
func Detect() (Type, error) {
	out, err := exec.Command("systemd-detect-virt").CombinedOutput()
	outStr := strings.Trim(string(out), "\n")

	// note: systemd-detect-virt exits with status 1 when "none" is detected
	//nolint:staticcheck
	if !(outStr == "none" || err == nil) {
		slog.Error("failed to detect virtualisation type", "output", out)

		return TypeNone, fmt.Errorf("failed to detect virtualisation type: %w", err)
	}

	// we use snake case, but systemd-detect-virt uses hyphen case
	virtType := strings.ReplaceAll(outStr, "-", "_")

	return TypeString(virtType)
}
