package mapper

import (
	"github.com/stretchr/testify/assert"
	pb "ms.api/protos/pb/types"
	"ms.api/types"
	"testing"
)

func TestGQLMapper_HydrateProduct(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Product{
		Id:             "id",
		Identification: "identification",
		Details: &pb.ProductDetails{
			Name: "Roava Classic GBP Current Account",
			ProductControl: &pb.ProductControl{
				DormancyPeriodDays: 12,
				OpeningBalance: &pb.OpeningBalance{
					Max: 20,
				},
			},
			OverdraftSetting: &pb.OverdraftSetting{
				AllowTechnicalOverdraft: true,
				InterestSettings: &pb.InterestSettings{
					DaysInYear: "123",
					RateTiers: []*pb.RateTiers{{
						EndingBalance: 12,
					}},
				},
			},
		},
	}

	var product types.Product
	err := mapper.Hydrate(from, &product)
	assert.Nil(t, err)

	assert.NotNil(t, product)
	assert.Equal(t, from.Id, product.ID)
	assert.Equal(t, from.Identification, *product.Identification)
	assert.Equal(t, from.Details.Name, *product.Details.Name)
	assert.Equal(t, from.Details.ProductControl.DormancyPeriodDays, *product.Details.ProductControl.DormancyPeriodDays)
	assert.Equal(t, from.Details.ProductControl.OpeningBalance.Max, *product.Details.ProductControl.OpeningBalance.Max)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.DaysInYear, *product.Details.OverdraftSetting.InterestSettings.DaysInYear)
	assert.Equal(t, from.Details.OverdraftSetting.AllowTechnicalOverdraft, *product.Details.OverdraftSetting.AllowTechnicalOverdraft)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance, *product.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance)
}

func TestGQLMapper_HydrateAccount(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Account{
		Id:    "id",
		Owner: "owner",
		AccountData: &pb.AccountData{
			Name: "name",
		},
	}

	var account types.Account
	err := mapper.Hydrate(from, &account)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, from.Id, account.ID)
}

func TestGQLMapper_HydrateTransaction(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Transaction{
		Id: "id",
		Account: &pb.Account{
			Id:    "accountId",
			Owner: "owner",
		},
		TransactionData: &pb.TransactionData{
			Id:     "id",
			Amount: 123,
			Fees: []*pb.TransactionFee{{
				Name: "name",
			}},
			AffectedAmounts: &pb.AffectedAmounts{
				FeesAmount: 43,
			},
		},
	}

	var transaction types.Transaction
	err := mapper.Hydrate(from, &transaction)
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, from.Id, transaction.ID)
	assert.Equal(t, from.Account.Id, transaction.Account.ID)
	assert.Equal(t, from.TransactionData.Fees[0].Name, *transaction.TransactionData.Fees[0].Name)
}

func TestGQLMapper_HydrateTag(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Tag{
		Id:   "id",
		Name: "name",
	}

	var tag types.Tag
	err := mapper.Hydrate(from, &tag)
	assert.Nil(t, err)

	assert.NotNil(t, tag)
	assert.Equal(t, from.Id, tag.ID)
	assert.Equal(t, from.Name, *tag.Name)
}
