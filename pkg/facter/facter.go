// Package facter contains types and utilities for scanning a system and generating a report, detailing key aspects of
// the system and its connected hardware.
package facter

import (
	"fmt"

	"github.com/numtide/nixos-facter/pkg/ephem"

	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/virt"
)

// Report represents a detailed report on the system’s hardware, virtualisation, SMBios, and swap entries.
type Report struct {
	// Version is a monotonically increasing number,
	// used to indicate breaking changes or new features in the report output.
	Version uint `json:"version"`

	// System indicates the system architecture e.g. x86_64-linux.
	System string `json:"system"`

	// Virtualisation indicates the type of virtualisation or container environment present on the system.
	Virtualisation virt.Type `json:"virtualisation"`

	// Hardware provides detailed information about the system’s hardware components, such as CPU, memory, and peripherals.
	Hardware Hardware `json:"hardware,omitempty"`

	// Smbios provides detailed information about the system's SMBios data, such as BIOS, board, chassis, memory, and processors.
	Smbios Smbios `json:"smbios,omitempty"`

	// Swap contains a list of swap entries representing the system's swap devices or files and their respective details.
	Swap []*ephem.SwapEntry `json:"swap,omitempty"`
}

// Scanner defines a type responsible for scanning and reporting system hardware information.
type Scanner struct {
	// Swap indicates whether the system swap information should be reported.
	Swap bool

	// Ephemeral indicates whether the scanner should report ephemeral details,
	// such as swap.
	Ephemeral bool

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
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	var smbios []hwinfo.Smbios
	var devices []hwinfo.HardwareDevice

	smbios, devices, err = hwinfo.Scan(s.Features)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	// read iommu groups
	iommuGroups, err := hwinfo.IOMMUGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to read iommu groups: %w", err)
	}

	for idx := range devices {
		// lookup iommu group before adding to the report
		device := devices[idx]
		groupId, ok := iommuGroups[device.SysfsId]
		if ok {
			device.SysfsIOMMUGroupId = &groupId
		}
		if err = report.Hardware.add(device); err != nil {
			return nil, fmt.Errorf("failed to add to hardware report: %w", err)
		}
	}

	for idx := range smbios {
		if err = report.Smbios.add(smbios[idx]); err != nil {
			return nil, fmt.Errorf("failed to add to smbios report: %w", err)
		}
	}

	if report.Virtualisation, err = virt.Detect(); err != nil {
		return nil, fmt.Errorf("failed to detect virtualisation: %w", err)
	}

	if s.Ephemeral || s.Swap {
		if report.Swap, err = ephem.SwapEntries(); err != nil {
			return nil, fmt.Errorf("failed to detect swap devices: %w", err)
		}
	}

	return &report, nil
}
