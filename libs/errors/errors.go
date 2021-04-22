package errors

import (
	coreError "github.com/roava/zebra/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	CddNotFound = 1105
)

// NewFromGrpc returns a formatted Roava Terror
func NewFromGrpc(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		if err == nil {
			return coreError.NewTerror(
				7000,
				"InvalidException",
				"unknown error",
				"",
			)
		}
		return coreError.NewTerror(
			7002,
			"ConnectionError",
			"error connecting to service",
			err.Error(),
		)
	}
	return coreError.NewTerror(
		7002,
		"ConnectionError",
		st.Code().String(),
		"",
	)
}

// FormatGqlTError formats the error given to a GQL error
func FormatGqlTError(err error, gqlErr *gqlerror.Error) *gqlerror.Error {
	if terror, ok := err.(*coreError.Terror); ok {
		gqlErr.Message = terror.Message()
		gqlErr.Extensions = map[string]interface{}{
			"code":      terror.Code(),
			"errorType": terror.ErrorType(),
			"status":    terror.Status(),
			"help":      terror.Help(),
		}
		return gqlErr
	}

	errStr := err.Error()
	idx := strings.Index(errStr, "{")
	message := errStr
	if idx != -1 {
		message = strings.TrimSpace(errStr[idx:])
	}

	terror, err := coreError.NewTerrorFromJSONString(message)
	if err != nil {
		gqlErr.Message = message
		return gqlErr
	}
	gqlErr.Message = terror.Message()
	gqlErr.Extensions = map[string]interface{}{
		"code":      terror.Code(),
		"errorType": terror.ErrorType(),
		"status":    terror.Status(),
		"help":      terror.Help(),
	}

	return gqlErr
}
