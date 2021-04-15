package mongo

import (
	"context"
	"github.com/pkg/errors"
	"github.com/roava/zebra/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"ms.api/libs/db"
	"ms.api/types"
	"time"
)

const (
	cddsCollection    = "cdds"
	checksCollection  = "checks"
	screensCollection = "screens"
	proofsCollection  = "proofs"
	personCollection  = "person"
	identityCollection  = "identities"
	orgsCollection    = "organizations"
	payeeCollection    = "payees"
)

func New(connectURI, databaseName string, logger *zap.Logger) (db.DataStore, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, err
	}
	return &mongoStore{mongoClient: client, databaseName: databaseName, logger: logger}, client, nil
}

type mongoStore struct {
	mongoClient  *mongo.Client
	databaseName string
	logger       *zap.Logger
}

func (s *mongoStore) GetCheck(id string) (*models.Check, error) {
	check := &models.Check{}
	err := s.col(checksCollection).FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(check)
	if err != nil {
		return nil, err
	}
	return check, nil
}

func (s *mongoStore) GetScreen(id string) (*models.Screen, error) {
	screen := &models.Screen{}
	err := s.col(screensCollection).FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(screen)
	if err != nil {
		return nil, err
	}
	return screen, nil
}

func (s *mongoStore) GetProof(id string) (*models.Proof, error) {
	proof := &models.Proof{}
	err := s.col(proofsCollection).FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(proof)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

func (s *mongoStore) GetCDDs(page, perPage int64) ([]*models.CDD, error) {
	opts := &options.FindOptions{}
	opts.SetSkip((page - 1) * perPage)
	opts.SetLimit(perPage)

	cursor, err := s.col(cddsCollection).Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query CDDs")
	}
	cdds := make([]*models.CDD, 0)
	for cursor.Next(context.Background()) {
		cdd := &models.CDD{}
		if err := cursor.Decode(cdd); err != nil {
			return nil, errors.Wrap(err, "failed to decode a single CDD")
		}
		cdds = append(cdds, cdd)
	}
	return cdds, nil
}

func (s *mongoStore) GetPerson(id string) (*models.Person, error) {
	person := &models.Person{}
	err := s.col(personCollection).FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (s *mongoStore) GetOrganization(id string) (*models.Organization, error) {
	org := &models.Organization{}
	err := s.col(orgsCollection).FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}


func (s *mongoStore) GetPayeesByOwner(owner string, opts *types.PayeeAggOpts) ([]*types.PayeeAggregate, error) {
	query := bson.D{
		{Key: "owner", Value: owner},
	}
	pipeline := getPayeePipeline(query, opts)
	cursor, err := s.col(payeeCollection).Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			s.logger.With(zap.Error(err)).Error("failed to close DB read cursor")
		}
	}()
	payees := make([]*types.PayeeAggregate, 0)
	if err := cursor.All(context.Background(), &payees); err != nil {
		return nil, err
	}
	return payees, nil
}

func (s *mongoStore) GetPayee(id, owner string, opts *types.PayeeAggOpts) (*types.PayeeAggregate, error) {
	query := bson.D{
		{Key: "id", Value: id},
		{Key: "owner", Value: owner},
	}
	pipeline := getPayeePipeline(query, opts)
	cursor, err := s.col(payeeCollection).Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			s.logger.With(zap.Error(err)).Error("failed to close DB read cursor")
		}
	}()

	payee := &types.PayeeAggregate{}
	if cursor.Next(context.Background()) {
		err = cursor.Decode(payee)
		if err != nil {
			return nil, err
		}
	}

	return payee, nil
}

func getPayeePipeline(query bson.D, opts *types.PayeeAggOpts) []bson.D {
	match := bson.D{
		{Key: "$match", Value: query},
	}

	// filter only payee accounts that have not been deleted
	filter := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "id", Value: 1},
			{Key: "owner", Value: 1},
			{Key: "name", Value: 1},
			{Key: "avatar", Value: 1},
			{Key: "ts", Value: 1},
			{Key: "accounts", Value: bson.D{
				{Key: "$filter", Value: bson.D{
					{Key: "input", Value: "$accounts"},
					{Key: "as", Value: "account"},
					{Key: "cond", Value: bson.D{{Key: "$eq", Value: bson.A{"$$account.deleted", false}}}},
				}},
			}},
		}},
	}

	lookupIdentitty, unwindIdentitty, projectId := bson.D{}, bson.D{}, bson.D{}
	if opts.Identity {
		lookupIdentitty = bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: identityCollection},
				{Key: "localField", Value: "owner"},
				{Key: "foreignField", Value: "id"},
				{Key: "as", Value: "identity"},
			}},
		}

		unwindIdentitty = bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$identity"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		}

		projectId = bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: "1"},
				{Key: "owner", Value: "$identity"},
				{Key: "name", Value: "1"},
				{Key: "avatar", Value: "1"},
				{Key: "accounts", Value: "1"},
				{Key: "ts", Value: "1"},
			}},
		}

	}

	lookupPerson, unwindPerson, projectPerson := bson.D{}, bson.D{}, bson.D{}
	if opts.Person {
		lookupIdentitty = bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: personCollection},
				{Key: "localField", Value: "owner.owner"},
				{Key: "foreignField", Value: "id"},
				{Key: "as", Value: "person"},
			}},
		}

		unwindPerson = bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$perosn"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		}

		projectPerson = bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: "1"},
				{Key: "owner.owner", Value: "$person"},
				{Key: "name", Value: "1"},
				{Key: "avatar", Value: "1"},
				{Key: "accounts", Value: "1"},
				{Key: "ts", Value: "1"},
			}},
		}
	}

	pipeline := mongo.Pipeline{
		match,
		filter,
		lookupIdentitty,
		unwindIdentitty,
		projectId,
		lookupPerson,
		unwindPerson,
		projectPerson,
	}

	return pipeline
}

func (s *mongoStore) col(collectionName string) *mongo.Collection {
	return s.mongoClient.Database(s.databaseName).Collection(collectionName)
}
