package mapper

import (
	"errors"

	coreError "github.com/roava/zebra/errors"
	zaplogger "github.com/roava/zebra/logger"
	"go.uber.org/zap"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/types"
)

type Mapper interface {
	Hydrate(from interface{}, to interface{}) error
	DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) pbTypes.DeviceToken_DeviceTokenTypes
	PreferenceInputFromModel(input types.DevicePreferencesTypes) pbTypes.DevicePreferences_DevicePreferencesTypes
}

// GQLMapper a mapper that returns Graphql types
type GQLMapper struct {
	logger *zap.Logger
}

var _ Mapper = &GQLMapper{
	logger: zaplogger.New(),
}

// Hydrate converts between types
func (G *GQLMapper) Hydrate(from interface{}, to interface{}) error {

	return errors.New("could not handle type")

}

var (
	MappingErr = coreError.NewTerror(
		7021,
		"InternalError",
		"failed to process the request, please try again later.",
		"",
	)
)

func String(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Int64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func Float64(i float64) *float64 {
	if i == 0 {
		return nil
	}
	return &i
}

func NewMapper() *GQLMapper {
	return &GQLMapper{
		logger: zaplogger.New(),
	}
}

func (G *GQLMapper) DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) pbTypes.DeviceToken_DeviceTokenTypes {
	switch tokenType {
	default:
		return pbTypes.DeviceToken_FIREBASE
	}
}

func (G *GQLMapper) PreferenceInputFromModel(input types.DevicePreferencesTypes) pbTypes.DevicePreferences_DevicePreferencesTypes {
	switch input {
	case types.DevicePreferencesTypesPush:
		return pbTypes.DevicePreferences_PUSH
	case types.DevicePreferencesTypesBiometrics:
		return pbTypes.DevicePreferences_BIOMETRICS
	default:
		return pbTypes.DevicePreferences_PUSH
	}
}
