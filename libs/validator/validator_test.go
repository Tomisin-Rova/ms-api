package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ms.api/types"
	"testing"
)

func TestValidateTransactionPassword(t *testing.T) {
	type tests struct {
		password string
		err      error
	}
	cases := []tests{
		{password: "abcd1234", err: nil},
		{password: "abcdabc1", err: nil},
		{password: "a1234567", err: nil},
		{password: "1234567a", err: nil},
		{password: "1111111111A", err: nil},
		{password: "1111&%/111A", err: nil},
		{password: "abcd123", err: ErrInvalidTransactionPassword},
		{password: "abcd", err: ErrInvalidTransactionPassword},
		{password: "abcdefgh", err: ErrInvalidTransactionPassword},
		{password: "77777777", err: ErrInvalidTransactionPassword},
	}
	for idx, next := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			err := ValidateTransactionPassword(next.password)
			assert.Equal(t, err, next.err)
		})
	}
}

func TestValidatePayeeAccount(t *testing.T) {
	type tests struct {
		payeeAccount *types.PayeeAccountInput
		err          error
	}

	cases := []tests{
		{
			payeeAccount: &types.PayeeAccountInput{
				Iban:          stringP("123414"),
				SortCode:      stringP("12345678"),
				AccountNumber: stringP("1223112422"),
			},
			err: nil,
		},
		{
			payeeAccount: &types.PayeeAccountInput{
				SortCode: stringP("12345678"),
			},
			err: nil,
		},
		{
			payeeAccount: &types.PayeeAccountInput{
				SortCode:      stringP("12345678aa"),
				AccountNumber: stringP("1223112422"),
			},
			err: ErrInvalidPayeeAccountDetails,
		},
		{
			payeeAccount: &types.PayeeAccountInput{
				SortCode:      stringP("12345678"),
				AccountNumber: stringP("some12422"),
			},
			err: ErrInvalidPayeeAccountDetails,
		},
	}

	for idx, next := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			_, err := ValidatePayeeAccount(next.payeeAccount)
			assert.Equal(t, next.err, err)
		})
	}
}

//stringP - return string pointer
func stringP(s string) *string {
	return &s
}
