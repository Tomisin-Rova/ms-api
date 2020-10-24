package datevalidator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	dateRegex = regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	ErrInvalidFormat = errors.New("invalid date format. Date format must be dd/mm/yyyy")
	ErrInvalidType   = errors.New("not a valid date")
	ErrInvalidAge    = errors.New("minimum age requirement for using Roava is 16years")
)

func ValidateDob(value string) error {
	if ok := dateRegex.MatchString(value); !ok {
		return ErrInvalidFormat
	}
	values := strings.Split(value, "/")
	if len(values) != 3 {
		return ErrInvalidFormat
	}

	year, err := strconv.Atoi(values[2])
	if err != nil {
		return ErrInvalidType
	}

	currentYear := time.Now().Year()
	age := currentYear - year
	if age < 16 {
		return ErrInvalidAge
	}
	return nil
}
