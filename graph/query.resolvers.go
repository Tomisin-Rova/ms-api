package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"ms.api/protos/pb/personService"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"ms.api/graph/connections"
	"ms.api/graph/generated"
	"ms.api/graph/models"
	"ms.api/protos/pb/authService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) Node(ctx context.Context, id string) (types.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*types.Person, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	person, err := r.authService.GetPerson(ctx, &authService.GetPersonRequest{
		PersonId:   claims.PersonId,
		IdentityId: claims.IdentityId,
		DeviceId:   claims.DeviceId,
	})
	if err != nil {
		r.logger.Error("failed to get person", zap.Error(err))
		return nil, err
	}
	p := &types.Person{}
	if err := copier.Copy(p, person); err != nil {
		r.logger.Error("copier failed", zap.Error(err))
		return nil, errors.New("failed to read profile information. please retry")
	}
	p.Dob = person.DOB
	return p, nil
}

func (r *queryResolver) Person(ctx context.Context, id string) (*types.Person, error) {
	person, err := r.personService.Person(ctx, &personService.PersonRequest{Id: id})
	if err != nil {
		return nil, err
	}
	identities := make([]*types.Identity, 0)
	emails := make([]*types.Email, 0)
	phones := make([]*types.Phone, 0)
	addresses := make([]*types.Address, 0)

	for _, id := range person.Identities {
		identities = append(identities, &types.Identity{
			ID:             id.Id,
			Owner:          id.Owner,
			Nickname:       &id.Nickname,
			Active:         &id.Active,
			Authentication: &id.Authentication,
		})
	}
	for _, email := range person.Emails {
		emails = append(emails, &types.Email{
			Value:    email.Value,
			Verified: email.Verified,
		})
	}
	for _, phone := range person.Phones {
		phones = append(phones, &types.Phone{
			Value:    phone.Number,
			Verified: phone.Verified,
		})
	}
	for _, addr := range person.Addresses {
		addresses = append(addresses, &types.Address{
			Street:   &addr.Street,
			Postcode: &addr.Postcode,
			Country:  &types.Country{CountryName: addr.Country},
			City:     &addr.Town,
		})
	}
	nationalities := make([]*string, 0)
	for _, next := range person.Nationality {
		nationalities = append(nationalities, &next)
	}
	return &types.Person{
		ID:               person.Id,
		Title:            &person.Title,
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		MiddleName:       &person.MiddleName,
		Phones:           phones,
		Emails:           emails,
		Dob:              person.Dob,
		CountryResidence: &person.CountryResidence,
		Nationality:      nationalities,
		Addresses:        addresses,
		Identities:       identities,
		Ts:               int64(person.Ts),
	}, nil
}

