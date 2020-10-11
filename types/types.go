// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type AuthResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
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
