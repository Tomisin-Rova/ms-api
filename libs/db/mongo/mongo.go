package mongo

import (
	"context"
	"github.com/roava/zebra/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"ms.api/libs/db"
	"time"
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

func (s *mongoStore) GetCDDs(page, perPage int64) ([]*models.CDD, error) {
	opts := &options.FindOptions{}
	opts.SetSkip((page - 1) * perPage)
	opts.SetLimit(perPage)

	cursor, err := s.col("").Find(context.Background(), nil, opts)
	if err != nil {
		return nil, err
	}
	cdds := make([]*models.CDD, 0)
	for cursor.Next(context.Background()) {
		cdd := &models.CDD{}
		if err := cursor.Decode(cdd); err != nil {
			return nil, err
		}
		cdds = append(cdds, cdd)
	}
	return cdds, nil
}

func (s *mongoStore) col(collectionName string) *mongo.Collection {
	return s.mongoClient.Database(s.databaseName).Collection(collectionName)
}
