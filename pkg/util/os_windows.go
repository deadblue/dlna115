//go:build windows
// +build windows

package util

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func getOsVersion() string {
	ver := windows.RtlGetVersion()
	return fmt.Sprintf(
		"Windows/%d.%d.%d",
		ver.MajorVersion, ver.MinorVersion, ver.BuildNumber,
	)
}
