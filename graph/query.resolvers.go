package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

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
	person, err := getPerson(personDto)
	if err != nil {
		return nil, err
	}

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
	person.Cdd = r.hydrateCDD(cddDto)

	return person, nil
}

func (r *queryResolver) Person(ctx context.Context, id string) (*types.Person, error) {
	person, err := r.personService.Person(ctx, &personService.PersonRequest{Id: id})
	if err != nil {
		return nil, err
	}

	p, err := getPerson(person)
	if err != nil {
		return nil, err
	}

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
	p.Cdd = r.hydrateCDD(cddDto)

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
	var chunkProcessErr error

	req := &cddService.CDDSRequest{
		Page:    1,
		PerPage: perPage,
		Status:  dataConverter.StateToStringSlice(status),
	}
	resp, err := r.cddService.CDDS(context.Background(), req)
	if err != nil {
		r.logger.With(zap.Error(err)).Error("failed to fetch cdds")
		return nil, terror.NewTerror(7013, "InternalError", "failed to load CDDs data. Internal system error", "internal system error")
	}

	cdds := resp.Results

	dataResolver := NewDataResolver(r.dataStore, r.logger)
	cddsValues := make([]*types.Cdd, len(cdds))
	// Wait until all cdds are inserted into the cddsValues slice
	var wg sync.WaitGroup
	maxGouroutinesCount := len(cdds)
	cddsChan := make(chan *CddChunk)
	errorsChan := make(chan error, maxGouroutinesCount+1)
	if maxGouroutinesCount > 0 {
		size := len(cdds) / maxGouroutinesCount
		for i := 0; i < maxGouroutinesCount; i++ {
			var rest int
			// Process the rest of the people in the slice
			if i == maxGouroutinesCount-1 {
				rest = len(cdds) % maxGouroutinesCount
			}
			cddsChunk := cdds[i*size : (i+1)*size+rest]
			wg.Add(1)
			go func(pos int, cddsChan chan *CddChunk, errorsChan chan error) {
				defer wg.Done()
				data, err := r.processCddChunk(ctx, cddsChunk, dataResolver, dataConverter)
				if err != nil {
					errorsChan <- err
					return
				}
				chunk := &CddChunk{
					pos:  pos * size,
					cdds: data,
				}
				cddsChan <- chunk
			}(i, cddsChan, errorsChan)
		}

		quit := make(chan int)
		go func(targetData []*types.Cdd, quit chan int) {
			for {
				select {
				case chunk := <-cddsChan:
					for i := range chunk.cdds {
						targetData[i+chunk.pos] = chunk.cdds[i]
					}
				case errValue := <-errorsChan:
					chunkProcessErr = errValue
				case <-quit:
					return
				}
			}
		}(cddsValues, quit)
		// Halt cddsChan consumer goroutine
		defer func() {
			quit <- 1
		}()
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}
	if chunkProcessErr != nil {
		r.logger.With(zap.Error(chunkProcessErr)).Error("failed to fetch some cdds")
	}

	//Remove null values
	cddsResult := make([]*types.Cdd, 0)
	for i := range cddsValues {
		if cddsValues[i] != nil {
			cddsResult = append(cddsResult, cddsValues[i])
		}
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
	validation := r.validation(validationDto)

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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.ProductConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context, id string) (*types.Account, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}
	account, err := r.accountService.GetAccount(ctx, &accountService.GetAccountRequest{
		Id:         id,
		IdentityId: claims.IdentityId,
	})
	if err != nil {
		r.logger.Error("failed to get account", zap.Error(err))
		return nil, err
	}
	p := r.hydrateAccount(account)
	return p, nil
}

func (r *queryResolver) Accounts(ctx context.Context, first *int64, after *string, last *int64, before *string) (*types.AccountConnection, error) {
	claims, err := middlewares.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
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

	var accountRes []*types.Account
	for _, c := range accounts.Accounts {
		p := r.hydrateAccount(c)
		accountRes = append(accountRes, p)
	}

	return connections.AccountConnectionCon(accountRes, edger, conn, input)
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
