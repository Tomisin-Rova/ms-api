// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type ActivateBioLoginResponse struct {
	Message           string `json:"message"`
	BiometricPasscode string `json:"biometricPasscode"`
}

type AuthResult struct {
	Token                  string `json:"token"`
	RefreshToken           string `json:"refreshToken"`
	RegistrationCheckpoint string `json:"registrationCheckpoint"`
}

type BioLoginInput struct {
	Email             string `json:"email"`
	BiometricPasscode string `json:"biometricPasscode"`
	DeviceID          string `json:"deviceId"`
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
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type CheckEmailExistenceResult struct {
	Exists  bool   `json:"exists"`
	Message string `json:"message"`
}

type CreateEmailInput struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Passcode string `json:"passcode"`
}

type CreatePasscodeInput struct {
	Token    string `json:"token"`
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

type InputAddress struct {
	Country  string `json:"country"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateBioDataInput struct {
	PersonID  string        `json:"personId"`
	Address   *InputAddress `json:"address"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Dob       string        `json:"dob"`
}
