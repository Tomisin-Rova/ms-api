package devicevalidator

import (
	coreErrors "github.com/roava/zebra/errors"
	"ms.api/types"
)

var ErrInvalidDevice = coreErrors.NewTerror(
	1101,
	"InvalidDeviceError",
	"invalid device",
	"invaliddevice",
)

type DeviceValidator interface {
	Validate(deviceInput *types.DeviceInput) error
}

type Validator struct{}

func (v *Validator) Validate(deviceInput *types.DeviceInput) error {
	if deviceInput == nil {
		return ErrInvalidDevice
	}
	if len(deviceInput.Identifier) == 0 {
		return ErrInvalidDevice
	}
	if len(deviceInput.Os) == 0 {
		return ErrInvalidDevice
	}
	if len(deviceInput.Brand) == 0 {
		return ErrInvalidDevice
	}
	return nil
}
