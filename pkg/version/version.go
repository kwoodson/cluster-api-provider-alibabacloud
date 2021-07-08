package version

import (
	"strings"

	"github.com/blang/semver"
)

var (
	// Raw is the string representation of the version. This will be replaced
	// with the calculated version at build time.
	Raw = "v0.1.0-alpha.1"

	// Version is semver representation of the version.
	Version = semver.MustParse(strings.TrimLeft(Raw, "v"))
)
