package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ms.api/libs/mapper"

	"ms.api/libs/preloader"

	"ms.api/config"
	"ms.api/libs/db"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/identityService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/onfidoService"
	"ms.api/protos/pb/payeeService"
	"ms.api/protos/pb/paymentService"
	"ms.api/protos/pb/personService"
	pb "ms.api/protos/pb/types"
	"ms.api/protos/pb/verifyService"
	"ms.api/server/http/middlewares"
	"ms.api/types"

	"github.com/roava/zebra/errors"
	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type OnboardStatus string

// CddChunk store the data to fill cdds slice
// on the right position
type CddChunk struct {
	pos  int
	cdds []*types.Cdd
}

const (
	errorMarshallingScreenValidation               = "marshall screen validation"
	Onboarded                        OnboardStatus = "ONBOARDED"
	NotOnboarded                     OnboardStatus = "NOT_ONBOARDED"
	IgnoreOnboardFilter              OnboardStatus = "IGNORE_ONBOARD_FILTER"
)

var (
	ErrUnAuthenticated = errors.NewTerror(
		7012, "InvalidOrExpiredTokenError", "user not authenticated", "user not authenticated")
)

//nolint
func (r *mutationResolver) validateAddress(addr types.AddressInput) error {
	if addr.Country == nil || *addr.Country == "" {
		return errors.NewTerror(7013, "InvalidCountryData", "country data is missing from address", "")
	}
	if addr.City == nil || *addr.City == "" {
		return errors.NewTerror(7014, "InvalidCityData", "city data is missing from address", "")
	}
	if addr.Street == nil || *addr.Street == "" {
		return errors.NewTerror(7015, "InvalidStreetData", "street data is missing from address", "")
	}
	return nil
}

type ResolverOpts struct {
	PayeeService      payeeService.PayeeServiceClient
	OnfidoClient      onfidoService.OnfidoServiceClient
	cddClient         cddService.CddServiceClient
	accountService    accountService.AccountServiceClient
	OnBoardingService onboardingService.OnBoardingServiceClient
	verifyService     verifyService.VerifyServiceClient
	AuthService       authService.AuthServiceClient
	paymentService    paymentService.PaymentServiceClient
	AuthMw            *middlewares.AuthMiddleware
	personService     personService.PersonServiceClient
	identityService   identityService.IdentityServiceClient
	DataStore         db.DataStore
	preloader         preloader.Preloader
	mapper            mapper.Mapper
}

type Resolver struct {
	PayeeService      payeeService.PayeeServiceClient
	cddService        cddService.CddServiceClient
	onBoardingService onboardingService.OnBoardingServiceClient
	accountService    accountService.AccountServiceClient
	personService     personService.PersonServiceClient
	verifyService     verifyService.VerifyServiceClient
	onfidoClient      onfidoService.OnfidoServiceClient
	authService       authService.AuthServiceClient
	paymentService    paymentService.PaymentServiceClient
	identityService   identityService.IdentityServiceClient
	authMw            *middlewares.AuthMiddleware
	logger            *zap.Logger
	dataStore         db.DataStore
	preloader         preloader.Preloader
	mapper            mapper.Mapper
}

func NewResolver(opt *ResolverOpts, logger *zap.Logger) *Resolver {
	return &Resolver{
		PayeeService:      opt.PayeeService,
		cddService:        opt.cddClient,
		onBoardingService: opt.OnBoardingService,
		accountService:    opt.accountService,
		personService:     opt.personService,
		verifyService:     opt.verifyService,
		onfidoClient:      opt.OnfidoClient,
		authService:       opt.AuthService,
		paymentService:    opt.paymentService,
		identityService:   opt.identityService,
		authMw:            opt.AuthMw,
		dataStore:         opt.DataStore,
		logger:            logger,
		preloader:         opt.preloader,
		mapper:            opt.mapper,
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	opts := &ResolverOpts{
		preloader: preloader.GQLPreloader{},
		mapper:    &mapper.GQLMapper{},
	}

	// OnBoarding
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
	}
	opts.OnBoardingService = onboardingService.NewOnBoardingServiceClient(connection)

	// CDD
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.CddServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.CddServiceURL)
	}
	opts.cddClient = cddService.NewCddServiceClient(connection)

	// Verify
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.VerifyServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.VerifyServiceURL)
	}
	opts.verifyService = verifyService.NewVerifyServiceClient(connection)

	// Auth
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.AuthServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.AuthServiceURL)
	}
	opts.AuthService = authService.NewAuthServiceClient(connection)

	// Account
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.AccountServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.AccountServiceURL)
	}
	opts.accountService = accountService.NewAccountServiceClient(connection)

	// Payment
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.PaymentServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.PaymentServiceURL)
	}
	opts.paymentService = paymentService.NewPaymentServiceClient(connection)

	// Person
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.PersonServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.PersonServiceURL)
	}
	opts.personService = personService.NewPersonServiceClient(connection)

	// Identity
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.IdentityServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.IdentityServiceURL)
	}
	opts.identityService = identityService.NewIdentityServiceClient(connection)

	return opts, nil
}

