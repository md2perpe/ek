// +build !windows

package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"syscall"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize return window width and height
func GetSize() (int, int) {
	var tty *os.File

	tty, err := os.OpenFile("/dev/tty", syscall.O_RDONLY, 0)

	if err != nil {
		return -1, -1
	}

	defer tty.Close()

	var sz winsize

	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL, tty.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)),
	)

	return int(sz.cols), int(sz.rows)
}
