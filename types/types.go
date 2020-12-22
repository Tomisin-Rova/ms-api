// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type Account struct {
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currencySymbol"`
	AccountNumber  string `json:"accountNumber"`
	AccountName    string `json:"accountName"`
	Balance        string `json:"balance"`
}

type AccountsResult struct {
	PrimaryAccount   *Account   `json:"primaryAccount"`
	CurrencyAccounts []*Account `json:"currencyAccounts"`
}

type ActivateBioLoginResponse struct {
	Message           string `json:"message"`
	BiometricPasscode string `json:"biometricPasscode"`
}

type APIPerson struct {
	FirstName               string `json:"firstName"`
	LastName                string `json:"lastName"`
	Email                   string `json:"email"`
	IsEmailActive           bool   `json:"isEmailActive"`
	IsBiometricLoginEnabled bool   `json:"isBiometricLoginEnabled"`
	IsTransactionPinEnabled bool   `json:"isTransactionPinEnabled"`
	RegistrationCheckPoint  string `json:"registrationCheckPoint"`
}

type AuthResult struct {
	Token        string     `json:"token"`
	RefreshToken string     `json:"refreshToken"`
	Person       *APIPerson `json:"person"`
}

type AuthenticateCustomerInput struct {
	Email    string  `json:"email"`
	Passcode string  `json:"passcode"`
	Device   *Device `json:"device"`
}

type Beneficiary struct {
	PayeeID  string          `json:"payeeId"`
	Owner    string          `json:"owner"`
	Name     string          `json:"name"`
	Accounts []*PayeeAccount `json:"accounts"`
}

type BioLoginInput struct {
	Email             string  `json:"email"`
	BiometricPasscode string  `json:"biometricPasscode"`
	Device            *Device `json:"device"`
}

type Cdd struct {
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Details     string `json:"details"`
	Status      string `json:"status"`
	TimeCreated int64  `json:"time_created"`
	TimeUpdated int64  `json:"time_updated"`
}

type CDDSummary struct {
	Status    string                `json:"status"`
	Documents []*CDDSummaryDocument `json:"documents"`
}

type CDDSummaryDocument struct {
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Reasons []string `json:"reasons"`
}

type CheckEmailExistenceResult struct {
	Exists  bool   `json:"exists"`
	Message string `json:"message"`
}

type Country struct {
	CountryID                     string `json:"CountryId"`
	Capital                       string `json:"Capital"`
	CountryName                   string `json:"CountryName"`
	Continent                     string `json:"Continent"`
	Dial                          string `json:"Dial"`
	GeoNameID                     string `json:"GeoNameId"`
	ISO4217CurrencyAlphabeticCode string `json:"ISO4217CurrencyAlphabeticCode"`
	ISO4217CurrencyNumericCode    int64  `json:"ISO4217CurrencyNumericCode"`
	IsIndependent                 string `json:"IsIndependent"`
	Languages                     string `json:"Languages"`
	OfficialNameEnglish           string `json:"officialNameEnglish"`
}

type CreatePasscodeInput struct {
	Token    string `json:"token"`
	Passcode string `json:"passcode"`
}

type CreatePayeeInput struct {
	Name           string  `json:"name"`
	AccountNumber  string  `json:"accountNumber"`
	SortCode       *string `json:"sortCode"`
	BankCode       *string `json:"bankCode"`
	RoutingNumber  *string `json:"RoutingNumber"`
	Bic            *string `json:"BIC"`
	Iban           *string `json:"IBAN"`
	Country        string  `json:"country"`
	BankName       *string `json:"bankName"`
	RoutingType    *string `json:"routingType"`
	TransactionPin string  `json:"transactionPin"`
}

type CreatePayeeResult struct {
	Message     string       `json:"message"`
	Beneficiary *Beneficiary `json:"Beneficiary"`
}

type CreatePersonInput struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Passcode string `json:"passcode"`
}

type CreatePhoneInput struct {
	Phone  string  `json:"phone"`
	Device *Device `json:"device"`
}

type CreatePhoneResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type DeactivateBioLoginInput struct {
	Email    string `json:"email"`
	DeviceID string `json:"deviceId"`
}

type Device struct {
	Os          string `json:"os"`
	Brand       string `json:"brand"`
	DeviceID    string `json:"deviceId"`
	DeviceToken string `json:"deviceToken"`
}

type GetPayeesByPhoneNumbers struct {
	Payees []*Payee `json:"payees"`
}

type InputAddress struct {
	Country  string `json:"country"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
}

type MakeTransferInput struct {
	FromAccountNumber string `json:"fromAccountNumber"`
	ToAccountNumber   string `json:"toAccountNumber"`
	Amount            int64  `json:"amount"`
	Notes             string `json:"notes"`
	TransactionPin    string `json:"transactionPin"`
}

type Payee struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	PersonID    string `json:"personId"`
}

type PayeeAccount struct {
	AccountNumber string  `json:"accountNumber"`
	SortCode      *string `json:"sortCode"`
	BankCode      *string `json:"bankCode"`
	RoutingNumber *string `json:"routingNumber"`
	Bic           *string `json:"BIC"`
	Iban          *string `json:"IBAN"`
	Country       string  `json:"country"`
	BankName      *string `json:"bankName"`
	RoutingType   *string `json:"routingType"`
}

type Person struct {
	PhoneNumber             string           `json:"phoneNumber"`
	FirstName               string           `json:"firstName"`
	LastName                string           `json:"lastName"`
	MiddleName              string           `json:"middleName"`
	Email                   string           `json:"email"`
	Nationality             string           `json:"nationality"`
	Addresses               []*PersonAddress `json:"addresses"`
	Dob                     string           `json:"dob"`
	IsEmailActive           bool             `json:"isEmailActive"`
	IsBiometricLoginEnabled bool             `json:"isBiometricLoginEnabled"`
	IsTransactionPinEnabled bool             `json:"isTransactionPinEnabled"`
	RegistrationCheckPoint  string           `json:"registrationCheckPoint"`
}

type PersonAddress struct {
	Country  string `json:"country"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
}

type Reason struct {
	ID          string `json:"Id"`
	Description string `json:"Description"`
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateBioDataInput struct {
	Address   *InputAddress `json:"address"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Dob       string        `json:"dob"`
}

type VerifyEmailMagicLInkInput struct {
	Email             string `json:"email"`
	VerificationToken string `json:"verificationToken"`
}

type CreateApplicationResponse struct {
	Token string `json:"token"`
}

type FetchCountriesResponse struct {
	Countries []*Country `json:"Countries"`
}

type FetchReasonResponse struct {
	Reasons []*Reason `json:"reasons"`
}
