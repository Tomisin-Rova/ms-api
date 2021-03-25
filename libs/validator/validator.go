package validator

import (
	coreError "github.com/roava/zebra/errors"
	"ms.api/types"
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
	if len(password) < 8 {
		return ErrInvalidTransactionPassword
	}
	return nil
}

// ValidatePayeeAccount validate and serialize payee account information
// TODO: add validation cases when it's added to the AC
func ValidatePayeeAccount(p *types.PayeeAccountInput) (*types.PayeeAccountInfo, error) {
	payeeAccount := &types.PayeeAccountInfo{
		Name: "",
		Currency: "",
		AccountNumber: "",
		SortCode: "",
		Iban: "",
		SwiftBic: "",
		BankCode: "",
		RoutingNumber: "",
		PhoneNumber: "",
	}
	if p.Name != nil { payeeAccount.Name = *p.Name }
	if p.Currency != nil { payeeAccount.Currency = *p.Currency }
	if p.AccountNumber != nil { payeeAccount.AccountNumber = *p.AccountNumber }
	if p.SortCode != nil { payeeAccount.SortCode = *p.SortCode }
	if p.Iban != nil { payeeAccount.Iban = *p.Iban }
	if p.SwiftBic != nil { payeeAccount.SwiftBic = *p.SwiftBic }
	if p.BankCode != nil { payeeAccount.BankCode = *p.BankCode }
	if p.RoutingNumber != nil { payeeAccount.RoutingNumber = *p.RoutingNumber }

	return payeeAccount, nil
}
