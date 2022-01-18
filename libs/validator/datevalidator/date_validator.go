package datevalidator

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	coreError "github.com/roava/zebra/errors"
)

var (
	dateRegex        = regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\d\d)`)
	ErrInvalidFormat = coreError.NewTerror(
		7007,
		"ErrInvalidDateFormat",
		"invalid date format. Date format must be dd/mm/yyyy",
		"",
	)
	ErrInvalidType = coreError.NewTerror(
		7008,
		"ErrInvalidType",
		"not a valid date",
		"",
	)
	ErrInvalidAge = coreError.NewTerror(
		7009,
		"ErrInvalidAge",
		"minimum age requirement for using Roava is 18years",
		"",
	)
)

func ValidateDob(value string) error {
	if ok := dateRegex.MatchString(value); !ok {
		return ErrInvalidFormat
	}
	values := strings.Split(value, "-")
	if len(values) != 3 {
		return ErrInvalidFormat
	}

	year, err := strconv.Atoi(values[2])
	if err != nil {
		return ErrInvalidType
	}

	currentYear := time.Now().Year()
	age := currentYear - year
	if age < 18 {
		return ErrInvalidAge
	}
	return nil
}
