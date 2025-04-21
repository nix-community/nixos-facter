package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/numtide/nixos-facter/pkg/udev"
	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/facter"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

var (
	outputPath       string
	logLevel         string
	hardwareFeatures []string
	version          bool

	scanner = facter.Scanner{}
)

func init() {
	// Define flags
	flag.StringVar(&outputPath, "output", "", "path to write the report")
	flag.StringVar(&outputPath, "o", "", "path to write the report")
	flag.BoolVar(&scanner.Swap, "swap", false, "capture swap entries")
	flag.BoolVar(
		&scanner.Ephemeral, "ephemeral", false,
		"capture all ephemeral properties e.g. swap, filesystems and so on",
	)
	flag.BoolVar(
		&version, "version", false,
		"print version and exit",
	)
	flag.StringVar(&logLevel, "log-level", "info", "log level")

	defaultFeatures := []string{
		"memory", "pci", "net", "serial", "cpu", "bios", "monitor", "scsi", "usb", "prom", "sbus", "sys", "sysfs",
		"udev", "block", "wlan",
	}

	var filteredFeatures []string

	for _, feature := range hwinfo.ProbeFeatureStrings() {
		if feature != "default" && feature != "int" {
			filteredFeatures = append(filteredFeatures, feature)
		}
	}

	hardwareFeatures = defaultFeatures

	flag.Func("hardware", "Hardware items to probe (comma separated).", func(flagValue string) error {
		hardwareFeatures = strings.Split(flagValue, ",")
		return nil
	})

	possibleValues := strings.Join(filteredFeatures, ",")
	defaultValues := strings.Join(defaultFeatures, ",")

	const usage = `nixos-facter [flags]
Hardware report generator %s (%s)

Usage:
  nixos-facter [flags]

Flags:
  --ephemeral          capture all ephemeral properties e.g. swap, filesystems and so on
  -h, --help           help for nixos-facter
  -o, --output string  path to write the report
  --swap               capture swap entries
  --version            version for nixos-facter
  --log-level string   log level, one of <error|warn|info|debug> (default "info")
  --hardware strings   Hardware items to probe.
                       Default: %s
                       Possible values: %s

`

	// Custom usage function
	flag.Usage = func() { fmt.Fprintf(os.Stderr, usage, build.Version, build.System, defaultValues, possibleValues) }
}

func Execute() {
	// check udev version
	if udevVersion, err := udev.Version(); err != nil {
		log.Fatalf("failed to get systemd version: %v", err)
	} else if udevVersion < 252 {
		log.Fatalf("udev version %d is too old, please upgrade to at least 252", udevVersion)
	}

	flag.Parse()

	if version {
		fmt.Printf("%s\n", build.Version)
		return
	}

	// Check if the effective user id is 0 e.g. root
	if os.Geteuid() != 0 {
		log.Fatalf("you must run this program as root")
	}

	// Convert the hardware features into probe features
	for _, str := range hardwareFeatures {
		probe, err := hwinfo.ProbeFeatureString(str)
		if err != nil {
			log.Fatalf("invalid hardware feature: %v", err)
		}

		scanner.Features = append(scanner.Features, probe)
	}

	// Set the log level

	var slogLevel slog.Level
	if err := slogLevel.UnmarshalText([]byte(logLevel)); err != nil {
		log.Fatalf("invalid log level: %v", err)
	}

	switch slogLevel {
	case slog.LevelDebug:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	case slog.LevelInfo:
		log.SetFlags(log.LstdFlags)
	case slog.LevelWarn, slog.LevelError:
		log.SetFlags(0)
	default:
		log.Fatalf("invalid log level: %s", logLevel)
	}

	slog.SetLogLoggerLevel(slogLevel)

	report, err := scanner.Scan()
	if err != nil {
		log.Fatalf("failed to scan: %v", err)
	}

	bytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal report to json: %v", err)
	}

	// If a file path is provided write the report to it, otherwise output the report on stdout
	if outputPath == "" {
		if _, err = os.Stdout.Write(bytes); err != nil {
			log.Fatalf("failed to write report to stdout: %v", err)
		}

		fmt.Println()
	} else if err = os.WriteFile(outputPath, bytes, 0o600); err != nil {
		log.Fatalf("failed to write report to output path: %v", err)
	}
}
