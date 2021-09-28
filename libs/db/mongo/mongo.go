package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/roava/zebra/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"ms.api/libs/db"
)

const (
	checksCollection   = "checks"
	screensCollection  = "screen"
	proofsCollection   = "proofs"
	personCollection   = "person"
	orgsCollection     = "organizations"
	identityCollection = "identities"
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
	check := models.Check{}
	ctx := context.Background()
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "id", Value: id}}}}
	lookupReportsStage := bson.D{primitive.E{
		Key: "$lookup",
		Value: bson.D{
			primitive.E{Key: "from", Value: "reports"},
			primitive.E{Key: "localField", Value: "data.reports"},
			primitive.E{Key: "foreignField", Value: "id"},
			primitive.E{Key: "as", Value: "data.reports"},
		},
	}}
	lookupTagsStage := bson.D{primitive.E{
		Key:   "$unset",
		Value: "data.tags",
	}}
	cursor, err := s.col(checksCollection).Aggregate(ctx, mongo.Pipeline{matchStage, lookupReportsStage, lookupTagsStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if cursor.Next(ctx) {
		if err := cursor.Decode(&check); err != nil {
			return nil, errors.Wrap(err, "failed to decode a single Validation")
		}
		return &check, nil
	}

	return nil, errors.New("failed to find check")
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

func (repo *mongoStore) GetIdentityById(identityId string) (*models.Identity, error) {
	identity := &models.Identity{}
	filter := bson.M{"id": identityId}
	err := repo.col(identityCollection).FindOne(context.Background(), filter).Decode(identity)
	if err != nil {
		return nil, err
	}
	return identity, nil
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
