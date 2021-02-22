package validator

import (
	"github.com/roava/zebra/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPassCode(t *testing.T) {
	const (
		success                 = iota
		invalid                 = iota
		invalidWithAlphanumeric = iota
	)

	var tests = []struct {
		name     string
		passCode string
		testType int
	}{
		{
			name:     "Test valid pass code",
			passCode: "123456",
			testType: success,
		},
		{
			name:     "Test invalid pass code",
			passCode: "12345",
			testType: invalid,
		},
		{
			name:     "Test invalid pass code with alphanumeric",
			passCode: "abcdef",
			testType: invalidWithAlphanumeric,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case success:
				err := IsValidPassCode(testCase.passCode)

				assert.NoError(t, err)
			case invalid:
				err := IsValidPassCode(testCase.passCode)

				assert.Error(t, err)
				assert.IsType(t, &errors.Terror{}, err)
				assert.Equal(t, 7011, err.(*errors.Terror).Code())
			case invalidWithAlphanumeric:
				err := IsValidPassCode(testCase.passCode)

				assert.Error(t, err)
				assert.IsType(t, &errors.Terror{}, err)
				assert.Equal(t, 7011, err.(*errors.Terror).Code())
			}
		})
	}
}
