// +build freebsd

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getTimes is copy of fsutil.GetTimes
func getTimes(path string) (time.Time, time.Time, time.Time, error) {
	if path == "" {
		return time.Time{}, time.Time{}, time.Time{}, errors.New("Path is empty")
	}

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	return time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec)),
		time.Unix(int64(stat.Mtimespec.Sec), int64(stat.Mtimespec.Nsec)),
		time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec)),
		nil
}

// getUserInfo find user info by name or id
func getUserInfo(nameOrID string) (*User, error) {
	cmd := exec.Command("getent", "passwd", nameOrID)

	out, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("User with this name/id %s does not exist", nameOrID)
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")
	aOut := strings.Split(sOut, ":")

	uid, _ := strconv.Atoi(aOut[2])
	gid, _ := strconv.Atoi(aOut[3])

	return &User{
		Name:     aOut[0],
		UID:      uid,
		GID:      gid,
		Comment:  aOut[4],
		HomeDir:  aOut[5],
		Shell:    aOut[6],
		RealName: aOut[0],
		RealUID:  uid,
		RealGID:  gid,
	}, nil
}
