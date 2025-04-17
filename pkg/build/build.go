// Package build contains constants and values set at build time via -X flags.
package build

var (
	// Name is the program name, typically set via Nix to match the derivation's `pname`.
	Name = "nixos-facter"
	// Version is the program version, typically set via Nix to match the derivation's `version`.
	Version = "v0.0.0+dev"
	// System is the architecture that this program was built for e.g. x86_64-linux.
	// It is set via Nix to match the Nixpkgs system.
	System = ""
	// ReportVersion is used to indicate significant changes in the report output and is embedded JSON report produced.
	ReportVersion uint16 = 1
)
