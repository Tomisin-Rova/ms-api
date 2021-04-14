package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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
	}
}

func ConnectServiceDependencies(secrets *config.Secrets) (*ResolverOpts, error) {
	opts := &ResolverOpts{}

	// OnBoarding
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err := dialRPC(ctx, secrets.OnboardingServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnboardingServiceURL)
	}
	opts.OnBoardingService = onboardingService.NewOnBoardingServiceClient(connection)

	// OnFido
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	connection, err = dialRPC(ctx, secrets.OnfidoServiceURL)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, secrets.OnfidoServiceURL)
	}
	opts.OnfidoClient = onfidoService.NewOnfidoServiceClient(connection)

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
	var person = types.Person{
		ID:               from.Id,
		Title:            &from.Title,
		FirstName:        from.FirstName,
		LastName:         from.LastName,
		MiddleName:       &from.MiddleName,
		Dob:              from.Dob,
		Ts:               from.Ts,
		CountryResidence: &from.CountryResidence,
	}
	// TODO: Fill other attributes

	addresses := make([]*types.Address, 0)
	for _, addr := range from.Addresses {
		addresses = append(addresses, &types.Address{
			Street:   &addr.Street,
			Postcode: &addr.Postcode,
			Country: &types.Country{
				CountryName: addr.Country,
			},
			City: &addr.City,
		})
	}
	person.Addresses = addresses

	// Add Phones
	phones := make([]*types.Phone, 0)
	for _, ph := range from.Phones {
		phones = append(phones, &types.Phone{
			Value:    ph.Number,
			Verified: ph.Verified,
		})
	}
	person.Phones = phones

	// Add Emails
	emails := make([]*types.Email, 0)
	for _, em := range from.Emails {
		emails = append(emails, &types.Email{
			Value:    em.Value,
			Verified: em.Verified,
		})
	}
	person.Emails = emails

	// Add Activity
	activities := make([]*types.Activity, 0)
	for _, ac := range from.Activities {
		activities = append(activities, &types.Activity{
			ID:            ac.Id,
			Description:   ac.Description,
			RiskWeighting: int64(ac.RiskWeighting),
			Supported:     &ac.Supported,
			Archived:      &ac.Archived,
			Ts:            &ac.Ts,
		})
	}
	person.Activities = activities

	return &person, nil
}

func (r *queryResolver) hydrateCDD(cddDto *pb.Cdd) *types.Cdd {
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
		validation := types.Validation{
			ID:             validationDto.Id,
			ValidationType: types.ValidationType(validationDto.ValidationType),
			Status:         types.State(validationDto.Status),
			Approved:       &validationDto.Approved,
			Ts:             &tsAsInt64,
		}
		// Fill validation Data
		switch validationDto.Data.TypeUrl {
		case models.SCREEN:
			var screen models.Screen
			err := json.Unmarshal(validationDto.Data.Value, &screen)
			if err != nil {
				r.logger.Error("marshall screen validation", zap.Error(err))
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
				r.logger.Error("marshall screen validation", zap.Error(err))
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
				r.logger.Error("marshall screen validation", zap.Error(err))
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

	return &cdd
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
