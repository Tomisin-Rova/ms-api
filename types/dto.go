package types

type PayeeAccountInfo struct {
	Name          string `json:"name"`
	Currency      string `json:"currency"`
	AccountNumber string `json:"account_number"`
	SortCode      string `json:"sort_code"`
	Iban          string `json:"iban"`
	SwiftBic      string `json:"swift_bic"`
	BankCode      string `json:"bank_code"`
	RoutingNumber string `json:"routing_number"`
	PhoneNumber   string `json:"phone_number"`
}
