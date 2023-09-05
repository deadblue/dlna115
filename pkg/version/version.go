package version

import (
	"fmt"
	"runtime"
)

const (
	name = "DLNA115"
)

var (
	version string

	shortVer = ""
	fullVer  = ""
)

func init() {
	shortVer = fmt.Sprintf("%s/%s", name, version)

	fullVer = fmt.Sprintf(
		"%s %s (%s %s/%s)",
		name, version, runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)
}

func Short() string {
	return shortVer
}

func Full() string {
	return fullVer
}

func Version() string {
	return version
}
