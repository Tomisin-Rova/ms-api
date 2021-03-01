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
	"time"
)

const (
	cddsCollection    = "cdds"
	checksCollection  = "checks"
	screensCollection = "screens"
	proofsCollection  = "proofs"
	personCollection  = "person"
	orgsCollection    = "organizations"
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

func (s *mongoStore) col(collectionName string) *mongo.Collection {
	return s.mongoClient.Database(s.databaseName).Collection(collectionName)
}
