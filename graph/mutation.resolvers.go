package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"ms.api/graph/generated"
	"ms.api/libs/validator/datevalidator"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/libs/validator/phonenumbervalidator"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/identityService"
	"ms.api/protos/pb/onboardingService"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *mutationResolver) CreatePhone(ctx context.Context, phone string, device types.DeviceInput) (*types.Response, error) {
	if err := phonenumbervalidator.ValidatePhoneNumber(phone); err != nil {
		r.logger.Error("failed to validate phone number",
			zap.Error(err),
			zap.String("phone", phone),
		)
		return nil, err
	}

	result, err := r.onBoardingService.CreatePhone(ctx,
		&onboardingService.CreatePhoneRequest{
			PhoneNumber: phone,
			Device: &protoTypes.Device{
				Identifier: device.Identifier,
				Brand:      device.Brand,
				Os:         device.Os,
			},
		},
	)
	if err != nil {
		r.logger.Error("error calling onBoardingService.CreatePhone()", zap.Error(err))
		return nil, err
	}
	return &types.Response{Message: result.Message, Success: true, Token: &result.Token}, nil
}

func (r *mutationResolver) ConfirmPhone(ctx context.Context, token string, code string) (*types.Response, error) {
	resp, err := r.onBoardingService.VerifySmsOtp(context.Background(), &onboardingService.OtpVerificationRequest{
		Token: token, Code: code,
	})
	if err != nil {
		r.logger.Error("error calling onBoardingService.ConfirmPhone()", zap.Error(err))
		return nil, err
	}
	return &types.Response{Success: resp.Match, Message: resp.Message}, nil
}

func (r *mutationResolver) Signup(ctx context.Context, token string, email string, passcode string) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(email); err != nil {
		return nil, err
	}

	result, err := r.onBoardingService.CreatePerson(ctx, &onboardingService.CreatePersonRequest{
		Email:    email,
		Passcode: passcode,
		Token:    token,
	})
	if err != nil {
		r.logger.Error("error calling onboardingService.CreatePerson.", zap.Error(err))
		return nil, err
	}

	return &types.AuthResponse{
		Message: "successful",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    result.JwtToken,
			Refresh: result.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) Registration(ctx context.Context, personid string, person types.PersonInput, address types.AddressInput) (*types.Person, error) {
	personId, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	if err := datevalidator.ValidateDob(person.Dob); err != nil {
		return nil, err
	}
	if err := r.validateAddress(address); err != nil {
		return nil, err
	}
	if *address.Postcode == "" {
		*address.Postcode = "NA"
	}

	payload := onboardingService.UpdatePersonRequest{
		PersonId: personId.PersonId,
		Address: &onboardingService.InputAddress{
			Postcode: *address.Postcode, Street: *address.Street,
			City: *address.City, Country: *address.Country,
		},
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Dob:       person.Dob,
	}
	res, err := r.onBoardingService.UpdatePersonBiodata(context.Background(), &payload)
	if err != nil {
		return nil, err
	}
	identities := make([]*types.Identity, 0)
	emails := make([]*types.Email, 0)
	phones := make([]*types.Phone, 0)
	addresses := make([]*types.Address, 0)
	for _, id := range res.Identities {
		identities = append(identities, &types.Identity{
			ID:             id.Id,
			Owner:          id.Owner,
			Nickname:       &id.Nickname,
			Active:         &id.Active,
			Authentication: &id.Authentication,
		})
	}
	for _, email := range res.Emails {
		emails = append(emails, &types.Email{
			Value:    email.Value,
			Verified: email.Verified,
		})
	}
	for _, phone := range res.Phones {
		phones = append(phones, &types.Phone{
			Value:    phone.Number,
			Verified: phone.Verified,
		})
	}
	for _, addr := range res.Addresses {
		addresses = append(addresses, &types.Address{
			Street:   &addr.Street,
			State:    &addr.State,
			Postcode: &addr.Postcode,
			Country:  &types.Country{CountryName: addr.Country},
		})
	}
	nationality := make([]*string, 0)
	for _, next := range res.Nationality {
		nationality = append(nationality, &next)
	}
	return &types.Person{
		ID:               res.Id,
		Title:            &res.Title,
		FirstName:        res.FirstName,
		LastName:         res.LastName,
		MiddleName:       &res.MiddleName,
		Phones:           phones,
		Emails:           emails,
		Dob:              res.Dob,
		CountryResidence: &res.CountryResidence,
		Nationality:      nationality,
		Addresses:        addresses,
		Identities:       identities,
		Ts:               int64(res.Ts),
	}, nil
}

