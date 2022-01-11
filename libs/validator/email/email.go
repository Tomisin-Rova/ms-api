package emailvalidator

import (
	"regexp"
	"strings"

	errors2 "github.com/roava/zebra/errors"
)

var ErrInvalidEmail = errors2.NewTerror(
	1100, "InvalidEmailError", "invalid email address", "invalid email address",
)
var userRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
var hostRegexp = regexp.MustCompile(`^[^\s]+\.[^\s]+$`)

func Validate(email string) (string, error) {
	email = strings.TrimSpace(email)
	if len(email) < 6 || len(email) > 254 {
		return "", ErrInvalidEmail
	}

	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return "", ErrInvalidEmail
	}

	user := email[:at]
	host := email[at+1:]

	if len(user) > 64 {
		return "", ErrInvalidEmail
	}

	if !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return "", ErrInvalidEmail
	}
	return strings.ToLower(email), nil
}

type EmailValidator interface {
	Validate(email string) (string, error)
}

type Validator struct{}

func (v *Validator) Validate(email string) (string, error) {
	return Validate(email)
}
