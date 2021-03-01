package mongo

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"log"
	"os"
	"testing"
)

var mongoDbPort = ""

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal(err)
	}

	resource, err := pool.Run("mongo", "4.2.9", []string{
		"MONGO_INITDB_DATABASE=roava",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	mongoDbPort = resource.GetPort("27017/tcp")
	log.Println(mongoDbPort)
	if err := pool.Retry(func() error {
		var err error
		connectUrl := fmt.Sprintf("mongodb://localhost:%s", mongoDbPort)
		log.Println(connectUrl)
		_, _, err = New(connectUrl, "roava", nil)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	code := m.Run()

	err = pool.Purge(resource)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestMongoStore_GetCDDs(t *testing.T) {
	connectUri := "mongodb://localhost:" + mongoDbPort
	repo, client, err := New(connectUri, "roava", zaptest.NewLogger(t))
	assert.Nil(t, err)
	assert.NotNil(t, repo)

	cdd := &models.CDD{
		ID:        "id",
		Owner:     "owner",
		Watchlist: false,
		Validations: []models.Validation{{
			Applicant: models.Person{ID: "personId"},
		}},
	}
	newId, err := client.Database("roava").Collection("cdds").InsertOne(context.Background(), cdd)
	assert.Nil(t, err)
	assert.NotNil(t, newId)

	values, err := repo.GetCDDs(1, 100)
	assert.Nil(t, err)
	assert.NotNil(t, values)
	assert.Equal(t, 1, len(values))
}
