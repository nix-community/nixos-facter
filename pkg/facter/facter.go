// Package facter contains types and utilities for scanning a system and generating a report, detailing key aspects of
// the system and its connected hardware.
package facter

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/numtide/nixos-facter/pkg/boot"
	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/cloud"
	"github.com/numtide/nixos-facter/pkg/ephem"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/virt"
)

// Report represents a detailed report on the system’s hardware, virtualisation, SMBios, and swap entries.
type Report struct {
	// Version is a monotonically increasing number,
	// used to indicate breaking changes or new features in the report output.
	Version uint16 `json:"version"`

	// System indicates the system architecture e.g. x86_64-linux.
	System string `json:"system"`

	// Virtualisation indicates the type of virtualisation or container environment present on the system.
	Virtualisation virt.Type `json:"virtualisation"`

	// UEFI contains information about UEFI firmware support.
	UEFI *boot.UEFIInfo `json:"uefi"`

	// Hardware provides detailed information about the system's hardware components, such as CPU, memory, and peripherals.
	Hardware Hardware `json:"hardware"`

	// Smbios provides detailed information about the system's SMBios data, such as BIOS, board, chassis, memory,
	// and processors.
	Smbios Smbios `json:"smbios"`

	// Swap contains a list of swap entries representing the system's swap devices or files and their respective details.
	Swap []*ephem.SwapEntry `json:"swap,omitempty"`

	// Cloud holds instance metadata captured from cloud provider metadata services.
	// It is only populated when the corresponding provider capture is enabled.
	Cloud *cloud.Cloud `json:"cloud,omitempty"`
}

// Scanner defines a type responsible for scanning and reporting system hardware information.
type Scanner struct {
	// Swap indicates whether the system swap information should be reported.
	Swap bool

	// Ephemeral indicates whether the scanner should report ephemeral details,
	// such as swap.
	Ephemeral bool

	Cloud struct {
		// Hetzner indicates whether instance metadata should be fetched from the Hetzner metadata service.
		Hetzner bool
	}

	// Features is a list of ProbeFeature types that should be scanned for.
	Features []hwinfo.ProbeFeature
}

// Scan scans the system's hardware and software information and returns a report.
// It also detects IOMMU groups and handles errors gracefully if scanning fails.
func (s *Scanner) Scan() (*Report, error) {
	var err error

	report := Report{
		Version: build.ReportVersion,
	}

	if build.System == "" {
		return nil, errors.New("system is not set")
	}

	report.System = build.System

	slog.Debug("building report", "system", report.System, "version", report.Version)

	var (
		smbios  []hwinfo.Smbios
		devices []hwinfo.HardwareDevice
	)

	slog.Debug("scanning hardware", "features", s.Features)

	smbios, devices, err = hwinfo.Scan(s.Features, s.Ephemeral)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	slog.Debug("reading IOMMU groups")

	// read iommu groups
	iommuGroups, err := hwinfo.IOMMUGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to read iommu groups: %w", err)
	}

	slog.Debug("processing devices", "count", len(devices))

	for idx := range devices {
		// lookup iommu group before adding to the report
		device := devices[idx]

		groupID, ok := iommuGroups[device.SysfsID]
		if ok {
			slog.Debug("IOMMU group found", "device", device.SysfsID, "groupID", groupID)
			device.SysfsIOMMUGroupID = &groupID
		}

		err = report.Hardware.add(device)
		if err != nil {
			return nil, fmt.Errorf("failed to add to hardware report: %w", err)
		}
	}

	slog.Debug("processing smbios entries", "count", len(smbios))

	for idx := range smbios {
		err = report.Smbios.add(smbios[idx])
		if err != nil {
			return nil, fmt.Errorf("failed to add to smbios report: %w", err)
		}
	}

	slog.Debug("detecting virtualisation")

	report.Virtualisation, err = virt.Detect()
	if err != nil {
		return nil, fmt.Errorf("failed to detect virtualisation: %w", err)
	}

	slog.Debug("detecting UEFI")

	report.UEFI, err = boot.DetectUEFI()
	if err != nil {
		return nil, fmt.Errorf("failed to detect UEFI: %w", err)
	}

	if s.Ephemeral || s.Swap {
		slog.Debug("processing swap devices")

		report.Swap, err = ephem.SwapEntries()
		if err != nil {
			return nil, fmt.Errorf("failed to detect swap devices: %w", err)
		}
	}

	if s.Cloud.Hetzner {
		slog.Debug("fetching hetzner metadata")

		metadata, err := cloud.HetznerMetadata(context.Background(), cloud.HetznerMetadataURL)
		if err != nil {
			return nil, fmt.Errorf("failed to capture hetzner provider metadata: %w", err)
		}

		report.Cloud = &cloud.Cloud{Hetzner: metadata}
	}

	slog.Debug("report complete")

	return &report, nil
}