func (r *mutationResolver) IntendedActivities(ctx context.Context, activities []string) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	req := &onboardingService.RoavaReasonsRequest{
		PersonId: claims.PersonId,
		Reasons:  activities,
	}

	resp, err := r.onBoardingService.AddReasonsForUsingRoava(ctx, req)
	if err != nil {
		r.logger.Error("Adding reasons for using roava failed", zap.Error(err))
		return nil, fmt.Errorf("an error ocurred saving the reasons of use, pls try again")
	}

	r.logger.Info("Adding reasons for using roava succeed")

	return &types.Response{Message: resp.Message, Success: true}, nil
}

func (r *mutationResolver) CreateApplication(ctx context.Context, applicant types.ApplicantInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	req := &onboardingService.CreateOnfidoApplicantRequest{
		PersonId: claims.PersonId,
	}
	resp, err := r.onBoardingService.CreateOnfidoApplicant(ctx, req)
	if err != nil {
		r.logger.Error("create applicant request failed", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: "successful",
		Success: true,
		Token:   &resp.Token,
	}, nil
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, email string, code string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResendOtp(ctx context.Context, phone string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResendEmailMagicLInk(ctx context.Context, email string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, credentials types.AuthInput) (*types.AuthResponse, error) {
	if err := emailvalidator.Validate(credentials.Email); err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", credentials.Email))
		return nil, err
	}
	// TODO: change authService.LoginRequest{}.Tokens to slice datatype
	req := &authService.LoginRequest{
		Email:    credentials.Email,
		Passcode: credentials.Passcode,
		Device: &protoTypes.Device{
			Os:         credentials.Device.Os,
			Brand:      credentials.Device.Brand,
			Identifier: credentials.Device.Identifier,
		}}
	resp, err := r.authService.Login(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.Login() failed: %v", err))
		return nil, err
	}
	return &types.AuthResponse{
		Message: "Successful!",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    resp.AccessToken,
			Refresh: resp.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (*types.AuthResponse, error) {
	req := &authService.RefreshTokenRequest{RefreshToken: token}
	resp, err := r.authService.RefreshToken(ctx, req)
	if err != nil {
		r.logger.Info(fmt.Sprintf("authService.RefreshToken() failed: %v", err))
		return nil, err
	}
	return &types.AuthResponse{
		Message: "Successful!",
		Success: true,
		Tokens: &types.AuthTokens{
			Auth:    resp.Token,
			Refresh: resp.RefreshToken,
		},
	}, nil
}

func (r *mutationResolver) UpdateDeviceToken(ctx context.Context, token []*types.DeviceTokenInput) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	// Get tokens
	var requestTokens []*identityService.DeviceTokens
	for _, token := range token {
		if token != nil {
			requestTokens = append(requestTokens, &identityService.DeviceTokens{
				Type:  string(token.Type),
				Value: token.Value,
			})
		}
	}
	response, err := r.identityService.UpdateDeviceTokens(ctx, &identityService.UpdateDeviceTokensRequest{
		DeviceId:   claims.DeviceId,
		IdentityId: claims.IdentityId,
		Tokens:     requestTokens,
	})
	if err != nil {
		r.logger.Error("error calling identityService.UpdateDeviceTokens()", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: "successful",
		Success: response.Success,
	}, nil
}

func (r *mutationResolver) ResetPasscode(ctx context.Context, credentials *types.AuthInput, token string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SubmitApplication(ctx context.Context) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptTerms(ctx context.Context, documents []*string) (*types.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