func dialRPC(ctx context.Context, address string) (*grpc.ClientConn, error) {
	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func getPerson(from *pb.Person) (*types.Person, error) {
	personStatus := types.PersonStatus(from.Status)
	var person = types.Person{
		ID:               from.Id,
		Title:            &from.Title,
		FirstName:        from.FirstName,
		LastName:         from.LastName,
		MiddleName:       &from.MiddleName,
		Dob:              from.Dob,
		Ts:               from.Ts,
		CountryResidence: &from.CountryResidence,
		Status:           &personStatus,
	}
	// TODO: Fill other attributes

	addresses := make([]*types.Address, len(from.Addresses))
	for i, addr := range from.Addresses {
		addresses[i] = &types.Address{
			Street:   &addr.Street,
			Postcode: &addr.Postcode,
			Country: &types.Country{
				CountryName: addr.Country,
			},
			City: &addr.City,
		}
	}
	person.Addresses = addresses

	// Add Phones
	phones := make([]*types.Phone, len(from.Phones))
	for i, ph := range from.Phones {
		phones[i] = &types.Phone{
			Value:    ph.Number,
			Verified: ph.Verified,
		}
	}
	person.Phones = phones

	// Add Emails
	emails := make([]*types.Email, len(from.Emails))
	for i, em := range from.Emails {
		emails[i] = &types.Email{
			Value:    em.Value,
			Verified: em.Verified,
		}
	}
	person.Emails = emails

	// Add Activity
	activities := make([]*types.Activity, len(from.Activities))
	for i, ac := range from.Activities {
		activities[i] = &types.Activity{
			ID:            ac.Id,
			Description:   ac.Description,
			RiskWeighting: int64(ac.RiskWeighting),
			Supported:     &ac.Supported,
			Archived:      &ac.Archived,
			Ts:            &ac.Ts,
		}
	}
	person.Activities = activities

	// Add Identities
	identities := make([]*types.Identity, len(from.Identities))
	for i, id := range from.Identities {
		identities[i] = &types.Identity{
			ID:             id.Id,
			Nickname:       &id.Nickname,
			Active:         &id.Active,
			Authentication: &id.Authentication,
			Credentials: &types.Credentials{
				Identifier:   id.Credentials.Identifier,
				RefreshToken: &id.Credentials.RefreshToken,
			},
			Organisation: &types.Organisation{
				ID:   id.Organisation.Id,
				Name: &id.Organisation.Name,
			},
			Ts: id.Ts,
		}
	}
	person.Identities = identities

	return &person, nil
}

func (r *Resolver) hydrateAccount(from *accountService.GetAccountResponse) *types.Account {
	dto := types.Account{
		ID:           &from.Id,
		Owner:        &from.Owner,
		Product:      &from.Product,
		Name:         &from.Name,
		Active:       &from.Active,
		Status:       &from.Status,
		Image:        &from.Image,
		Organisation: &from.Organisation,
		Ts:           Int64(int64(from.Ts)),
		AccountDetails: &types.AccountDetails{
			VirtualAccountID: &from.AccountDetails.VirtualAccountID,
			Iban:             &from.AccountDetails.Iban,
			AccountNumber:    &from.AccountDetails.AccountNumber,
			SortCode:         &from.AccountDetails.SortCode,
			SwiftBic:         &from.AccountDetails.SwiftBic,
			BankCode:         &from.AccountDetails.BankCode,
			RoutingNumber:    &from.AccountDetails.RoutingNumber,
		},
	}

	// Add tags
	tags := make([]*string, len(from.Tags))
	for index, tag := range from.Tags {
		tags[index] = &tag
	}
	dto.Tags = tags

	// Add Account Data

	dto.AccountData = &types.AccountData{
		AccountHolderKey:  &from.AccountData.AccountHolderKey,
		AccountHolderType: &from.AccountData.AccountHolderType,
		AccountState:      &from.AccountData.AccountState,
		AccountType:       &from.AccountData.AccountType,
		AccruedAmounts: &types.AccruedAmounts{
			InterestAccrued:                   Int64(int64(from.AccountData.AccruedAmounts.InterestAccrued)),
			OverdraftInterestAccrued:          Int64(int64(from.AccountData.AccruedAmounts.OverdraftInterestAccrued)),
			TechnicalOverdraftInterestAccrued: Int64(int64(from.AccountData.AccruedAmounts.TechnicalOverdraftInterestAccrued)),
		},
		ActivationDate:    String(string(from.AccountData.ActivationDate)),
		ApprovedDate:      String(string(from.AccountData.ApprovedDate)),
		AssignedBranchKey: &from.AccountData.AssignedBranchKey,
		AssignedCentreKey: &from.AccountData.AssignedCentreKey,
		AssignedUserKey:   &from.AccountData.AssignedUserKey,
		Balances: &types.Balances{
			AvailableBalance:              Int64(int64(from.AccountData.Balances.AvailableBalance)),
			BlockedBalance:                Int64(int64(from.AccountData.Balances.BlockedBalance)),
			FeesDue:                       Int64(int64(from.AccountData.Balances.FeesDue)),
			ForwardAvailableBalance:       Int64(int64(from.AccountData.Balances.ForwardAvailableBalance)),
			HoldBalance:                   Int64(int64(from.AccountData.Balances.HoldBalance)),
			LockedBalance:                 Int64(int64(from.AccountData.Balances.LockedBalance)),
			OverdraftAmount:               Int64(int64(from.AccountData.Balances.OverdraftAmount)),
			OverdraftInterestDue:          Int64(int64(from.AccountData.Balances.OverdraftInterestDue)),
			TechnicalOverdraftAmount:      Int64(int64(from.AccountData.Balances.TechnicalOverdraftAmount)),
			TechnicalOverdraftInterestDue: Int64(int64(from.AccountData.Balances.TechnicalOverdraftInterestDue)),
			TotalBalance:                  Int64(int64(from.AccountData.Balances.TotalBalance)),
		},
		ClosedDate:           &from.AccountData.ClosedDate,
		CreationDate:         &from.AccountData.CreationDate,
		CreditArrangementKey: &from.AccountData.CreditArrangementKey,
		CurrencyCode:         &from.AccountData.CurrencyCode,
		EncodedKey:           &from.AccountData.EncodedKey,
		InterestSettings: &types.InterestSettings{
			InterestPaymentSettings: &types.InterestPaymentSettings{
				InterestPaymentPoint: &from.AccountData.InterestSettings.InterestPaymentSettings.InterestPaymentPoint,
			},
			InterestRateSettings: &types.InterestRateSettings{
				EncodedKey:                   &from.AccountData.InterestSettings.InterestRateSettings.EncodedKey,
				InterestChargeFrequency:      &from.AccountData.InterestSettings.InterestRateSettings.InterestChargeFrequency,
				InterestChargeFrequencyCount: Int64(int64(from.AccountData.InterestSettings.InterestRateSettings.InterestChargeFrequencyCount)),
				InterestRate:                 Int64(int64(from.AccountData.InterestSettings.InterestRateSettings.InterestRate)),
				InterestRateTerms:            &from.AccountData.InterestSettings.InterestRateSettings.InterestRateTerms,
			},
		},
		InternalControls: &types.InternalControls{
			MaxWithdrawalAmount:      Int64(int64(from.AccountData.InternalControls.MaxWithdrawalAmount)),
			RecommendedDepositAmount: Int64(int64(from.AccountData.InternalControls.RecommendedDepositAmount)),
			TargetAmount:             Int64(int64(from.AccountData.InternalControls.TargetAmount)),
		},
		LastAccountAppraisalDate:        &from.AccountData.LastAccountAppraisalDate,
		LastInterestCalculationDate:     &from.AccountData.LastInterestCalculationDate,
		LastInterestStoredDate:          &from.AccountData.LastInterestStoredDate,
		LastModifiedDate:                &from.AccountData.LastModifiedDate,
		LastOverdraftInterestReviewDate: &from.AccountData.LastOverdraftInterestReviewDate,
		LastSetToArrearsDate:            &from.AccountData.LastSetToArrearsDate,
		LockedDate:                      &from.AccountData.LockedDate,
		MaturityDate:                    &from.AccountData.MaturityDate,
		MigrationEventKey:               &from.AccountData.MigrationEventKey,
		Name:                            &from.AccountData.Name,
		Notes:                           &from.AccountData.Notes,
		OverdraftSettings: &types.OverdraftSettings{
			AllowOverdraft: &from.AccountData.OverdraftSettings.AllowOverdraft,
			OverdraftLimit: Int64(int64(from.AccountData.OverdraftSettings.OverdraftLimit)),
		},
		ProductTypeKey:          &from.AccountData.ProductTypeKey,
		WithholdingTaxSourceKey: &from.AccountData.WithholdingTaxSourceKey,
	}

	return &dto
}

func (r *queryResolver) hydrateCDD(cddDto *pb.Cdd) (*types.Cdd, error) {
	if cddDto == nil {
		return nil, nil
	}

	tsAsInt64 := int64(cddDto.Ts)
	var cdd = types.Cdd{
		ID:        cddDto.Id,
		Watchlist: &cddDto.Watchlist,
		Status:    types.State(cddDto.Status),
		Onboard:   &cddDto.Onboard,
		Ts:        &tsAsInt64,
	}
	// Add validations
	for _, validationDto := range cddDto.Validations {
		tsAsInt64 := int64(validationDto.Ts)
		actions := make([]*types.Action, len(validationDto.Actions))

		for index, action := range validationDto.Actions {
			person, err := getPerson(action.Reporter)
			if err != nil {
				return nil, err
			}
			actions[index] = &types.Action{
				ID:       action.Id,
				Reporter: person,
				Notes:    action.Notes,
				Status:   action.Status,
				Ts:       tsAsInt64,
			}
		}

		validation := types.Validation{
			ID:             validationDto.Id,
			ValidationType: types.ValidationType(validationDto.ValidationType),
			Status:         types.State(validationDto.Status),
			Approved:       &validationDto.Approved,
			Ts:             &tsAsInt64,
			Actions:        actions,
		}
		// Fill validation Data
		switch validationDto.Data.TypeUrl {
		case models.SCREEN:
			var screen models.Screen
			err := json.Unmarshal(validationDto.Data.Value, &screen)
			if err != nil {
				r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
				continue
			}
			tsAsInt64 := screen.Timestamp.Unix()
			var data = types.Screen{
				ID:     screen.ID,
				Data:   string(screen.Data),
				Status: types.State(screen.Status),
				Ts:     &tsAsInt64,
			}

			// Add data to validation
			validation.Data = &data
		case models.CHECK:
			var check models.Check
			err := json.Unmarshal(validationDto.Data.Value, &check)
			if err != nil {
				r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
				continue
			}
			tsAsInt64 := check.Timestamp.Unix()
			createdAtAsString := check.Data.CreatedAt.Format(time.RFC3339)
			var data = types.Check{
				ID: check.ID,
				Data: &types.CheckData{
					ID:                    check.Data.ID,
					CreatedAt:             &createdAtAsString,
					Status:                types.State(check.Data.Status),
					Sandbox:               &check.Data.Sandbox,
					ResultsURI:            &check.Data.ResultsURI,
					FormURI:               &check.Data.FormURI,
					Paused:                &check.Data.Paused,
					Version:               &check.Data.Version,
					Href:                  &check.Data.HREF,
					ApplicantID:           &check.Data.ApplicantID,
					ApplicantProvidesData: &check.Data.ApplicantProvidesData,
				},
				Status: types.State(check.Status),
				Ts:     &tsAsInt64,
			}
			// Add reports
			for _, reportDto := range check.Data.Reports {
				tsAsInt64 := reportDto.Timestamp.Unix()
				var report = types.Report{
					ID:     reportDto.ID,
					Data:   string(reportDto.Data),
					Status: types.State(reportDto.Status),
					Ts:     &tsAsInt64,
					Review: &types.ReportReviewStatus{
						Resubmit: &reportDto.Review.Resubmit,
						Message:  &reportDto.Review.Message,
					},
				}
				data.Data.Reports = append(data.Data.Reports, &report)
			}
			// TODO: Tags connection

			// Add data to validation
			validation.Data = &data
		case models.PROOF:
			var proof models.Proof
			err := json.Unmarshal(validationDto.Data.Value, &proof)
			if err != nil {
				r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
				continue
			}
			tsAsInt64 := proof.Timestamp.Unix()
			var data = types.Proof{
				ID:   proof.ID,
				Type: types.ProofType(proof.Type),
				Data: string(proof.Data),
				Review: &types.ReportReviewStatus{
					Resubmit: &proof.Review.Resubmit,
					Message:  &proof.Review.Message,
				},
				Status: types.State(proof.Status),
				Ts:     &tsAsInt64,
			}

			// Add data to validation
			validation.Data = &data
		}

		// Append validation
		cdd.Validations = append(cdd.Validations, &validation)
	}

	return &cdd, nil
}

// TODO: Refactor this function to use it on the hydrateCDD
// It is possible by introducing a goto statement but arguable, though.
func (r *queryResolver) validation(ctx context.Context, validationDto *pb.Validation, dataResolver *DataResolver) *types.Validation {
	tsAsInt64 := int64(validationDto.Ts)
	//Build Validation Action
	actions := make([]*types.Action, len(validationDto.Actions))

	for index, action := range validationDto.Actions {
		person, err := getPerson(action.Reporter)
		if err != nil {
			r.logger.Error("get person data", zap.Error(err))
		}
		actions[index] = &types.Action{
			ID:       action.Id,
			Reporter: person,
			Notes:    action.Notes,
			Status:   action.Status,
			Ts:       int64(action.Ts),
		}
	}
	validation := types.Validation{
		ID:             validationDto.Id,
		ValidationType: types.ValidationType(validationDto.ValidationType),
		Status:         types.State(validationDto.Status),
		Approved:       &validationDto.Approved,
		Organisation: &types.Organisation{
			ID:   validationDto.Organisation.Id,
			Name: &validationDto.Organisation.Name,
		},
		Actions: actions,
		Ts:      &tsAsInt64,
	}
	// Fill validation Data
	switch validationDto.Data.TypeUrl {
	case models.SCREEN:
		var screen models.Screen
		err := json.Unmarshal(validationDto.Data.Value, &screen)
		if err != nil {
			r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
		}
		tsAsInt64 := screen.Timestamp.Unix()
		var data = types.Screen{
			ID:     screen.ID,
			Data:   string(screen.Data),
			Status: types.State(screen.Status),
			Ts:     &tsAsInt64,
		}

		// Add data to validation
		validation.Data = &data
	case models.CHECK:
		var check models.Check
		err := json.Unmarshal(validationDto.Data.Value, &check)
		if err != nil {
			r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
		}
		pbOwner, err := r.personService.Person(ctx, &personService.PersonRequest{Id: validationDto.Applicant})
		var owner *types.Person
		if err != nil {
			r.logger.Error(errorGettingPersonMsg, zap.Error(err))
		} else {
			owner, err = getPerson(pbOwner)
			if err != nil {
				r.logger.Error("resolve person data", zap.Error(err))
			}
		}

		tsAsInt64 := check.Timestamp.Unix()
		createdAtAsString := check.Data.CreatedAt.Format(time.RFC3339)
		var data = types.Check{
			ID: check.ID,
			Data: &types.CheckData{
				ID:                    check.Data.ID,
				CreatedAt:             &createdAtAsString,
				Status:                types.State(check.Data.Status),
				Sandbox:               &check.Data.Sandbox,
				ResultsURI:            &check.Data.ResultsURI,
				FormURI:               &check.Data.FormURI,
				Paused:                &check.Data.Paused,
				Version:               &check.Data.Version,
				Href:                  &check.Data.HREF,
				ApplicantID:           &check.Data.ApplicantID,
				ApplicantProvidesData: &check.Data.ApplicantProvidesData,
			},
			Status: types.State(check.Status),
			Ts:     &tsAsInt64,
			Owner:  owner,
			Organisation: &types.Organisation{
				ID:   validationDto.Organisation.Id,
				Name: &validationDto.Organisation.Name,
			},
		}
		// Add reports
		reports := make([]*types.Report, len(check.Data.Reports))
		for i := range check.Data.Reports {
			reportDto := check.Data.Reports[i]
			tsAsInt64 := reportDto.Timestamp.Unix()

			organization, err := r.dataStore.GetOrganization(reportDto.Organisation)
			if err != nil {
				r.logger.Error("get organization data", zap.Error(err))
				organization = &models.Organization{}
			}
			var report = types.Report{
				ID:     reportDto.ID,
				Data:   string(reportDto.Data),
				Status: types.State(reportDto.Status),
				Ts:     &tsAsInt64,
				Review: &types.ReportReviewStatus{
					Resubmit: &reportDto.Review.Resubmit,
					Message:  &reportDto.Review.Message,
				},
				Organisation: &types.Organisation{
					ID:   organization.ID,
					Name: &organization.Name,
				},
			}
			reports[i] = &report
		}
		data.Data.Reports = reports
		// TODO: Tags connection

		// Add data to validation
		validation.Data = &data
	case models.PROOF:
		var proof models.Proof
		err := json.Unmarshal(validationDto.Data.Value, &proof)
		if err != nil {
			r.logger.Error(errorMarshallingScreenValidation, zap.Error(err))
		}
		tsAsInt64 := proof.Timestamp.Unix()
		var data = types.Proof{
			ID:   proof.ID,
			Type: types.ProofType(proof.Type),
			Data: string(proof.Data),
			Review: &types.ReportReviewStatus{
				Resubmit: &proof.Review.Resubmit,
				Message:  &proof.Review.Message,
			},
			Status: types.State(proof.Status),
			Ts:     &tsAsInt64,
		}

		// Add data to validation
		validation.Data = &data
	}
	return &validation
}

func personWithCdd(from *pb.Person) (*types.Person, error) {
	person, err := getPerson(from)
	if err != nil {
		return nil, err
	}
	if from.Cdd != nil {
		ts := int64(from.Cdd.Ts)
		person.Cdd = &types.Cdd{
			Status: types.State(from.Cdd.Status),
			Ts:     &ts,
		}
	}
	return person, nil
}

func (r *queryResolver) processCddChunk(ctx context.Context, cdds []*pb.Cdd, dataResolver *DataResolver, dataConverter *DataConverter) ([]*types.Cdd, error) {
	cddsValues := make([]*types.Cdd, len(cdds))
	for i, next := range cdds {
		validations := make([]*types.Validation, 0)
		for _, validation := range next.Validations {
			modelValidation, err := dataConverter.ProtoValidationToModel(validation)
			if err != nil {
				r.logger.With(zap.Error(err)).Error("cannot convert validation")
				continue
			}
			nextValidation, err := dataResolver.ResolveValidation(*modelValidation)
			if err != nil {
				r.logger.With(zap.Error(err)).Error("cannot resolve validation data")
				continue
			}
			validations = append(validations, nextValidation)
		}

		owner, err := dataResolver.ResolvePerson(next.Owner, nil)
		if err != nil {
			return nil, err
		}
		cddValue := &types.Cdd{
			ID:          next.Id,
			Owner:       owner,
			Watchlist:   &next.Watchlist,
			Details:     &next.Details,
			Status:      types.State(next.Status),
			Onboard:     &next.Onboard,
			Version:     Int64(int64(next.Version)),
			Validations: validations,
			Active:      &next.Active,
			Ts:          Int64(int64(next.Ts)),
		}

		cddsValues[i] = cddValue
	}
	return cddsValues, nil
}

// TODO: Converts from cursor-based pagination to number based pagination
func (r *queryResolver) perPageCddsQuery(first *int64, after *string, last *int64, before *string) int64 {
	if after == nil && before == nil && first != nil {
		return *first
	}
	return 100
}
func String(s string) *string {
	return &s
}

func Int64(i int64) *int64 {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}
