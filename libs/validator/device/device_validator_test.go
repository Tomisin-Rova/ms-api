package devicevalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"ms.api/types"
)

func TestValidate(t *testing.T) {
	const (
		failDeviceIsNil = iota
		failNoIdentifier
		failNoOs
		failNoBrand
		successValidDevice
	)
	testCases := []struct {
		name        string
		testType    int
		deviceInput *types.DeviceInput
		err         error
	}{
		{
			name:        "Fail device is nil",
			testType:    failDeviceIsNil,
			deviceInput: nil,
			err:         ErrInvalidDevice,
		},
		{
			name:     "Fail no identifier",
			testType: failNoIdentifier,
			deviceInput: &types.DeviceInput{
				Os:    "valid-os",
				Brand: "valid-brand",
			},
			err: ErrInvalidDevice,
		},
		{
			name:     "Fail no os",
			testType: failNoOs,
			deviceInput: &types.DeviceInput{
				Identifier: "valid-identifier",
				Brand:      "valid-brand",
			},
			err: ErrInvalidDevice,
		},
		{
			name:     "Fail no brand",
			testType: failNoBrand,
			deviceInput: &types.DeviceInput{
				Identifier: "valid-identifier",
				Os:         "valid-os",
			},
			err: ErrInvalidDevice,
		},
		{
			name:     "Success valid device",
			testType: successValidDevice,
			deviceInput: &types.DeviceInput{
				Identifier: "valid-identifier",
				Os:         "valid-os",
				Brand:      "valid-brand",
			},
			err: nil,
		},
	}

	validator := new(Validator)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failDeviceIsNil:
				err := validator.Validate(testCase.deviceInput)
				assert.Error(t, err)
				assert.IsType(t, ErrInvalidDevice, err)
			case failNoIdentifier:
				err := validator.Validate(testCase.deviceInput)
				assert.Error(t, err)
				assert.IsType(t, ErrInvalidDevice, err)
			case failNoOs:
				err := validator.Validate(testCase.deviceInput)
				assert.Error(t, err)
				assert.IsType(t, ErrInvalidDevice, err)
			case failNoBrand:
				err := validator.Validate(testCase.deviceInput)
				assert.Error(t, err)
				assert.IsType(t, ErrInvalidDevice, err)
			case successValidDevice:
				err := validator.Validate(testCase.deviceInput)
				assert.NoError(t, err)
				assert.Nil(t, err)
			}
		})
	}
}
