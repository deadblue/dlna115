//go:build windows
// +build windows

package util

import (
	"log"

	"golang.org/x/sys/windows"
)

func socketControl(fd uintptr) {
	if err := windows.SetsockoptInt(
		windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1,
	); err != nil {
		log.Printf("Set socket reuse_addr failed: %s", err)
	}
}
