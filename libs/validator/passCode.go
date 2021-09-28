package validator

import (
	coreErrors "github.com/roava/zebra/errors"
	"regexp"
	"strings"
)

var (
	passCodeRegex      = regexp.MustCompile(`^\d{6}$`)
	errInvalidPassCode = coreErrors.NewTerror(
		7011, "InvalidPassCode", "invalid pass code", "invalid pass code",
	)
)

func IsValidPassCode(passCode string) error {
	passCode = strings.TrimSpace(passCode)
	if ok := passCodeRegex.MatchString(passCode); !ok {
		return errInvalidPassCode
	}

	return nil
}
