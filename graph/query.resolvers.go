package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"ms.api/graph/generated"
	"ms.api/protos/pb/authService"
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) People(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.PersonConnection, error) {
	panic(fmt.Errorf("not implemented"))
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
