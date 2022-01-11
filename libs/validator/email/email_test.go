package emailvalidator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	type testCase struct {
		email string
		err   error
	}
	cases := []testCase{
		{email: "foo@bar.com", err: nil},
		{email: "foo@bar.", err: ErrInvalidEmail},
		{email: "@bar.com", err: ErrInvalidEmail},
		{email: "s@s.c", err: ErrInvalidEmail},
		{email: strings.Repeat("u", 65) + "@email.com", err: ErrInvalidEmail},
	}
	validator := new(Validator)
	for _, c := range cases {
		_, err := validator.Validate(c.email)
		assert.Equal(t, c.err, err)
	}
}
