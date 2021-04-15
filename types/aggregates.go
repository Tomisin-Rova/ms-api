package types

import "github.com/roava/zebra/models"

type PayeeAggregate struct {
	models.Payee
	Owner  IdentityAggregate `json:"owner" bson:"owner"`
}

type IdentityAggregate struct {
	models.Identity
	Owner  models.Person `json:"owner" bson:"owner"`
}

type PayeeAggOpts struct {
	Person bool
	Identity bool
}
