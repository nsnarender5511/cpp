package version

// Version is the current version of cursor++.
// This is overridden during build by ldflags
var Version = "dev"

// GetVersion returns the current version
func GetVersion() string {
	return Version
}
