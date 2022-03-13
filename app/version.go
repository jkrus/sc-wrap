package app

import "sc-wrap/config"

// version returns the application version as a properly formed string
// per the semantic versioning 2.0.0 spec (http://semver.org/).
func version() string {
	// start with the major, minor, and patch versions.
	ver := config.AppVersion
	// append pre-release if there is one.
	if config.AppVerPreRelease {
		ver += "-dev"
	}

	return ver
}
