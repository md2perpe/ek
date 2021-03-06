// +build !windows

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"time"

	PATH "pkg.re/essentialkaos/ek.v7/path"
	"pkg.re/essentialkaos/ek.v7/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_IFMT   = 0xf000
	_IFSOCK = 0xc000
	_IFLNK  = 0xa000
	_IFREG  = 0x8000
	_IFBLK  = 0x6000
	_IFDIR  = 0x4000
	_IFCHR  = 0x2000
	_IRUSR  = 0x100
	_IWUSR  = 0x80
	_IXUSR  = 0x40
	_IRGRP  = 0x20
	_IWGRP  = 0x10
	_IXGRP  = 0x8
	_IROTH  = 0x4
	_IWOTH  = 0x2
	_IXOTH  = 0x1
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrEmptyPath error
var ErrEmptyPath = errors.New("Path is empty")

// ////////////////////////////////////////////////////////////////////////////////// //

// CheckPerms check many props at once.
//
// F - is file
// D - is directory
// X - is executable
// L - is link
// W - is writable
// R - is readable
// S - not empty (only for files)
//
func CheckPerms(props, path string) bool {
	if len(props) == 0 || path == "" {
		return false
	}

	path = PATH.Clean(path)
	props = strings.ToUpper(props)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	var user *system.User

	for _, k := range props {
		switch k {

		case 'F':
			if stat.Mode&_IFMT != _IFREG {
				return false
			}

		case 'D':
			if stat.Mode&_IFMT != _IFDIR {
				return false
			}

		case 'L':
			if !IsLink(path) {
				return false
			}

		case 'X':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

			if !isExecutableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'W':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

			if !isWritableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'R':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

			if !isReadableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'S':
			if stat.Size == 0 {
				return false
			}
		}
	}

	return true
}

// ProperPath return first proper path from given slice
func ProperPath(props string, paths []string) string {
	for _, path := range paths {
		path = PATH.Clean(path)

		if CheckPerms(props, path) {
			return path
		}
	}

	return ""
}

// IsExist check if target is exist in fs or not
func IsExist(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return syscall.Access(path, syscall.F_OK) == nil
}

// IsRegular check if target is regular file or not
func IsRegular(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFREG
}

// IsSocket check if target is socket or not
func IsSocket(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFSOCK
}

// IsBlockDevice check if target is block device or not
func IsBlockDevice(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFBLK
}

// IsCharacterDevice check if target is character device or not
func IsCharacterDevice(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFCHR
}

// IsDir check if target is directory or not
func IsDir(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFDIR
}

// IsLink check if file is link or not
func IsLink(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var buf = make([]byte, 1)
	_, err := syscall.Readlink(path, buf)

	return err == nil
}

// IsReadable check if file is readable or not
func IsReadable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isReadableStat(stat, user.UID, getGIDList(user))
}

// IsWritable check if file is writable or not
func IsWritable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isWritableStat(stat, user.UID, getGIDList(user))
}

// IsExecutable check if file is executable or not
func IsExecutable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isExecutableStat(stat, user.UID, getGIDList(user))
}

// IsNonEmpty check if file is empty or not
func IsNonEmpty(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return GetSize(path) > 0
}

// IsEmptyDir check if directory empty or not
func IsEmptyDir(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	fd, err := syscall.Open(path, syscall.O_RDONLY, 0)

	if err != nil {
		return false
	}

	defer syscall.Close(fd)

	n, err := syscall.ReadDirent(fd, make([]byte, 4096))

	if n == 0x30 || err != nil {
		return true
	}

	return false
}

// GetOwner return object owner pid and gid
func GetOwner(path string) (int, int, error) {
	if path == "" {
		return -1, -1, ErrEmptyPath
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, -1, err
	}

	return int(stat.Uid), int(stat.Gid), nil
}

// GetATime return time of last access
func GetATime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	atime, _, _, err := GetTimes(path)

	return atime, err
}

// GetCTime return time of creation
func GetCTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, _, ctime, err := GetTimes(path)

	return ctime, err
}

// GetMTime return time of modification
func GetMTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, mtime, _, err := GetTimes(path)

	return mtime, err
}

// GetSize return file size in bytes
func GetSize(path string) int64 {
	if path == "" {
		return -1
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1
	}

	return stat.Size
}

// GetPerms return file permissions
func GetPerms(path string) os.FileMode {
	if path == "" {
		return 0
	}

	path = PATH.Clean(path)

	return os.FileMode(getMode(path) & 0777)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMode(path string) uint32 {
	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return 0
	}

	return uint32(stat.Mode)
}

func isReadableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if uid == 0 {
		return true
	}

	if stat.Mode&_IROTH == _IROTH {
		return true
	}

	if stat.Mode&_IRUSR == _IRUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IRGRP == _IRGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

func isWritableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if uid == 0 {
		return true
	}

	if stat.Mode&_IWOTH == _IWOTH {
		return true
	}

	if stat.Mode&_IWUSR == _IWUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IWGRP == _IWGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

func isExecutableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if uid == 0 {
		return true
	}

	if stat.Mode&_IXOTH == _IXOTH {
		return true
	}

	if stat.Mode&_IXUSR == _IXUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IXGRP == _IXGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

func getGIDList(user *system.User) []int {
	if user == nil {
		return nil
	}

	var result []int

	for _, group := range user.Groups {
		result = append(result, group.GID)
	}

	return result
}
