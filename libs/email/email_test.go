package emailvalidator

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	}
	for _, c := range cases {
		err := Validate(c.email)
		assert.Equal(t, c.err, err)
	}
}
