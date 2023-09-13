//go:build unix
// +build unix

package util

import (
	"log"

	"golang.org/x/sys/unix"
)

func socketControl(fd uintptr) {
	var err error
	if err = unix.SetsockoptInt(
		int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1,
	); err != nil {
		log.Printf("Set socket reuse_addr failed: %s", err)
		return
	}
	if err = unix.SetsockoptInt(
		int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1,
	); err != nil {
		log.Printf("Set socket reuse_port failed: %s", err)
	}
}
