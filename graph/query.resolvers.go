package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"
	terror "github.com/roava/zebra/errors"
	"go.uber.org/zap"
	"ms.api/graph/connections"
	"ms.api/graph/generated"
	"ms.api/graph/models"
	mainErrors "ms.api/libs/errors"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/accountService"
	"ms.api/protos/pb/cddService"
	"ms.api/protos/pb/onboardingService"
	"ms.api/protos/pb/paymentService"
	"ms.api/protos/pb/personService"
	"ms.api/server/http/middlewares"
	"ms.api/types"
)

func (r *queryResolver) Me(ctx context.Context) (*types.Person, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	personDto, err := r.personService.Person(ctx, &personService.PersonRequest{
		Id: claims.PersonId,
	})
	if err != nil {
		r.logger.Error(errorGettingPersonMsg, zap.Error(err))
		return nil, err
	}
	person := getPerson(personDto)
	// Add CDD to response
	cddDto, err := r.cddService.GetCDDByOwner(ctx, &cddService.GetCDDByOwnerRequest{
		PersonId: claims.PersonId,
	})
	if err != nil {
		// If error it's CddNotFound don't return error
		newTerror := mainErrors.NewFromGrpc(err)
		if newTerror == nil || newTerror.Code() != mainErrors.CddNotFound {
			r.logger.Error("get cdd", zap.Error(err))
			return nil, err
		}

		r.logger.Info("no cdd found", zap.String("owner", claims.PersonId))
	}
	dataConverter := NewDataConverter(r.logger)
	person.Cdd = dataConverter.makeCdd(cddDto)

	for i := range person.Identities {
		person.Identities[i].Owner = person
	}

	return person, nil
}

func (r *queryResolver) MeStaff(ctx context.Context) (*types.Staff, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	req := &personService.StaffRequest{
		Id: claims.PersonId,
	}
	staffDto, err := r.personService.GetStaffById(ctx, req)
	if err != nil {
		r.logger.Error(errorGettingPersonMsg, zap.Error(err))
		return nil, err
	}

	emails := make([]*types.Email, len(staffDto.Emails))
	for i, email := range staffDto.Emails {
		emails[i] = &types.Email{
			Value:    email.Value,
			Verified: email.Verified,
		}
	}

	phones := make([]*types.Phone, len(staffDto.Phones))
	for i, phone := range staffDto.Phones {
		phones[i] = &types.Phone{
			Value:    phone.Number,
			Verified: phone.Verified,
		}
	}

	identities := make([]*types.Identity, len(staffDto.Identities))
	for i, identity := range staffDto.Identities {
		org := &types.Organisation{
			ID:   identity.Organisation.Id,
			Name: &identity.Organisation.Name,
		}
		identities[i] = &types.Identity{
			ID:             identity.Id,
			Active:         &identity.Active,
			Authentication: &identity.Authentication,
			Credentials: &types.Credentials{
				Identifier:   identity.Credentials.Identifier,
				RefreshToken: &identity.Credentials.RefreshToken,
			},
			Organisation: org,
			Ts:           identity.Ts,
		}
		identities[i].Owner = &types.Person{
			ID:               identity.Owner.Id,
			Title:            &identity.Owner.Title,
			FirstName:        identity.Owner.FirstName,
			LastName:         identity.Owner.LastName,
			MiddleName:       &identity.Owner.MiddleName,
			Dob:              identity.Owner.Dob,
			Employer:         org,
			Ts:               identity.Owner.Ts,
			CountryResidence: &identity.Owner.CountryResidence,
		}
	}

	return &types.Staff{
		ID:         staffDto.Id,
		FirstName:  staffDto.FirstName,
		LastName:   staffDto.LastName,
		Status:     types.StaffStatus(staffDto.Status),
		Emails:     emails,
		Phones:     phones,
		Identities: identities,
	}, nil
}

