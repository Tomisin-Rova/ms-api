package validator

import (
	"regexp"

	coreError "github.com/roava/zebra/errors"
)

var (
	characterRegex                = regexp.MustCompile(`[A-Za-z]+`)
	digitRegex                    = regexp.MustCompile(`[\d]+`)
	ErrInvalidTransactionPassword = coreError.NewTerror(
		7010,
		"InvalidPassword",
		"Your transaction password must have at least one number and at least one letter and must be at least 8-characters long.",
		"",
	)
	ErrInvalidPayeeAccountDetailsCode    = 7011
	ErrInvalidPayeeAccountDetailsType    = "InvalidPayeeDetails"
	ErrInvalidPayeeAccountDetailsMessage = "Invalid payee account details"
	ErrInvalidPayeeAccountDetails        = coreError.NewTerror(
		ErrInvalidPayeeAccountDetailsCode,
		ErrInvalidPayeeAccountDetailsType,
		ErrInvalidPayeeAccountDetailsMessage,
		"",
	)
	ErrInvalidPaymentDetails = coreError.NewTerror(
		7012,
		"InvalidPaymentDetails",
		"Invalid payment details",
		"",
	)
)

func ValidateTransactionPassword(password string) error {
	if len(password) < 8 ||
		!characterRegex.MatchString(password) ||
		!digitRegex.MatchString(password) {
		return ErrInvalidTransactionPassword
	}

	return nil
}
