package db

import "github.com/roava/zebra/models"

//go:generate mockgen -source=datastore.go -destination=../../mocks/datastore_mock.go -package=mocks
type DataStore interface {
	GetCDDs(page, perPage int64) ([]*models.CDD, error)
	GetCheck(id string) (*models.Check, error)
	GetScreen(id string) (*models.Screen, error)
	GetProof(id string) (*models.Proof, error)
	GetPerson(id string) (*models.Person, error)
	GetOrganization(id string) (*models.Organization, error)
}
