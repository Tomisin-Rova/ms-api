package validator

import (
	coreError "github.com/roava/zebra/errors"
	"regexp"
)

var (
	alphaNumericRegex             = regexp.MustCompile(`^([A-Za-z]+[0-9]|[0-9]+[A-Za-z])[A-Za-z0-9]*$`)
	ErrInvalidTransactionPassword = coreError.NewTerror(
		7010,
		"InvalidPassword",
		"8 digit password must have at least one alphabet and one letter",
		"",
	)
)

func ValidateTransactionPassword(password string) error {
	if ok := alphaNumericRegex.MatchString(password); !ok {
		return ErrInvalidTransactionPassword
	}
	if len(password) != 8 {
		return ErrInvalidTransactionPassword
	}
	return nil
}
