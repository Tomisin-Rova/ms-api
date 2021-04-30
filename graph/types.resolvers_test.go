package graph

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ms.api/types"
)

var (
	mockOwner = "personid"
	mockCdds  = []*types.Cdd{
		{
			ID: "id1",
			Owner: &types.Person{
				ID: mockOwner,
			},
			Validations: []*types.Validation{
				{ValidationType: "CHECK", Status: "APPROVED"},
				{ValidationType: "SCREEN", Status: "REJECTED"},
			},
		},
		{
			ID: "id2",
			Owner: &types.Person{
				ID: mockOwner,
			},
			Validations: []*types.Validation{
				{ValidationType: "CHECK", Status: "PENDING"},
				{ValidationType: "CHECK", Status: "PENDING"},
				{ValidationType: "SCREEN", Status: "DECLINED"},
				{ValidationType: "SCREEN", Status: "REJECTED"},
			},
		},
	}
)

func Test_Validations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Filter Empty Query Validations", func(t *testing.T) {
		resolver := cDDResolver{}
		ctx := context.Background()
		filteredValidations, err := resolver.Validations(ctx, mockCdds[0], nil, nil)
		assert.Nil(t, err, "querying without validation filter should not return an error")
		assert.Equal(t, len(mockCdds[0].Validations), len(filteredValidations), "query without validation filter should return all validations")
	})
	t.Run("Filter Type Validations", func(t *testing.T) {
		resolver := cDDResolver{}
		ctx := context.Background()
		validationType := types.ValidationType("CHECK")
		filteredValidations, err := resolver.Validations(ctx, mockCdds[0], &validationType, nil)
		assert.Nil(t, err, "querying with validation type filter should not return an error")
		assert.Equal(t, 1, len(filteredValidations), "query with validation type filter should return the correct number of validations")
	})
	t.Run("Filter by Status Validations", func(t *testing.T) {
		resolver := cDDResolver{}
		ctx := context.Background()
		states := []types.State{types.StateDeclined}
		filteredValidations, err := resolver.Validations(ctx, mockCdds[1], nil, states)
		assert.Nil(t, err, "querying with status filter should not return an error")
		assert.Equal(t, 1, len(filteredValidations), "query with validation type filter should return the correct number of validations")
	})
	t.Run("Filter by validation type and Status Validations", func(t *testing.T) {
		resolver := cDDResolver{}
		ctx := context.Background()
		validationType := types.ValidationType("SCREEN")
		states := []types.State{types.StateRejected}
		filteredValidations, err := resolver.Validations(ctx, mockCdds[1], &validationType, states)
		assert.Nil(t, err, "querying with validation type and status filter should not return an error")
		assert.Equal(t, 1, len(filteredValidations), "query with validation type and status filter should return the correct number of validations")
	})

}
