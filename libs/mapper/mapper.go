package mapper

import (
	"errors"
	coreError "github.com/roava/zebra/errors"
	pb "ms.api/protos/pb/types"
	"ms.api/types"
)

type Mapper interface {
	Hydrate(from interface{}, to interface{}) error
}

// GQLMapper a mapper that returns Graphql types
type GQLMapper struct{}

var _ Mapper = &GQLMapper{}

// Hydrate converts between types
func (G *GQLMapper) Hydrate(from interface{}, to interface{}) error {
	switch value := from.(type) {
	case *pb.Product:
		return G.hydrateProduct(value, to)
	default:
		return errors.New("could not handle type")
	}
}

var (
	MappingErr = coreError.NewTerror(
		7021,
		"InternalError",
		"failed to process the request, please try again later.",
		"",
	)
)

func (G *GQLMapper) hydrateProduct(data *pb.Product, to interface{}) error {
	product, ok := to.(*types.Product)
	if !ok {
		return errors.New("invalid to type")
	}

	*product = types.Product{
		ID:             data.Id,
		Identification: &data.Identification,
		Scheme:         &data.Scheme,
	}

	if data.Details == nil {
		return nil
	}

	product.Details = &types.ProductDetails{}

	// ProductControl
	if data.Details.ProductControl != nil {
		product.Details.ProductControl = &types.ProductControl{
			DormancyPeriodDays:       &data.Details.ProductControl.DormancyPeriodDays,
			MaxWithdrawalAmount:      &data.Details.ProductControl.MaxWithdrawalAmount,
			RecommendedDepositAmount: &data.Details.ProductControl.RecommendedDepositAmount,
		}

		if data.Details.ProductControl.OpeningBalance != nil {
			product.Details.ProductControl.OpeningBalance = &types.OpeningBalance{
				DefaultValue: &data.Details.ProductControl.OpeningBalance.DefaultValue,
				Max:          &data.Details.ProductControl.OpeningBalance.Max,
				Min:          &data.Details.ProductControl.OpeningBalance.Min,
			}
		}
	}

	// ProductMaturity
	if data.Details.ProductMaturity != nil {
		product.Details.ProductMaturity = &types.ProductMaturity{
			Unit:         &data.Details.ProductMaturity.Unit,
			DefaultValue: &data.Details.ProductMaturity.DefaultValue,
			Max:          &data.Details.ProductMaturity.Max,
			Min:          &data.Details.ProductMaturity.Min,
		}
	}

	// OverdraftSetting
	if data.Details.OverdraftSetting != nil {
		product.Details.OverdraftSetting = &types.OverdraftSetting{
			AllowOverdraft:          &data.Details.OverdraftSetting.AllowOverdraft,
			AllowTechnicalOverdraft: &data.Details.OverdraftSetting.AllowTechnicalOverdraft,
			MaxLimit:                &data.Details.OverdraftSetting.MaxLimit,
		}

		// InterestSettings
		if data.Details.OverdraftSetting.InterestSettings != nil {
			rateTiers := make([]*types.RateTiers, len(data.Details.OverdraftSetting.InterestSettings.RateTiers))
			for index, value := range data.Details.OverdraftSetting.InterestSettings.RateTiers {
				rateTiers[index] = &types.RateTiers{}
				rateTiers[index].EncodedKey = String(value.EncodedKey)
				rateTiers[index].EndingBalance = Int64(value.EndingBalance)
				rateTiers[index].EndingDay = Int64(value.EndingDay)
				rateTiers[index].InterestRate = Int64(value.InterestRate)
			}

			product.Details.OverdraftSetting.InterestSettings = &types.InterestSettings{
				DaysInYear:                 &data.Details.OverdraftSetting.InterestSettings.DaysInYear,
				InterestCalculationBalance: &data.Details.OverdraftSetting.InterestSettings.InterestCalculationBalance,
				ChargeFrequency:            &data.Details.OverdraftSetting.InterestSettings.DaysInYear,
				IndexSourceKey:             &data.Details.OverdraftSetting.InterestSettings.IndexSourceKey,
				ChargeFrequencyCount:       Int64(int64(data.Details.OverdraftSetting.InterestSettings.ChargeFrequencyCount)),
				RateReviewCount:            Int64(int64(data.Details.OverdraftSetting.InterestSettings.RateReviewCount)),
				InterestRateReviewUnit:     &data.Details.OverdraftSetting.InterestSettings.InterestRateReviewUnit,
				RateSource:                 &data.Details.OverdraftSetting.InterestSettings.RateSource,
				RateTerms:                  &data.Details.OverdraftSetting.InterestSettings.RateTerms,
				RateTiers:                  rateTiers,
			}

			// Interest rate
			if data.Details.OverdraftSetting.InterestSettings.InterestRate != nil {
				product.Details.OverdraftSetting.InterestSettings.InterestRate = &types.InterestRate{
					DefaultValue: &data.Details.OverdraftSetting.InterestSettings.InterestRate.DefaultValue,
					MaxValue:     &data.Details.OverdraftSetting.InterestSettings.InterestRate.MaxValue,
					MinValue:     &data.Details.OverdraftSetting.InterestSettings.InterestRate.MinValue,
				}
			}

			// InterestPaymentSettings
			if data.Details.OverdraftSetting.InterestSettings.InterestPaymentSettings != nil {
				product.Details.OverdraftSetting.InterestSettings.InterestPaymentSettings = &types.InterestPaymentSettings{
					InterestPaymentPoint: &data.Details.OverdraftSetting.InterestSettings.InterestPaymentSettings.InterestPaymentPoint,
				}
			}

			// InterestRateSettings
			if data.Details.OverdraftSetting.InterestSettings.InterestRateSettings != nil {
				interestRateTiers := make([]*types.InterestRateTiers, len(data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateTiers))
				for index, value := range data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateTiers {
					interestRateTiers[index] = &types.InterestRateTiers{}
					interestRateTiers[index].EncodedKey = &value.EncodedKey
					interestRateTiers[index].EndingBalance = Int64(int64(value.EndingBalance))
					interestRateTiers[index].EndingDay = Int64(int64(value.EndingDay))
					interestRateTiers[index].InterestRate = Int64(int64(value.InterestRate))
				}

				product.Details.OverdraftSetting.InterestSettings.InterestRateSettings = &types.InterestRateSettings{
					EncodedKey:                   &data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.EncodedKey,
					InterestChargeFrequency:      &data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestChargeFrequency,
					InterestChargeFrequencyCount: Int64(int64(data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestChargeFrequencyCount)),
					InterestRateTerms:            &data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateTerms,
					InterestRate:                 Int64(int64(data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRate)),
					InterestRateReviewCount:      Int64(int64(data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateReviewCount)),
					InterestRateReviewUnit:       &data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateReviewUnit,
					InterestRateSource:           &data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestRateSource,
					InterestSpread:               Int64(int64(data.Details.OverdraftSetting.InterestSettings.InterestRateSettings.InterestSpread)),
					InterestRateTiers:            interestRateTiers,
				}
			}
		}
	}

	// InterestSetting
	if data.Details.InterestSetting != nil {
		interestPaymentDates := make([]*types.InterestPaymentDates, len(data.Details.InterestSetting.InterestPaymentDates))
		for index, value := range data.Details.InterestSetting.InterestPaymentDates {
			interestPaymentDates[index] = &types.InterestPaymentDates{}
			interestPaymentDates[index].Day = Int64(int64(value.Day))
			interestPaymentDates[index].Month = Int64(int64(value.Month))
		}
		product.Details.InterestSetting = &types.ProductInterestSetting{
			CollectInterestWhenLocked:  &data.Details.InterestSetting.CollectInterestWhenLocked,
			DaysInYear:                 &data.Details.InterestSetting.DaysInYear,
			InterestCalculationBalance: &data.Details.InterestSetting.InterestCalculationBalance,
			InterestPaidIntoAccount:    &data.Details.InterestSetting.CollectInterestWhenLocked,
			InterestPaymentPoint:       &data.Details.InterestSetting.InterestPaymentPoint,
			MaximumBalance:             &data.Details.InterestSetting.MaximumBalance,
			InterestPaymentDates:       interestPaymentDates,
		}

		// RateSetting
		if data.Details.InterestSetting.RateSetting != nil {
			rateSettingTiers := make([]*types.RateTiers, len(data.Details.InterestSetting.RateSetting.RateTiers))
			for index, value := range data.Details.InterestSetting.RateSetting.RateTiers {
				rateSettingTiers[index] = &types.RateTiers{}
				rateSettingTiers[index].EncodedKey = &value.EncodedKey
				rateSettingTiers[index].EndingBalance = &value.EndingBalance
				rateSettingTiers[index].EndingDay = &value.EndingDay
				rateSettingTiers[index].InterestRate = &value.InterestRate
			}

			product.Details.InterestSetting.RateSetting = &types.RateSetting{
				AccrueAfterMaturity:  &data.Details.InterestSetting.RateSetting.AccrueAfterMaturity,
				IndexSourceKey:       &data.Details.InterestSetting.RateSetting.IndexSourceKey,
				ChargeFrequency:      &data.Details.InterestSetting.RateSetting.ChargeFrequency,
				ChargeFrequencyCount: &data.Details.InterestSetting.RateSetting.ChargeFrequencyCount,
				RateSource:           &data.Details.InterestSetting.RateSetting.RateSource,
				RateTerms:            &data.Details.InterestSetting.RateSetting.RateTerms,
				RateTiers:            rateSettingTiers,
			}

			// Interest rate
			if data.Details.InterestSetting.RateSetting.InterestRate != nil {
				product.Details.InterestSetting.RateSetting.InterestRate = &types.InterestRate{
					DefaultValue: &data.Details.InterestSetting.RateSetting.InterestRate.DefaultValue,
					MaxValue:     &data.Details.InterestSetting.RateSetting.InterestRate.MaxValue,
					MinValue:     &data.Details.InterestSetting.RateSetting.InterestRate.MinValue,
				}
			}
		}
	}

	// InterestSetting
	if data.Details.InterestSetting != nil {
		accountingRules := make([]*types.AccountingRules, len(data.Details.ProductSetting.AccountingRules))
		for index, value := range data.Details.ProductSetting.AccountingRules {
			accountingRules[index] = &types.AccountingRules{}
			accountingRules[index].EncodedKey = &value.EncodedKey
			accountingRules[index].FinancialResource = &value.FinancialResource
			accountingRules[index].GlKey = &value.GlKey
		}
		product.Details.ProductSetting = &types.ProductSetting{
			AccountingMethod:   &data.Details.ProductSetting.AccountingMethod,
			InterestAccounting: &data.Details.ProductSetting.InterestAccounting,
			AccountingRules:    accountingRules,
		}
	}

	return nil
}

func String(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Int64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}
