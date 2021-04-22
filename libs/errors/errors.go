package errors

import (
	coreError "github.com/roava/zebra/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"strings"
)

const (
	CddNotFound = 1105
)

// NewFromGrpc returns a formatted Roava Terror
func NewFromGrpc(err error) *coreError.Terror {
	errStr := err.Error()
	idx := strings.Index(errStr, "{")
	message := errStr
	if idx != -1 {
		message = strings.TrimSpace(errStr[idx:])
	}

	terror, err := coreError.NewTerrorFromJSONString(message)
	if err != nil {
		return nil
	}

	return terror
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
