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
	if p.Currency != nil {
		payeeAccount.Currency = *p.Currency
	}
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
	payment := &types.PaymentInput{}
	if p.IdempotencyKey != " " {
		payment.IdempotencyKey = p.IdempotencyKey
	}
	if p.Owner != " " {
		payment.Owner = p.Owner
	}
	if p.Charge != nil {
		payment.Charge = p.Charge
	}
	if p.Currency != nil {
		payment.Currency = p.Currency
	}
	if p.Reference != nil {
		payment.Reference = p.Reference
	}
	if p.Status != nil {
		payment.Status = p.Status
	}
	if p.Image != nil {
		payment.Image = p.Image
	}
	if p.Notes != nil {
		payment.Notes = p.Notes
	}
	if p.Quote != nil {
		payment.Quote = p.Quote
	}
	if p.Tags != nil {
		payment.Tags = p.Tags
	}
	if p.Beneficiary != nil {
		payment.Beneficiary = p.Beneficiary
	}
	if p.FundingSource != " " {
		payment.FundingSource = p.FundingSource
	}
	if p.Currency != nil {
		payment.Currency = p.Currency
	}
	if p.FundingAmount != 0 {
		payment.FundingAmount = p.FundingAmount
	}
	return payment, nil
}
