package phonenumbervalidator

import (
	"regexp"

	coreError "github.com/roava/zebra/errors"
)

var (
	phoneRegex            = regexp.MustCompile(`^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$`)
	ErrInvalidPhoneNumber = coreError.NewTerror(
		7010,
		"InvalidPhoneNumberException",
		"phone number is not valid",
		"",
	)
)

type PhoneNumberValidator interface {
	ValidatePhoneNumber(phoneNumber string) error
}

type Validator struct{}

func (v *Validator) ValidatePhoneNumber(phoneNumber string) error {
	if len(phoneNumber) > 15 || len(phoneNumber) < 7 {
		return ErrInvalidPhoneNumber
	}
	if ok := phoneRegex.MatchString(phoneNumber); !ok {
		return ErrInvalidPhoneNumber
	}
	return nil
}
