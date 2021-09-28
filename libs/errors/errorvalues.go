package errors

import (
	"fmt"
	"github.com/roava/zebra/errors"
)

const (
	InternalErr = 7021
)

var (
	internalErrMsg = "failed to process the request, please try again later."

	errorTypes = map[int]string{
		InternalErr: "InternalErr",
	}

	errorMessages = map[int]string{
		InternalErr: internalErrMsg,
	}

	errorDetail = map[int]string{
		InternalErr: internalErrMsg,
	}
)

func Type(code int) string {
	if value, ok := errorTypes[code]; ok {
		return value
	}
	return "UnKnownError"
}

func Message(code int) string {
	if value, ok := errorMessages[code]; ok {
		return value
	}
	return "unknown"
}

func Detail(code int, err error) string {
	if value, ok := errorDetail[code]; ok {
		return fmt.Sprintf("%s: %v", value, err)
	}
	return "unknown"
}

func Format(code int, err error) error {
	return errors.NewTerror(code, Type(code), Message(code), Detail(code, err))
}
