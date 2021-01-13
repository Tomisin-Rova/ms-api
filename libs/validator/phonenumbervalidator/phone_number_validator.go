package phonenumbervalidator

import (
	"regexp"

	coreError "github.com/roava/zebra/errors"
)

var (
	phoneRegex            = regexp.MustCompile(`^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$`)
	ErrInvalidPhoneNumber = coreError.NewTerror(
		7001,
		"InvalidPhoneNumberException",
		"phone number is not valid",
		"",
	)
)

func ValidatePhoneNumber(phoneNumber string) error {
	if ok := phoneRegex.MatchString(phoneNumber); !ok {
		return ErrInvalidPhoneNumber
	}
	if len(phoneNumber) > 15 || len(phoneNumber) < 7 {
		return ErrInvalidPhoneNumber
	}
	return nil
}
