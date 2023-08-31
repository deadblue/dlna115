//go:build unix
// +build unix

package util

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func getOsVersion() string {
	buf := &unix.Utsname{}
	unix.Uname(buf)
	return fmt.Sprintf(
		"%s/%s",
		unix.ByteSliceToString(buf.Sysname[:]),
		unix.ByteSliceToString(buf.Release[:]),
	)
}
