package db

import "github.com/roava/zebra/models"

type DataStore interface {
	GetCDDs(page, perPage int64) ([]*models.CDD, error)
}
