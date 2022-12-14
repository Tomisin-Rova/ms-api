package phonenumbervalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePhoneNumber(t *testing.T) {
	type testCase struct {
		phoneNumber string
		err         error
	}
	cases := []testCase{
		{phoneNumber: "+2347035452307", err: nil},
		{phoneNumber: "07035452307", err: nil},
		{phoneNumber: "+2347035452307a", err: ErrInvalidPhoneNumber},
		{phoneNumber: "+adftoio", err: ErrInvalidPhoneNumber},
		{phoneNumber: "+23470a5452307", err: ErrInvalidPhoneNumber},
		{phoneNumber: "+234703545230745", err: ErrInvalidPhoneNumber},
		{phoneNumber: "+23490", err: ErrInvalidPhoneNumber},
	}
	validator := new(Validator)
	for idx, next := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			err := validator.ValidatePhoneNumber(next.phoneNumber)
			assert.Equal(t, err, next.err)
		})
	}
}
