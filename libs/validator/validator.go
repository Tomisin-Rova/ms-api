package validator

import (
	"regexp"

	coreError "github.com/roava/zebra/errors"
	"ms.api/types"
)

var (
	characterRegex                = regexp.MustCompile(`[A-Za-z]+`)
	digitRegex                    = regexp.MustCompile(`[\d]+`)
	onlyDiigtsRegex               = regexp.MustCompile(`^[0-9]*$`)
	ErrInvalidTransactionPassword = coreError.NewTerror(
		7010,
		"InvalidPassword",
		"Your transaction password must have at least one number and at least one letter and must be at least 8-characters long.",
		"",
	)
	ErrInvalidPayeeAccountDetails = coreError.NewTerror(
		7011,
		"InvalidPayeeDetails",
		"Invalid payee account details",
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

// ValidatePayeeAccount validate and serialize payee account information
func ValidatePayeeAccount(p *types.PayeeAccountInput) (*types.PayeeAccountInfo, error) {
	payeeAccount := &types.PayeeAccountInfo{}
	if p.Name != nil {
		payeeAccount.Name = *p.Name
	}
	if p.Currency == nil || *p.Currency == "" {
		return nil, ErrInvalidPayeeAccountDetails
	}
	payeeAccount.Currency = *p.Currency
	if p.AccountNumber != nil {
		payeeAccount.AccountNumber = *p.AccountNumber
	}
	if p.SortCode != nil {
		payeeAccount.SortCode = *p.SortCode
	}
	if p.Iban != nil {
		payeeAccount.Iban = *p.Iban
	}
	if p.SwiftBic != nil {
		payeeAccount.SwiftBic = *p.SwiftBic
	}
	if p.BankCode != nil {
		payeeAccount.BankCode = *p.BankCode
	}
	if p.RoutingNumber != nil {
		payeeAccount.RoutingNumber = *p.RoutingNumber
	}

	if (!onlyDiigtsRegex.MatchString(payeeAccount.SortCode) && payeeAccount.SortCode != "") ||
		(!onlyDiigtsRegex.MatchString(payeeAccount.AccountNumber) && payeeAccount.AccountNumber != "") {
		return nil, ErrInvalidPayeeAccountDetails
	}

	return payeeAccount, nil
}

func ValidatePayment(p *types.PaymentInput) (*types.PaymentInput, error) {
	if p.Beneficiary == nil ||
		p.Beneficiary.Amount == nil {
		return nil, ErrInvalidPaymentDetails
	}

	return p, nil
}