func (r *queryResolver) Person(ctx context.Context, id string) (*types.Person, error) {
	person, err := r.personService.Person(ctx, &personService.PersonRequest{Id: id})
	if err != nil {
		return nil, err
	}

	p := getPerson(person)

	// Add CDD to response
	cddDto, err := r.cddService.GetCDDByOwner(ctx, &cddService.GetCDDByOwnerRequest{
		PersonId: p.ID,
	})
	if err != nil {
		// If error it's CddNotFound don't return error
		newTerror := mainErrors.NewFromGrpc(err)
		if newTerror == nil || newTerror.Code() != mainErrors.CddNotFound {
			r.logger.Error("get cdd", zap.Error(err))
			return nil, err
		}

		r.logger.Info("no cdd found", zap.String("owner", p.ID))
	}

	dataConverter := NewDataConverter(r.logger)
	p.Cdd = dataConverter.makeCdd(cddDto)

	return p, nil
}

func (r *queryResolver) People(ctx context.Context, keywords *string, first *int64, after *string, last *int64, before *string, onboarded *bool) (*types.PersonConnection, error) {
	var kw string
	if keywords != nil {
		kw = *keywords
	}
	onboardedStatus := IgnoreOnboardFilter
	if onboarded != nil && *onboarded {
		onboardedStatus = Onboarded
	} else if onboarded != nil {
		onboardedStatus = NotOnboarded
	}
	res, err := r.personService.People(ctx, &personService.PeopleRequest{
		Page:      1,
		PerPage:   100,
		Keywords:  kw,
		Onboarded: string(onboardedStatus),
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Person, len(res.Persons))

	for i, person := range res.Persons {
		pto, err := personWithCdd(person)
		if err != nil {
			return nil, err
		}
		data[i] = pto
	}

	if err != nil {
		return nil, err
	}

	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(person *types.Person, offset int) connections.Edge {
		return types.PersonEdge{
			Node:   person,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.PersonEdge, nodes []*types.Person, info *types.PageInfo, totalCount int) (*types.PersonConnection, error) {
		var personNodes []*types.Person
		personNodes = append(personNodes, nodes...)
		count := int64(totalCount)
		return &types.PersonConnection{
			Edges:      edges,
			Nodes:      personNodes,
			PageInfo:   info,
			TotalCount: &count,
		}, nil
	}
	return connections.PeopleLookupCon(data, edger, conn, input)
}

func (r *queryResolver) Identity(ctx context.Context, id string) (*types.Identity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Identities(ctx context.Context) ([]*types.Identity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (*bool, error) {
	newEmail, err := emailvalidator.Validate(email)
	if err != nil {
		r.logger.Info("invalid email supplied", zap.String("email", email))
		return nil, err
	}
	resp, err := r.onBoardingService.CheckEmailExistence(ctx, &onboardingService.CheckEmailExistenceRequest{Email: newEmail})
	if err != nil {
		r.logger.Error("error calling onboardingService.checkEmailExistence()", zap.Error(err))
		return nil, err
	}
	return &resp.Exists, nil
}

func (r *queryResolver) Address(ctx context.Context, id string) (*types.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Addresses(ctx context.Context) (*types.AddressConnection, error) {
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
			TotalCount: Int64(int64(totalCount)),
		}, nil
	}

	var addressRes []*types.Address
	for _, c := range addresses.Addresses {
		lon, err := strconv.ParseFloat(c.Longitude, 64)
		if err != nil {
			r.logger.Info("can't convert longitude value", zap.Error(err))
		}
		lat, err := strconv.ParseFloat(c.Latitude, 64)
		if err != nil {
			r.logger.Info("can't convert latitude value", zap.Error(err))
		}
		address := &types.Address{
			Street:   &c.Summaryline,
			City:     &c.Posttown,
			State:    &c.County,
			Postcode: &c.Postcode,
			Location: &types.Location{
				Longitude: &lon,
				Latitude:  &lat,
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

func (r *queryResolver) Activities(ctx context.Context, supported *bool) ([]*types.Activity, error) {
	reason, err := r.onBoardingService.FetchReasons(ctx, &onboardingService.FetchReasonsRequest{
		Supported: supported != nil && *supported,
	})
	if err != nil {
		r.logger.Error(errorGettingPersonMsg, zap.Error(err))
		return nil, err
	}
	p := make([]*types.Activity, 0)

	for _, v := range reason.Reasons {
		p = append(p, &types.Activity{
			ID:          v.Id,
			Description: v.Description,
		})
	}

	return p, nil
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

func (r *queryResolver) Cdds(ctx context.Context, keywords *string, status []types.State, first *int64, after *string, last *int64, before *string) (*types.CDDConnection, error) {
	dataConverter := NewDataConverter(r.logger)
	perPage := r.perPageCddsQuery(first, after, last, before)

	req := &cddService.CDDSRequest{
		Page:    1,
		PerPage: perPage,
		Status:  dataConverter.StateToStringSlice(status),
	}
	if keywords != nil {
		req.Keywords = *keywords
	}
	resp, err := r.cddService.CDDS(context.Background(), req)
	if err != nil {
		r.logger.With(zap.Error(err)).Error("failed to fetch cdds")
		return nil, terror.NewTerror(7013, "InternalError", "failed to load CDDs data. Internal system error", "internal system error")
	}

	// dataResolver := NewDataResolver(r.dataStore, r.logger)
	cdds := resp.Results
	cddsResult := make([]*types.Cdd, len(cdds))
	for i, cdd := range cdds {
		cddsResult[i] = dataConverter.makeCdd(cdd)
	}

	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(cdd *types.Cdd, offset int) connections.Edge {
		return types.CDDEdge{
			Node:   cdd,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.CDDEdge, nodes []*types.Cdd, info *types.PageInfo, totalCount int) (*types.CDDConnection, error) {
		var cddNodes []*types.Cdd
		cddNodes = append(cddNodes, nodes...)
		count := int64(totalCount)
		return &types.CDDConnection{
			Edges:      edges,
			Nodes:      cddNodes,
			PageInfo:   info,
			TotalCount: &count,
		}, nil
	}
	return connections.CddLookupCon(cddsResult, edger, conn, input)
}

func (r *queryResolver) Validation(ctx context.Context, id string) (*types.Validation, error) {
	validationDto, err := r.cddService.GetValidationById(ctx, &cddService.GetValidationByIdRequest{
		ValidationId: id,
	})
	if err != nil {
		r.logger.Error("get validation", zap.Error(err))
		return nil, err
	}
	dataResolver := NewDataResolver(r.dataStore, r.logger)
	validation := r.validation(ctx, validationDto, dataResolver)

	return validation, nil
}

func (r *queryResolver) Validations(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ValidationConnection, error) {
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

func (r *queryResolver) Report(ctx context.Context, id string) (*types.Report, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reports(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ReportConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Proof(ctx context.Context, id string) (*types.Proof, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Proofs(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ProofConnection, error) {
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
	product, err := r.accountService.GetProduct(ctx, &accountService.GetProductRequest{Id: id})
	if err != nil {
		r.logger.Error("failed to get product", zap.Error(err))
		return nil, err
	}

	var productRes types.Product
	if err := r.mapper.Hydrate(product, &productRes); err != nil {
		err := mainErrors.Format(mainErrors.InternalErr, nil)
		r.logger.Error("debug", zap.Error(err))
		return nil, err
	}
	return &productRes, nil
}

func (r *queryResolver) Products(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ProductConnection, error) {
	products, err := r.accountService.GetProducts(ctx, &accountService.GetProductsRequest{
		Page:    1,
		PerPage: 100,
	})
	if err != nil {
		r.logger.Error("failed to get products", zap.Error(err))
		return nil, err
	}

	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(p *types.Product, offset int) connections.Edge {
		return types.ProductEdge{
			Node:   p,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.ProductEdge, nodes []*types.Product, info *types.PageInfo, totalCount int) (*types.ProductConnection, error) {
		var productNodes []*types.Product
		productNodes = append(productNodes, nodes...)

		return &types.ProductConnection{
			Edges:      edges,
			Nodes:      productNodes,
			PageInfo:   info,
			TotalCount: Int64(int64(totalCount)),
		}, nil
	}

	var productRes []*types.Product
	for _, p := range products.Products {
		var product types.Product
		if err := r.mapper.Hydrate(p, &product); err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, nil)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}

		productRes = append(productRes, &product)
	}

	return connections.ProductConnectionCon(productRes, edger, conn, input)
}

func (r *queryResolver) Account(ctx context.Context, id string) (*types.Account, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	preloads := r.preloader.GetPreloads(ctx)

	var opts struct {
		ProductRequested      bool
		TagsRequested         bool
		TransactionsRequested bool
	}
	for _, item := range preloads {
		if item == "product" {
			opts.ProductRequested = true
		}
		if item == "tags" {
			opts.TagsRequested = true
		}
		if item == "transactions" {
			opts.TransactionsRequested = true
		}
	}

	account, err := r.accountService.GetAccount(ctx, &accountService.GetAccountRequest{
		Id:         id,
		IdentityId: claims.IdentityId,
	})
	if err != nil {
		r.logger.Error("failed to get account", zap.Error(err))
		return nil, err
	}

	var accountRes types.Account
	if err := r.mapper.Hydrate(account, &accountRes); err != nil {
		err := mainErrors.Format(mainErrors.InternalErr, nil)
		r.logger.Error("debug", zap.Error(err))
		return nil, err
	}

	// Add Products if requested
	if opts.ProductRequested {
		product, err := r.accountService.GetProduct(ctx, &accountService.GetProductRequest{Id: account.Product})
		if err != nil {
			r.logger.Error("failed to get product", zap.Error(err))
			return nil, err
		}

		var productRes types.Product
		if err := r.mapper.Hydrate(product, &productRes); err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, nil)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}
		accountRes.Product = &productRes
	}

	// Add Transactions if requested
	if opts.TransactionsRequested && account.Transactions != nil {
		var transactionRes []*types.Transaction
		for _, p := range account.Transactions {
			var transaction types.Transaction
			if err := r.mapper.Hydrate(p, &transaction); err != nil {
				err := mainErrors.Format(mainErrors.InternalErr, nil)
				r.logger.Error("debug", zap.Error(err))
				return nil, err
			}

			transactionRes = append(transactionRes, &transaction)
		}

		transArgMap := r.preloader.GetArgMap(ctx, "Transactions")
		transConnInput := r.getConnInput(transArgMap)

		edger := func(p *types.Transaction, offset int) connections.Edge {
			return types.TransactionEdge{
				Node:   p,
				Cursor: connections.OffsetToCursor(offset),
			}
		}

		conn := func(edges []*types.TransactionEdge, nodes []*types.Transaction, info *types.PageInfo, totalCount int) (*types.TransactionConnection, error) {
			var transactionNodes []*types.Transaction
			transactionNodes = append(transactionNodes, nodes...)

			return &types.TransactionConnection{
				Edges:      edges,
				Nodes:      transactionNodes,
				PageInfo:   info,
				TotalCount: Int64(int64(totalCount)),
			}, nil
		}

		transConn, err := connections.TransactionConnectionCon(transactionRes, edger, conn, transConnInput)
		if err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, err)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}

		accountRes.Transactions = transConn
	}

	// Add Tags if requested
	if opts.TagsRequested && account.Tags != nil {
		var tagsRes []*types.Tag
		for _, p := range account.Tags {
			var tag types.Tag
			if err := r.mapper.Hydrate(p, &tag); err != nil {
				err := mainErrors.Format(mainErrors.InternalErr, nil)
				r.logger.Error("debug", zap.Error(err))
				return nil, err
			}

			tagsRes = append(tagsRes, &tag)
		}

		tagsArgMap := r.preloader.GetArgMap(ctx, "Tags")
		tagsConnInput := r.getConnInput(tagsArgMap)

		edger := func(p *types.Tag, offset int) connections.Edge {
			return types.TagEdge{
				Node:   p,
				Cursor: connections.OffsetToCursor(offset),
			}
		}

		conn := func(edges []*types.TagEdge, nodes []*types.Tag, info *types.PageInfo, totalCount int) (*types.TagConnection, error) {
			var TagNodes []*types.Tag
			TagNodes = append(TagNodes, nodes...)

			return &types.TagConnection{
				Edges:      edges,
				Nodes:      TagNodes,
				PageInfo:   info,
				TotalCount: Int64(int64(totalCount)),
			}, nil
		}

		tagsConn, err := connections.TagConnectionCon(tagsRes, edger, conn, tagsConnInput)
		if err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, err)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}

		accountRes.Tags = tagsConn
	}

	return &accountRes, nil
}

func (r *queryResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.AccountConnection, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	preloads := r.preloader.GetPreloads(ctx)

	var opts struct {
		ProductRequested      bool
		TagsRequested         bool
		TransactionsRequested bool
	}
	for _, item := range preloads {
		if item == "nodes.product" {
			opts.ProductRequested = true
		}
		if item == "nodes.tags" {
			opts.TagsRequested = true
		}
		if item == "nodes.transactions" {
			opts.TransactionsRequested = true
		}
	}

	accounts, err := r.accountService.GetAccounts(ctx, &accountService.GetAccountsRequest{
		IdentityId: claims.IdentityId,
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

	edger := func(account *types.Account, offset int) connections.Edge {
		return types.AccountEdge{
			Node:   account,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.AccountEdge, nodes []*types.Account, info *types.PageInfo, totalCount int) (*types.AccountConnection, error) {
		var accountNodes []*types.Account
		accountNodes = append(accountNodes, nodes...)

		return &types.AccountConnection{
			Edges:      edges,
			Nodes:      accountNodes,
			PageInfo:   info,
			TotalCount: Int64(int64(totalCount)),
		}, nil
	}

	var accountsRes []*types.Account
	for _, c := range accounts.Accounts {
		var account types.Account
		if err := r.mapper.Hydrate(c, &account); err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, nil)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}

		// Add Products if requested
		if opts.ProductRequested {
			product, err := r.accountService.GetProduct(ctx, &accountService.GetProductRequest{Id: account.Product.ID})
			if err != nil {
				r.logger.Error("failed to get product", zap.Error(err))
				return nil, err
			}

			var productRes types.Product
			if err := r.mapper.Hydrate(product, &productRes); err != nil {
				err := mainErrors.Format(mainErrors.InternalErr, nil)
				r.logger.Error("debug", zap.Error(err))
				return nil, err
			}
			account.Product = &productRes
		}

		// Add Transactions if requested
		if opts.TransactionsRequested && c.Transactions != nil {
			var transactionRes []*types.Transaction
			for _, p := range c.Transactions {
				var transaction types.Transaction
				if err := r.mapper.Hydrate(p, &transaction); err != nil {
					err := mainErrors.Format(mainErrors.InternalErr, nil)
					r.logger.Error("debug", zap.Error(err))
					return nil, err
				}

				transactionRes = append(transactionRes, &transaction)
			}

			transArgMap := r.preloader.GetArgMap(ctx, "Transactions")
			transConnInput := r.getConnInput(transArgMap)

			edger := func(p *types.Transaction, offset int) connections.Edge {
				return types.TransactionEdge{
					Node:   p,
					Cursor: connections.OffsetToCursor(offset),
				}
			}

			conn := func(edges []*types.TransactionEdge, nodes []*types.Transaction, info *types.PageInfo, totalCount int) (*types.TransactionConnection, error) {
				var transactionNodes []*types.Transaction
				transactionNodes = append(transactionNodes, nodes...)

				return &types.TransactionConnection{
					Edges:      edges,
					Nodes:      transactionNodes,
					PageInfo:   info,
					TotalCount: Int64(int64(totalCount)),
				}, nil
			}

			transConn, err := connections.TransactionConnectionCon(transactionRes, edger, conn, transConnInput)
			if err != nil {
				err := mainErrors.Format(mainErrors.InternalErr, err)
				r.logger.Error("debug", zap.Error(err))
				return nil, err
			}
			account.Transactions = transConn
		}

		// Add Tags if requested
		if opts.TagsRequested && c.Tags != nil {
			var tagsRes []*types.Tag
			for _, p := range c.Tags {
				var tag types.Tag
				if err := r.mapper.Hydrate(p, &tag); err != nil {
					err := mainErrors.Format(mainErrors.InternalErr, nil)
					r.logger.Error("debug", zap.Error(err))
					return nil, err
				}

				tagsRes = append(tagsRes, &tag)
			}

			tagsArgMap := r.preloader.GetArgMap(ctx, "Tags")
			tagsConnInput := r.getConnInput(tagsArgMap)

			edger := func(p *types.Tag, offset int) connections.Edge {
				return types.TagEdge{
					Node:   p,
					Cursor: connections.OffsetToCursor(offset),
				}
			}

			conn := func(edges []*types.TagEdge, nodes []*types.Tag, info *types.PageInfo, totalCount int) (*types.TagConnection, error) {
				var TagNodes []*types.Tag
				TagNodes = append(TagNodes, nodes...)

				return &types.TagConnection{
					Edges:      edges,
					Nodes:      TagNodes,
					PageInfo:   info,
					TotalCount: Int64(int64(totalCount)),
				}, nil
			}

			tagsConn, err := connections.TagConnectionCon(tagsRes, edger, conn, tagsConnInput)
			if err != nil {
				err := mainErrors.Format(mainErrors.InternalErr, err)
				r.logger.Error("debug", zap.Error(err))
				return nil, err
			}

			account.Tags = tagsConn
		}

		accountsRes = append(accountsRes, &account)
	}

	return connections.AccountConnectionCon(accountsRes, edger, conn, input)
}

func (r *queryResolver) Payee(ctx context.Context, id string) (*types.Payee, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	preloads := r.preloader.GetPreloads(ctx)

	var opts struct {
		PersonRequested   bool
		IdentityRequested bool
	}
	for _, item := range preloads {
		if item == "owner" {
			opts.IdentityRequested = true
		}
		if item == "owner.owner" {
			opts.PersonRequested = true
		}
	}

	payee, err := r.paymentService.GetPayee(ctx, &paymentService.GetPayeeRequest{
		PayeeId:    id,
		IdentityId: claims.IdentityId,
	})
	if err != nil {
		r.logger.Error("failed to get payee", zap.Error(err))
		return nil, err
	}

	payeeRes := &types.Payee{}
	if err := copier.Copy(payeeRes, &payee); err != nil {
		r.logger.Error(errorCopierFailedMsg, zap.Error(err))
		return nil, errors.New("failed to read payee information. please retry")
	}

	// update missing copier fields
	payeeRes.ID = payee.Id
	for index, account := range payee.Accounts {
		payeeRes.Accounts[index].ID = account.Id
	}

	if opts.IdentityRequested {
		identity, err := r.dataStore.GetIdentityById(claims.IdentityId)
		if err != nil {
			r.logger.Error("failed to get payee", zap.Error(err))
			return nil, err
		}

		identityRes := &types.Identity{}
		if err := copier.Copy(identityRes, &identity); err != nil {
			r.logger.Error(errorCopierFailedMsg, zap.Error(err))
			return nil, errors.New("failed to read identity information. please retry")
		}

		payeeRes.Owner = identityRes
	}

	if opts.PersonRequested {
		person, err := r.personService.Person(ctx, &personService.PersonRequest{Id: claims.PersonId})
		if err != nil {
			r.logger.Error(errorGettingPersonMsg, zap.Error(err))
			return nil, err
		}
		personRes := &types.Person{}
		if err := copier.Copy(payeeRes, &payee); err != nil {
			r.logger.Error(errorCopierFailedMsg, zap.Error(err))
			return nil, errors.New("failed to read person information. please retry")
		}
		// update missing copier fields
		personRes.ID = person.Id

		payeeRes.Owner.Owner = personRes
	}

	return payeeRes, nil
}

func (r *queryResolver) Payees(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.PayeeConnection, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	payees, err := r.paymentService.GetPayees(ctx, &paymentService.GetPayeesRequest{
		IdentityId: claims.IdentityId,
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

	edger := func(p *types.Payee, offset int) connections.Edge {
		return types.PayeeEdge{
			Node:   p,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.PayeeEdge, nodes []*types.Payee, info *types.PageInfo, totalCount int) (*types.PayeeConnection, error) {
		var payeeNodes []*types.Payee
		payeeNodes = append(payeeNodes, nodes...)

		return &types.PayeeConnection{
			Edges:      edges,
			Nodes:      payeeNodes,
			PageInfo:   info,
			TotalCount: Int64(int64(totalCount)),
		}, nil
	}

	var payeeRes []*types.Payee
	for _, p := range payees.Payee {
		payee := &types.Payee{}
		if err := copier.Copy(payee, &p); err != nil {
			r.logger.Error(errorCopierFailedMsg, zap.Error(err))
			return nil, errors.New("failed to read payee information. please retry")
		}

		// update missing copier fields
		payee.ID = p.Id
		for index, account := range p.Accounts {
			payee.Accounts[index].ID = account.Id
		}
		payeeRes = append(payeeRes, payee)
	}

	return connections.PayeeConnectionCon(payeeRes, edger, conn, input)
}

func (r *queryResolver) Transaction(ctx context.Context, id string) (*types.Transaction, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	transaction, err := r.accountService.GetTransaction(ctx, &accountService.GetTransactionRequest{
		Id: id,
	})
	if err != nil {
		r.logger.Error("failed to get transaction", zap.Error(err))
		return nil, err
	}
	var transactionRes types.Transaction
	if err := r.mapper.Hydrate(transaction, &transactionRes); err != nil {
		err := mainErrors.Format(mainErrors.InternalErr, nil)
		r.logger.Error("debug", zap.Error(err))
		return nil, err
	}

	var accountRes types.Account
	if err := r.mapper.Hydrate(transaction.Account, &accountRes); err != nil {
		err := mainErrors.Format(mainErrors.InternalErr, nil)
		r.logger.Error("debug", zap.Error(err))
		return nil, err
	}
	transactionRes.Account = &accountRes

	return &transactionRes, nil
}

func (r *queryResolver) Transactions(ctx context.Context, first *int64, after *string, last *int64, before *string, account string) (*types.TransactionConnection, error) {
	_, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	transactions, err := r.accountService.GetTransactions(ctx, &accountService.GetTransactionsRequest{
		Account: account,
	})
	if err != nil {
		r.logger.Error("failed to get transaction", zap.Error(err))
		return nil, err
	}

	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(p *types.Transaction, offset int) connections.Edge {
		return types.TransactionEdge{
			Node:   p,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conn := func(edges []*types.TransactionEdge, nodes []*types.Transaction, info *types.PageInfo, totalCount int) (*types.TransactionConnection, error) {
		var transactionNodes []*types.Transaction
		transactionNodes = append(transactionNodes, nodes...)

		return &types.TransactionConnection{
			Edges:      edges,
			Nodes:      transactionNodes,
			PageInfo:   info,
			TotalCount: Int64(int64(totalCount)),
		}, nil
	}

	fmt.Println(len(transactions.Transactions))
	var transactionRes []*types.Transaction
	for _, p := range transactions.Transactions {
		var transaction types.Transaction
		if err := r.mapper.Hydrate(p, &transaction); err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, nil)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}

		var account types.Account
		if err := r.mapper.Hydrate(p.Account, &account); err != nil {
			err := mainErrors.Format(mainErrors.InternalErr, nil)
			r.logger.Error("debug", zap.Error(err))
			return nil, err
		}
		transaction.Account = &account

		transactionRes = append(transactionRes, &transaction)
	}
	fmt.Println(len(transactions.Transactions), len(transactionRes))

	return connections.TransactionConnectionCon(transactionRes, edger, conn, input)
}

func (r *queryResolver) Acceptance(ctx context.Context, id string) (*types.Acceptance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Acceptances(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.AcceptanceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (types.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOnfidoSDKToken(ctx context.Context) (*types.Response, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	req := &onboardingService.GetOnfidoSDKTokenRequest{
		PersonId: claims.PersonId,
	}
	resp, err := r.onBoardingService.GetOnfidoSDKToken(ctx, req)
	if err != nil {
		r.logger.Error("Get sdk token request failed", zap.Error(err))
		return nil, err
	}
	return &types.Response{
		Message: "successful",
		Success: true,
		Token:   &resp.Token,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const (
	// Error messages
	errorGettingPersonMsg = "failed to get person"
	errorCopierFailedMsg  = "copier failed"
)

func (r *queryResolver) OnfidoReport(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) ComplyAdvReport(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
