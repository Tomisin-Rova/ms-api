package mapper

import (
	"errors"

	coreError "github.com/roava/zebra/errors"
	zaplogger "github.com/roava/zebra/logger"
	"go.uber.org/zap"
)

type Mapper interface {
	Hydrate(from interface{}, to interface{}) error
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
