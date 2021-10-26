package mapper

import (
	"errors"

	coreError "github.com/roava/zebra/errors"
	zaplogger "github.com/roava/zebra/logger"
	"go.uber.org/zap"
	pb "ms.api/protos/pb/types"
	"ms.api/types"
)

type Mapper interface {
	Hydrate(from interface{}, to interface{}) error
}

// GQLMapper a mapper that returns Graphql types
type GQLMapper struct {
	logger *zap.Logger
}

var _ Mapper = &GQLMapper{
	logger: zaplogger.New(),
}

// Hydrate converts between types
func (G *GQLMapper) Hydrate(from interface{}, to interface{}) error {
	switch value := from.(type) {
	case *pb.Product:
		return G.hydrateProduct(value, to)
	case *pb.Account:
		return G.hydrateAccount(value, to)
	case *pb.Tag:
		return G.hydrateTag(value, to)
	case *pb.Transaction:
		return G.hydrateTransaction(value, to)
	case *pb.Payment:
		return G.hydratePayment(value, to)
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

	product.Details = &types.ProductDetails{
		Category:              &data.Details.Category,
		Type:                  &data.Details.Type,
		Name:                  &data.Details.Name,
		State:                 &data.Details.State,
		Currency:              &data.Details.Currency,
		Notes:                 &data.Details.Notes,
		CreditRequirement:     &data.Details.CreditRequirement,
		WithholdingTaxEnabled: &data.Details.WithholdingTaxEnabled,
		AllowOffset:           &data.Details.AllowOffset,
	}

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
					interestRateTiers[index].EndingBalance = &value.EndingBalance
					interestRateTiers[index].EndingDay = Int64(int64(value.EndingDay))
					interestRateTiers[index].InterestRate = &value.InterestRate
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

func (G *GQLMapper) hydrateAccount(data *pb.Account, to interface{}) error {
	account, ok := to.(*types.Account)
	if !ok {
		return errors.New("invalid to type")
	}

	*account = types.Account{
		ID:           data.Id,
		Owner:        types.Person{ID: data.Owner},
		Product:      &types.Product{ID: data.Product},
		Name:         &data.Name,
		Active:       &data.Active,
		Status:       &data.Status,
		Image:        &data.Image,
		Organisation: &types.Organisation{ID: data.Organisation},
		Ts:           Int64(int64(data.Ts)),
	}

	// Account details
	if data.AccountDetails != nil {
		account.AccountDetails = &types.AccountDetails{
			VirtualAccountID: &data.AccountDetails.VirtualAccountID,
			Iban:             &data.AccountDetails.Iban,
			AccountNumber:    &data.AccountDetails.AccountNumber,
			SortCode:         &data.AccountDetails.SortCode,
			SwiftBic:         &data.AccountDetails.SwiftBic,
			BankCode:         &data.AccountDetails.BankCode,
			RoutingNumber:    &data.AccountDetails.RoutingNumber,
		}
	}

	// Account data
	if data.AccountData != nil {
		account.AccountData = &types.AccountData{
			AccountHolderKey:                &data.AccountData.AccountHolderKey,
			AccountHolderType:               &data.AccountData.AccountHolderType,
			AccountState:                    &data.AccountData.AccountState,
			AccountType:                     &data.AccountData.AccountType,
			ActivationDate:                  String(string(data.AccountData.ActivationDate)),
			ApprovedDate:                    String(string(data.AccountData.ApprovedDate)),
			AssignedBranchKey:               &data.AccountData.AssignedBranchKey,
			AssignedCentreKey:               &data.AccountData.AssignedCentreKey,
			AssignedUserKey:                 &data.AccountData.AssignedUserKey,
			ClosedDate:                      &data.AccountData.ClosedDate,
			CreationDate:                    &data.AccountData.CreationDate,
			CreditArrangementKey:            &data.AccountData.CreditArrangementKey,
			CurrencyCode:                    &data.AccountData.CurrencyCode,
			EncodedKey:                      &data.AccountData.EncodedKey,
			LastAccountAppraisalDate:        &data.AccountData.LastAccountAppraisalDate,
			LastInterestCalculationDate:     &data.AccountData.LastInterestCalculationDate,
			LastInterestStoredDate:          &data.AccountData.LastInterestStoredDate,
			LastModifiedDate:                &data.AccountData.LastModifiedDate,
			LastOverdraftInterestReviewDate: &data.AccountData.LastOverdraftInterestReviewDate,
			LastSetToArrearsDate:            &data.AccountData.LastSetToArrearsDate,
			LockedDate:                      &data.AccountData.LockedDate,
			MaturityDate:                    &data.AccountData.MaturityDate,
			MigrationEventKey:               &data.AccountData.MigrationEventKey,
			Name:                            &data.AccountData.Name,
			Notes:                           &data.AccountData.Notes,
			ProductTypeKey:                  &data.AccountData.ProductTypeKey,
			WithholdingTaxSourceKey:         &data.AccountData.WithholdingTaxSourceKey,
		}

		// AccruedAmounts
		if data.AccountData.AccruedAmounts != nil {
			account.AccountData.AccruedAmounts = &types.AccruedAmounts{
				InterestAccrued:                   &data.AccountData.AccruedAmounts.InterestAccrued,
				OverdraftInterestAccrued:          &data.AccountData.AccruedAmounts.OverdraftInterestAccrued,
				TechnicalOverdraftInterestAccrued: &data.AccountData.AccruedAmounts.TechnicalOverdraftInterestAccrued,
			}
		}

		// Balances
		if data.AccountData.Balances != nil {
			account.AccountData.Balances = &types.Balances{
				AvailableBalance:              &data.AccountData.Balances.AvailableBalance,
				BlockedBalance:                &data.AccountData.Balances.BlockedBalance,
				FeesDue:                       &data.AccountData.Balances.FeesDue,
				ForwardAvailableBalance:       &data.AccountData.Balances.ForwardAvailableBalance,
				HoldBalance:                   &data.AccountData.Balances.HoldBalance,
				LockedBalance:                 &data.AccountData.Balances.LockedBalance,
				OverdraftAmount:               &data.AccountData.Balances.OverdraftAmount,
				OverdraftInterestDue:          &data.AccountData.Balances.OverdraftInterestDue,
				TechnicalOverdraftAmount:      &data.AccountData.Balances.TechnicalOverdraftAmount,
				TechnicalOverdraftInterestDue: &data.AccountData.Balances.TechnicalOverdraftInterestDue,
				TotalBalance:                  &data.AccountData.Balances.TotalBalance,
			}
		}

		// InternalControls
		if data.AccountData.InternalControls != nil {
			account.AccountData.InternalControls = &types.InternalControls{
				MaxWithdrawalAmount:      &data.AccountData.InternalControls.MaxWithdrawalAmount,
				RecommendedDepositAmount: &data.AccountData.InternalControls.RecommendedDepositAmount,
				TargetAmount:             &data.AccountData.InternalControls.TargetAmount,
			}
		}

		// OverdraftSettings
		if data.AccountData.OverdraftSettings != nil {
			account.AccountData.OverdraftSettings = &types.OverdraftSettings{
				AllowOverdraft: &data.AccountData.OverdraftSettings.AllowOverdraft,
				OverdraftLimit: Int64(int64(data.AccountData.OverdraftSettings.OverdraftLimit)),
			}
		}

		// InterestSettings
		if data.AccountData.InterestSettings != nil && data.AccountData.InterestSettings.InterestPaymentSettings != nil {

			account.AccountData.InterestSettings = &types.InterestSettings{}

			// InterestPaymentSettings
			if data.AccountData.InterestSettings.InterestPaymentSettings != nil {
				account.AccountData.InterestSettings.InterestPaymentSettings = &types.InterestPaymentSettings{
					InterestPaymentPoint: &data.AccountData.InterestSettings.InterestPaymentSettings.InterestPaymentPoint,
				}
			}

			// InterestRateSettings
			if data.AccountData.InterestSettings.InterestPaymentSettings != nil {
				account.AccountData.InterestSettings.InterestRateSettings = &types.InterestRateSettings{
					EncodedKey:                   &data.AccountData.InterestSettings.InterestRateSettings.EncodedKey,
					InterestChargeFrequency:      &data.AccountData.InterestSettings.InterestRateSettings.InterestChargeFrequency,
					InterestChargeFrequencyCount: Int64(int64(data.AccountData.InterestSettings.InterestRateSettings.InterestChargeFrequencyCount)),
					InterestRate:                 Int64(int64(data.AccountData.InterestSettings.InterestRateSettings.InterestRate)),
					InterestRateTerms:            &data.AccountData.InterestSettings.InterestRateSettings.InterestRateTerms,
				}
			}

		}

	}
	return nil
}

func (G *GQLMapper) hydrateTransaction(data *pb.Transaction, to interface{}) error {
	transaction, ok := to.(*types.Transaction)
	if !ok {
		return errors.New("invalid to type")
	}

	*transaction = types.Transaction{
		ID: data.Id,
		Ts: Int64(int64(data.Ts)),
	}

	// Account data
	if data.TransactionData != nil {
		transaction.TransactionData = &types.TransactionData{
			ID:               data.TransactionData.Id,
			Amount:           &data.TransactionData.Amount,
			BookingDate:      &data.TransactionData.BookingDate,
			CreationDate:     &data.TransactionData.CreationDate,
			CurrencyCode:     &data.TransactionData.CurrencyCode,
			EncodedKey:       &data.TransactionData.EncodedKey,
			ExternalID:       &data.TransactionData.ExternalID,
			Notes:            &data.TransactionData.Notes,
			ParentAccountKey: &data.TransactionData.ParentAccountKey,
			PaymentOrderID:   &data.TransactionData.PaymentOrderID,
			Type:             &data.TransactionData.Type,
			UserKey:          &data.TransactionData.UserKey,
			ValueDate:        &data.TransactionData.ValueDate,
		}

		// TransferDetails
		if data.TransactionData.TransferDetails != nil {
			transaction.TransactionData.TransferDetails = &types.TransferDetails{
				LinkedLoanTransactionKey: &data.TransactionData.TransferDetails.LinkedLoanTransactionKey,
			}
		}

		// AffectedAmounts
		if data.TransactionData.AffectedAmounts != nil {
			transaction.TransactionData.AffectedAmounts = &types.AffectedAmounts{
				FeesAmount:                       &data.TransactionData.AffectedAmounts.FeesAmount,
				FractionAmount:                   &data.TransactionData.AffectedAmounts.FractionAmount,
				FundsAmount:                      &data.TransactionData.AffectedAmounts.FundsAmount,
				InterestAmount:                   &data.TransactionData.AffectedAmounts.InterestAmount,
				OverdraftAmount:                  &data.TransactionData.AffectedAmounts.OverdraftAmount,
				OverdraftFeesAmount:              &data.TransactionData.AffectedAmounts.OverdraftFeesAmount,
				OverdraftInterestAmount:          &data.TransactionData.AffectedAmounts.OverdraftInterestAmount,
				TechnicalOverdraftAmount:         &data.TransactionData.AffectedAmounts.TechnicalOverdraftAmount,
				TechnicalOverdraftInterestAmount: &data.TransactionData.AffectedAmounts.TechnicalOverdraftInterestAmount,
			}
		}

		// AccountBalances
		if data.TransactionData.AccountBalances != nil {
			transaction.TransactionData.AccountBalances = &types.AccountBalances{
				TotalBalance: &data.TransactionData.AccountBalances.TotalBalance,
			}
		}

		// Fees
		if data.TransactionData.Fees != nil {
			transFees := make([]*types.TransactionFee, len(data.TransactionData.Fees))
			for index, tfee := range data.TransactionData.Fees {

				transFees[index] = &types.TransactionFee{
					Name:             &tfee.Name,
					Amount:           Int64(int64(tfee.Amount)),
					PredefinedFeeKey: &tfee.PredefinedFeeKey,
					TaxAmount:        Int64(int64(tfee.TaxAmount)),
					Trigger:          &tfee.Trigger,
				}
			}
			transaction.TransactionData.Fees = transFees
		}

	}

	if data.Account != nil {

		// remove second level of transactions
		data.Account.Transactions = nil

		var account types.Account
		err := G.hydrateAccount(data.Account, &account)
		if err != nil {
			return err
		}

		transaction.Account = &account
	}

	return nil
}

func (G *GQLMapper) hydrateTag(data *pb.Tag, to interface{}) error {
	tag, ok := to.(*types.Tag)
	if !ok {
		return errors.New("invalid to type")
	}
	*tag = types.Tag{
		ID:   data.Id,
		Name: &data.Name,
	}
	return nil
}

func (G *GQLMapper) hydrateOwner(data *pb.Person, to interface{}) error {
	owner, ok := to.(*types.Owner)
	if !ok {
		return errors.New("invalid to type")
	}

	person := data
	*owner = types.Person{
		ID:               person.Id,
		Title:            &person.Title,
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		MiddleName:       &person.MiddleName,
		Dob:              person.Dob,
		Status:           (*types.PersonStatus)(&person.Status),
		Ts:               person.Ts,
		CountryResidence: &person.CountryResidence,
		Bvn:              &person.Bvn,
		Phones:           phonesFromProto(person.Phones),
		Emails:           emailsFromProto(person.Emails),
	}
	return nil
}

func (G *GQLMapper) hydratePayeeAccount(data *pb.PayeeAccount, to interface{}) error {
	payeeAccount, ok := to.(*types.PayeeAccount)
	if !ok {
		return errors.New("invalid to type")
	}
	*payeeAccount = types.PayeeAccount{
		ID:            data.Id,
		Name:          &data.Name,
		Currency:      &data.Currency,
		AccountNumber: &data.AccountNumber,
		SortCode:      &data.SortCode,
		Iban:          &data.Iban,
		SwiftBic:      &data.SwiftBic,
		BankCode:      &data.BankCode,
		RoutingNumber: &data.RoutingNumber,
		PhoneNumber:   &data.PhoneNumber,
	}
	return nil
}

func emailsFromProto(protoEmails []*pb.Email) []*types.Email {
	emails := []*types.Email{}
	for _, protoEmail := range protoEmails {
		email := &types.Email{
			Value:    protoEmail.Value,
			Verified: protoEmail.Verified,
		}
		emails = append(emails, email)
	}
	return emails
}

func phonesFromProto(protoEmails []*pb.PhoneNumber) []*types.Phone {
	phones := []*types.Phone{}
	for _, protoPhone := range protoEmails {
		phone := &types.Phone{
			Value:    protoPhone.Number,
			Verified: protoPhone.Verified,
		}
		phones = append(phones, phone)
	}
	return phones
}

func (G *GQLMapper) hydratePayment(data *pb.Payment, to interface{}) error {
	payment, ok := to.(*types.Payment)
	if !ok {
		return errors.New("invalid to type")
	}
	*payment = types.Payment{
		ID:             &data.Id,
		IdempotencyKey: data.IdempotencyKey,
		Charge:         Float64(float64(data.Charge)),
		Reference:      String(data.Reference),
		Status:         (*types.State)(&data.Status),
		Image:          String(data.Image),
		Notes:          String(data.Notes),
		Currency:       &types.Currency{},
		FundingAmount:  float64(data.FundingAmount),
		Ts:             &data.Ts,
		Beneficiary:    new(types.Beneficiary),
		FundingSource:  new(types.Account),
	}

	if data.Tags != nil {
		tags := make([]*types.Tag, len(data.Tags))
		for index, v := range data.Tags {
			tags[index] = &types.Tag{
				Name: String(v),
			}
		}
		payment.Tags = tags
	}

	if data.Owner != nil {
		err := G.hydrateOwner(data.Owner, &payment.Owner)
		if err != nil {
			G.logger.Error("hydrate owner", zap.String("payment_id", data.Id), zap.Error(err))
			return err
		}
	}

	if data.Source != nil && data.Source.Accounts != nil {
		switch data.Source.Accounts.(type) {
		case *pb.PaymentAccount_Account:
			sourceAccount := data.Source.Accounts.(*pb.PaymentAccount_Account)
			if sourceAccount.Account == nil {
				G.logger.Error("source account decoding", zap.String("payment_id", data.Id))
				payment.FundingSource = nil
				return errors.New("invalid source account")
			}
			err := G.hydrateAccount(sourceAccount.Account, payment.FundingSource)
			if err != nil {
				G.logger.Error("hydrate source account", zap.String("payment_id", data.Id))
				return err
			}
			payment.FundingSource.Owner = payment.Owner
		case *pb.PaymentAccount_PayeeAccount:
			payment.FundingSource = nil
		}
	}

	if data.Target != nil {
		if data.Target.Accounts != nil {
			switch data.Target.Accounts.(type) {
			case *pb.PaymentAccount_Account:
				targetAccount := data.Target.Accounts.(*pb.PaymentAccount_Account)
				if targetAccount.Account == nil {
					G.logger.Error("target account decoding", zap.String("payment_id", data.Id))
					payment.Beneficiary.Account = nil
					return errors.New("undefined target account")
				}

				account := types.Account{}
				err := G.hydrateAccount(targetAccount.Account, &account)
				if err != nil {
					G.logger.Error("hydrate target account", zap.String("payment_id", data.Id))
					return err
				}
				account.Owner = payment.Owner
				payment.Beneficiary.Account = &account
			case *pb.PaymentAccount_PayeeAccount:
				targetPayeeAccount := data.Target.Accounts.(*pb.PaymentAccount_PayeeAccount)
				if targetPayeeAccount.PayeeAccount == nil {
					G.logger.Error("target payee account decoding", zap.String("payment_id", data.Id))
					payment.Beneficiary.Account = nil
					return errors.New("undefined target account")
				}
				payeeAccount := types.PayeeAccount{}
				err := G.hydratePayeeAccount(targetPayeeAccount.PayeeAccount, &payeeAccount)
				if err != nil {
					G.logger.Error("hydrate target payee account", zap.String("payment_id", data.Id))
					return err
				}
				payment.Beneficiary.Account = &payeeAccount
			}
		}
		payment.Beneficiary.Amount = data.Target.Amount
		payment.Beneficiary.Currency = &types.Currency{
			Code: data.Target.Currency,
		}
	}

	if data.Quote != nil {
		var fee *types.Fee
		var fx *types.Fx

		if data.Quote.Fee != nil {
			fee = &types.Fee{
				LowerBoundary: &data.Quote.Fee.LowerBoundary,
				UpperBoundary: &data.Quote.Fee.UpperBoundary,
				Fee:           data.Quote.Fee.Fee,
			}
		}
		if data.Quote.Fx != nil {
			fx = &types.Fx{
				Currency:     data.Quote.Fx.Currency,
				BaseCurrency: data.Quote.Fx.BaseCurrency,
				BuyRate:      data.Quote.Fx.BuyRate,
				SellRate:     data.Quote.Fx.SellRate,
				Ts:           data.Quote.Fx.Ts,
			}
		}

		payment.Quote = &types.Quote{
			ID:        data.Quote.Id,
			HasExpiry: &data.Quote.HasExpiry,
			Expires:   &data.Quote.Expires,
			Ts:        &data.Quote.Ts,
			Fee:       fee,
			Fx:        fx,
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

func Float64(i float64) *float64 {
	if i == 0 {
		return nil
	}
	return &i
}

func NewMapper() *GQLMapper {
	return &GQLMapper{
		logger: zaplogger.New(),
	}
}
