package validator

import (
	"regexp"

	coreError "github.com/roava/zebra/errors"
	"ms.api/types"
)

const (
	ngnCurrency = "NGN"
	gbpCurrency = "GBP"
)

var (
	characterRegex                = regexp.MustCompile(`[A-Za-z]+`)
	digitRegex                    = regexp.MustCompile(`[\d]+`)
	ngnAccountNumberRegex         = regexp.MustCompile(`^\d{10}$`)
	gbpAccountNumberRegex         = regexp.MustCompile(`^\d{8}$`)
	sortCodeRegex                 = regexp.MustCompile(`^\d{6}$`)
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
		switch *p.Currency {
		case ngnCurrency:
			if !ngnAccountNumberRegex.MatchString(*p.AccountNumber) {
				return nil, coreError.NewTerror(
					ErrInvalidPayeeAccountDetailsCode,
					ErrInvalidPayeeAccountDetailsType,
					ErrInvalidPayeeAccountDetailsMessage,
					"",
					coreError.WithHelp("account number must be a 10 digit value"),
				)
			}
		default:
			if !gbpAccountNumberRegex.MatchString(*p.AccountNumber) {
				return nil, coreError.NewTerror(
					ErrInvalidPayeeAccountDetailsCode,
					ErrInvalidPayeeAccountDetailsType,
					ErrInvalidPayeeAccountDetailsMessage,
					"",
					coreError.WithHelp("account number must be an 8 digit value"),
				)
			}
		}
		payeeAccount.AccountNumber = *p.AccountNumber
	}
	if p.SortCode != nil {
		switch *p.Currency {
		case gbpCurrency:
			if !sortCodeRegex.MatchString(*p.SortCode) {
				return nil, coreError.NewTerror(
					ErrInvalidPayeeAccountDetailsCode,
					ErrInvalidPayeeAccountDetailsType,
					ErrInvalidPayeeAccountDetailsMessage,
					"",
					coreError.WithHelp("sort code must be a 6 digit value"),
				)
			}
		}
		payeeAccount.SortCode = *p.SortCode
	}
	if p.Iban != nil {
		payeeAccount.Iban = *p.Iban
	}
	if p.SwiftBic != nil {
		payeeAccount.SwiftBic = *p.SwiftBic
	}
	if *p.Currency == ngnCurrency && (p.BankCode == nil || *p.BankCode == "") {
		return nil, coreError.NewTerror(
			ErrInvalidPayeeAccountDetailsCode,
			ErrInvalidPayeeAccountDetailsType,
			ErrInvalidPayeeAccountDetailsMessage,
			"",
			coreError.WithHelp("bank code is required for NGN account"),
		)
	}
	if p.BankCode != nil {
		payeeAccount.BankCode = *p.BankCode
	}
	if p.RoutingNumber != nil {
		payeeAccount.RoutingNumber = *p.RoutingNumber
	}

	return payeeAccount, nil
}

func ValidatePayment(p *types.PaymentInput) (*types.PaymentInput, error) {
	if p.Beneficiary == nil {
		return nil, ErrInvalidPaymentDetails
	}

	return p, nil
}
