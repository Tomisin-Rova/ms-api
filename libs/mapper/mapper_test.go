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
	assert.Equal(t, from.Details.ProductControl.DormancyPeriodDays, *product.Details.ProductControl.DormancyPeriodDays)
	assert.Equal(t, from.Details.ProductControl.OpeningBalance.Max, *product.Details.ProductControl.OpeningBalance.Max)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.DaysInYear, *product.Details.OverdraftSetting.InterestSettings.DaysInYear)
	assert.Equal(t, from.Details.OverdraftSetting.AllowTechnicalOverdraft, *product.Details.OverdraftSetting.AllowTechnicalOverdraft)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance, *product.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance)
}
