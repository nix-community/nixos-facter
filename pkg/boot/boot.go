// Package boot contains utilities for detecting boot-related information.
package boot

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const (
	efiFirmwarePath     = "/sys/firmware/efi"
	efiPlatformSizePath = "/sys/firmware/efi/fw_platform_size"
)

// UEFIInfo contains information about UEFI firmware support.
type UEFIInfo struct {
	// Supported indicates whether the system booted with UEFI.
	Supported bool `json:"supported"`

	// PlatformSize indicates the firmware platform size in bits (32 or 64).
	// Only populated when Supported is true.
	PlatformSize *uint8 `json:"platform_size,omitempty"`
}

// DetectUEFI detects UEFI boot information.
func DetectUEFI() (*UEFIInfo, error) {
	info := &UEFIInfo{}

	// Check if the EFI firmware directory exists
	stat, err := os.Stat(efiFirmwarePath)
	if errors.Is(err, fs.ErrNotExist) || (err == nil && !stat.IsDir()) {
		slog.Debug("UEFI boot not detected", "path", efiFirmwarePath)

		return info, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to stat %s: %w", efiFirmwarePath, err)
	}

	info.Supported = true

	slog.Debug("UEFI boot detected", "path", efiFirmwarePath)

	// Detect platform size (32 or 64 bit)
	data, err := os.ReadFile(efiPlatformSizePath)
	if err == nil {
		sizeStr := strings.TrimSpace(string(data))

		size, err := strconv.ParseUint(sizeStr, 10, 8)
		if err == nil {
			size8 := uint8(size)
			info.PlatformSize = &size8

			slog.Debug("UEFI platform size detected", "bits", size)
		}
	}

	return info, nil
}
