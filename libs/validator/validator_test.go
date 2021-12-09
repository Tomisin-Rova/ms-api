package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

//stringP - return string pointer
// func stringP(s string) *string {
// 	return &s
// }
