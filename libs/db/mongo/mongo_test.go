package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
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

func TestIdentityRepository_GetIdentityById(t *testing.T) {
	connectUri := fmt.Sprintf("mongodb://localhost:%s", mongoDbPort)
	repo, client, err := New(connectUri, "roava", zaptest.NewLogger(t))
	assert.Nil(t, err)
	assert.NotNil(t, client)

	identity := &models.Identity{
		ID:        "identityId",
		Owner:     "owner",
		Timestamp: time.Now(),
	}
	r, err := client.Database("roava").Collection(identityCollection).
		InsertOne(context.Background(), identity)
	assert.Nil(t, err)
	assert.NotNil(t, r.InsertedID)

	id, err := repo.GetIdentityById(identity.ID)
	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, id.ID, identity.ID)
}