func (r *queryResolver) People(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.PersonConnection, error) {
	res, err := r.personService.People(ctx, &personService.PeopleRequest{
		Page:    1,
		PerPage: 50,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Person, 0)
	for _, next := range res.Persons {
		person := next
		identities := make([]*types.Identity, 0)
		emails := make([]*types.Email, 0)
		phones := make([]*types.Phone, 0)
		addresses := make([]*types.Address, 0)

		for _, id := range person.Identities {
			identities = append(identities, &types.Identity{
				ID:             id.Id,
				Owner:          id.Owner,
				Nickname:       &id.Nickname,
				Active:         &id.Active,
				Authentication: &id.Authentication,
			})
		}
		for _, email := range person.Emails {
			emails = append(emails, &types.Email{
				Value:    email.Value,
				Verified: email.Verified,
			})
		}
		for _, phone := range person.Phones {
			phones = append(phones, &types.Phone{
				Value:    phone.Number,
				Verified: phone.Verified,
			})
		}
		for _, addr := range person.Addresses {
			addresses = append(addresses, &types.Address{
				Street:   &addr.Street,
				Postcode: &addr.Postcode,
				Country:  &types.Country{CountryName: addr.Country},
				City:     &addr.Town,
			})
		}
		nationalities := make([]*string, 0)
		for _, next := range person.Nationality {
			nationalities = append(nationalities, &next)
		}
		p := &types.Person{
			ID:               person.Id,
			Title:            &person.Title,
			FirstName:        person.FirstName,
			LastName:         person.LastName,
			MiddleName:       &person.MiddleName,
			Phones:           phones,
			Emails:           emails,
			Dob:              person.Dob,
			CountryResidence: &person.CountryResidence,
			Nationality:      nationalities,
			Addresses:        addresses,
			Identities:       identities,
			Ts:               int64(person.Ts),
		}
		data = append(data, p)
	}
	return &types.PersonConnection{Nodes: data}, nil
}

func (r *queryResolver) Identity(ctx context.Context, id string) (*types.Identity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Identities(ctx context.Context) ([]*types.Identity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Address(ctx context.Context, id string) (*types.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Addresses(ctx context.Context) ([]*types.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AddressLookup(ctx context.Context, text *string, first *int64, after *string, last *int64, before *string) (*types.AddressConnection, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	addresses, err := r.onBoardingService.AddressLookup(ctx, &onboardingService.AddressLookupRequest{
		Text: *text,
	})

	if err != nil {
		return nil, err
	}

	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(address *types.Address, offset int) connections.Edge {
		return types.AddressEdge{
			Node:   address,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.AddressEdge, nodes []*types.Address, info *types.PageInfo, totalCount int) (*types.AddressConnection, error) {
		var addressNodes []*types.Address
		addressNodes = append(addressNodes, nodes...)

		return &types.AddressConnection{
			Edges:      edges,
			Nodes:      addressNodes,
			PageInfo:   info,
			TotalCount: int64(totalCount),
		}, nil
	}

	var addressRes []*types.Address
	for _, c := range addresses.Addresses {
		address := &types.Address{
			Street:   &c.Street,
			City:     &c.Town,
			State:    &c.State,
			Postcode: &c.Postcode,
			Country: &types.Country{
				CountryName: c.Country,
			},
		}

		addressRes = append(addressRes, address)
	}

	return connections.AddressLookupCon(addressRes, edger, conn, input)
}

func (r *queryResolver) Device(ctx context.Context, identifier string) (*types.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Devices(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.DeviceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Auths(ctx context.Context) ([]*types.Auth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Organisation(ctx context.Context, id string) (*types.Organisation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Organisations(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.OrganisationConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Verification(ctx context.Context, code string) (*types.Verification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Verifications(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.VerificationConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Country(ctx context.Context, code string) (*types.Country, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Countries(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.CountryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Currency(ctx context.Context, code string) (*types.Currency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Currencies(ctx context.Context, supported *bool, first *int64, after *string, last *int64, before *string) (*types.CurrencyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Industry(ctx context.Context, code string) (*types.Industry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Industries(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.IndustryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Activity(ctx context.Context, id string) (*types.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Activities(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ActivityConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Message(ctx context.Context, id string) (*types.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Messages(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.MessageConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quote(ctx context.Context, id string) (*types.Quote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quotes(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.QuoteConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Price(ctx context.Context, pair *string, ts *int64) (*types.Fx, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Prices(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.FxConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tag(ctx context.Context, id string) (*types.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.TagConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cdd(ctx context.Context, id string) (*types.Cdd, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cdds(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string) (*types.CDDConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Check(ctx context.Context, id string) (*types.Check, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Checks(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.CheckConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Screen(ctx context.Context, id string) (*types.Screen, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Screens(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ScreenConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OnfidoReport(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ComplyAdvReport(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Task(ctx context.Context, id string) (*types.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comment(ctx context.Context, id string) (*types.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.CommentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*types.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ProductConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context, id string) (*types.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.AccountConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Transaction(ctx context.Context, id string) (*types.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Transactions(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.TransactionConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Acceptance(ctx context.Context, id string) (*types.Acceptance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Acceptances(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.AcceptanceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
