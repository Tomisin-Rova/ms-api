package utils

import (
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {
	type Out struct {
		Name          string  `json:"name"`
		Age           int     `json:"age"`
		WalletBalance float64 `json:"walletBalance"`
	}
	output := Out{
		Name:          "Justice Nefe",
		Age:           23,
		WalletBalance: 1000.05,
	}

	tests := []struct {
		name           string
		input          interface{}
		expectedResult interface{}
	}{
		{
			"Byte Slice Input",
			[]byte(fmt.Sprintf(`
{
	"name": "Justice Nefe",
	"age": 23,
	"walletBalance": 1000.05
}
				`)),
			output,
		},
		{
			"Struct Input",
			Out{
				Name:          "Justice Nefe",
				Age:           23,
				WalletBalance: 1000.05,
			},
			output,
		},
	}

	for _, test := range tests {
		var o Out
		if err := Pack(test.input, &o); err != nil {
			t.Errorf("%s with input: %v , expected %v but got %v instead", test.name, test.input, test.expectedResult, err)
		}
		if o != test.expectedResult {
			t.Errorf("%s with input: %v , expected %v but got %v instead", test.name, test.input, test.expectedResult, o)
		}
	}
}
