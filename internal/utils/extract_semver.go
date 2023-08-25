package utils

import (
	"errors"
	"regexp"
)

var (
	semverRegex = regexp.MustCompile(`v?(\d+\.\d+\.\d+)`)
	ErrNoSemver = errors.New("could not find semver in version")
)

func ExtractSemver(version string) (string, error) {
	matches := semverRegex.FindStringSubmatch(version)

	if len(matches) < 2 {
		return "", ErrNoSemver
	}

	return matches[1], nil
}
