package datevalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDob(t *testing.T) {
	type testCase struct {
		dob string
		err error
	}
	cases := []testCase{
		{dob: "12/10/1990", err: nil},
		{dob: "12/10/190", err: ErrInvalidFormat},
		{dob: "1a/10/1990", err: ErrInvalidFormat},
		{dob: "11/10/990", err: ErrInvalidFormat},
		{dob: "11/1a/1990", err: ErrInvalidFormat},
		{dob: "40/10/1990", err: ErrInvalidFormat},
		{dob: "40/10/1990", err: ErrInvalidFormat},
		{dob: "01/15/1990", err: ErrInvalidFormat},
		{dob: "01/11/19b0", err: ErrInvalidFormat},
		{dob: "01/01/2020", err: ErrInvalidAge},
		{dob: "01/01/2003", err: nil},
	}
	for idx, c := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			err := ValidateDob(c.dob)
			assert.Equal(t, c.err, err)
		})
	}
}
