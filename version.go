package main

import (
	"fmt"
)

// Version information - will be overridden during build
var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

// GetVersion returns the version
func GetVersion() string {
	return Version
}

// GetVersionInfo returns the full version information
func GetVersionInfo() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", Version, CommitSHA, BuildDate)
}
