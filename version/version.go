// Package version provides methods for parsing semver version info
package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Version contains version data
type Version struct {
	raw        string
	slice      [3]int
	preRelease string
	build      string
	size       int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrEmpty           = errors.New("Version can't be empty")
	ErrEmptyBuild      = errors.New("Build number is empty")
	ErrEmptyPrerelease = errors.New("Prerelease number is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var preRegExp = regexp.MustCompile(`([a-zA-Z-.]{1,})([0-9]{0,})`)

// ////////////////////////////////////////////////////////////////////////////////// //

// Parse parse version string and return version struct
func Parse(v string) (Version, error) {
	if v == "" {
		return Version{}, ErrEmpty
	}

	var slice = [3]int{0, 0, 0}
	var raw = v

	var (
		preRelease string
		build      string
		size       int
	)

	if strings.Contains(v, "+") {
		bs := strings.Split(v, "+")

		if bs[1] == "" {
			return Version{}, ErrEmptyBuild
		} else {
			v = bs[0]
			build = bs[1]
		}
	}

	if strings.Contains(v, "-") {
		ps := strings.Split(v, "-")

		if ps[1] == "" {
			return Version{}, ErrEmptyPrerelease
		} else {
			v = ps[0]
			preRelease = ps[1]
		}
	}

	for index, version := range strings.Split(v, ".") {
		if index < 3 {
			iv, err := strconv.Atoi(version)

			if err != nil {
				return Version{}, err
			}

			slice[index] = iv
			size = index + 1
		}
	}

	return Version{
		raw:        raw,
		slice:      slice,
		preRelease: preRelease,
		build:      build,
		size:       size,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Major return major version
func (v Version) Major() int {
	if v.raw == "" || len(v.slice) == 0 {
		return -1
	}

	return v.slice[0]
}

// Minor return minor version
func (v Version) Minor() int {
	if v.raw == "" || len(v.slice) == 0 {
		return -1
	}

	return v.slice[1]
}

// Patch return patch version
func (v Version) Patch() int {
	if v.raw == "" || len(v.slice) == 0 {
		return -1
	}

	return v.slice[2]
}

// PreRelease return prerelease version
func (v Version) PreRelease() string {
	if v.raw == "" {
		return ""
	}

	return v.preRelease
}

// Build return build
func (v Version) Build() string {
	if v.raw == "" {
		return ""
	}

	return v.build
}

// Equal return true if version are equal to given
func (v Version) Equal(version Version) bool {
	if v.Major() != version.Major() {
		return false
	}

	if v.Minor() != version.Minor() {
		return false
	}

	if v.Patch() != version.Patch() {
		return false
	}

	if v.PreRelease() != version.PreRelease() {
		return false
	}

	if v.Build() != version.Build() {
		return false
	}

	return true
}

// Less return true if given version is greater
func (v Version) Less(version Version) bool {
	if v.Major() > version.Major() {
		return false
	}

	if v.Minor() > version.Minor() {
		return false
	}

	if v.Patch() > version.Patch() {
		return false
	}

	pr1, pr2 := v.PreRelease(), version.PreRelease()

	if pr1 != pr2 {
		return prereleaseLess(pr1, pr2)
	}

	if v.slice == version.slice {
		return false
	}

	return true
}

// Greater return true if given version is less
func (v Version) Greater(version Version) bool {
	if v.Major() < version.Major() {
		return false
	}

	if v.Minor() < version.Minor() {
		return false
	}

	if v.Patch() < version.Patch() {
		return false
	}

	pr1, pr2 := v.PreRelease(), version.PreRelease()

	if pr1 != pr2 {
		return !prereleaseLess(pr1, pr2)
	}

	if v.slice == version.slice {
		return false
	}

	return true
}

// Contains check is current version contains given version
func (v Version) Contains(version Version) bool {
	if v.Major() != version.Major() {
		return false
	}

	if v.size == 1 {
		return true
	}

	if v.Minor() != version.Minor() {
		return false
	}

	if v.size == 2 {
		return true
	}

	if v.Patch() != version.Patch() {
		return false
	}

	return false
}

// String return version as string
func (v Version) String() string {
	return v.raw
}

// ////////////////////////////////////////////////////////////////////////////////// //

// prereleaseLess
func prereleaseLess(pr1, pr2 string) bool {
	// Current version is release and given is prerelease
	if pr1 == "" && pr2 != "" {
		return false
	}

	// Current version is prerelease and given is release
	if pr1 != "" && pr2 == "" {
		return true
	}

	// Parse prerelease
	pr1Re := preRegExp.FindStringSubmatch(pr1)
	pr2Re := preRegExp.FindStringSubmatch(pr2)

	pr1Name := pr1Re[1]
	pr2Name := pr2Re[1]

	if pr1Name > pr2Name {
		return false
	}

	if pr1Name < pr2Name {
		return true
	}

	// Errors not important, because if subver is empty
	// Atoi return 0
	pr1Ver, _ := strconv.Atoi(pr1Re[2])
	pr2Ver, _ := strconv.Atoi(pr2Re[2])

	return pr1Ver < pr2Ver
}
