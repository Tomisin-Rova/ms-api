package graph

import (
	"github.com/stretchr/testify/assert"
	pbTypes "ms.api/protos/pb/types"
	"ms.api/types"
	"testing"
)

func TestHelpers_ScheduledTransactionRepeatType(t *testing.T) {
	const (
		protoToSchema = iota
		schemaToProto
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test proto to schema",
			testType: protoToSchema,
		},
		{
			name:     "Test schema to proto",
			testType: schemaToProto,
		},
	}

	data := map[pbTypes.ScheduledTransaction_ScheduledTransactionRepeatType]types.ScheduledTransactionRepeatType{
		pbTypes.ScheduledTransaction_ONE_TIME: types.ScheduledTransactionRepeatTypeOneTime,
		pbTypes.ScheduledTransaction_WEEKLY:   types.ScheduledTransactionRepeatTypeWeekly,
		pbTypes.ScheduledTransaction_MONTHLY:  types.ScheduledTransactionRepeatTypeMonthly,
		pbTypes.ScheduledTransaction_ANNUALLY: types.ScheduledTransactionRepeatTypeAnnually,
		// Invalid types for coverage
		-100: "",
	}

	h := &helpersfactory{}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			for proto, schema := range data {
				switch testCase.testType {
				case protoToSchema:
					result := h.MapScheduledTransactionRepeatType(proto)
					assert.Equal(t, schema, result)
				case schemaToProto:
					result := h.MapProtoScheduledTransactionRepeatType(schema)
					assert.Equal(t, proto, result)
				}
			}
		})
	}
}

func TestHelpers_ScheduledTransactionStatus(t *testing.T) {
	const (
		protoToSchema = iota
		schemaToProto
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test proto to schema",
			testType: protoToSchema,
		},
		{
			name:     "Test schema to proto",
			testType: schemaToProto,
		},
	}

	data := map[pbTypes.ScheduledTransaction_ScheduledTransactionStatuses]types.ScheduledTransactionStatus{
		pbTypes.ScheduledTransaction_ACTIVE:   types.ScheduledTransactionStatusActive,
		pbTypes.ScheduledTransaction_INACTIVE: types.ScheduledTransactionStatusInactive,
		// Invalid types for coverage
		-100: "",
	}

	h := &helpersfactory{}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			for proto, schema := range data {
				switch testCase.testType {
				case protoToSchema:
					result := h.MapScheduledTransactionStatus(proto)
					assert.Equal(t, schema, result)
				case schemaToProto:
					result := h.MapProtoScheduledTransactionStatus(schema)
					assert.Equal(t, proto, result)
				}
			}
		})
	}
}

func TestHelpers_FeeType(t *testing.T) {
	const (
		protoToSchema = iota
		schemaToProto
	)

	var tests = []struct {
		name     string
		testType int
	}{
		{
			name:     "Test proto to schema",
			testType: protoToSchema,
		},
		{
			name:     "Test schema to proto",
			testType: schemaToProto,
		},
	}

	data := map[pbTypes.Fee_FeeTypes]types.FeeTypes{
		pbTypes.Fee_FIXED:    types.FeeTypesFixed,
		pbTypes.Fee_VARIABLE: types.FeeTypesVariable,
		// Invalid types for coverage
		-100: "",
	}

	h := &helpersfactory{}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			for proto, schema := range data {
				switch testCase.testType {
				case protoToSchema:
					result := h.MapFeeTypes(proto)
					assert.Equal(t, schema, result)
				case schemaToProto:
					result := h.MapProtoFeeTypes(schema)
					assert.Equal(t, proto, result)
				}
			}
		})
	}
}
