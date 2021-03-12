// Package models contains the various GraphQL data models
package models

type ConnectionInput struct {
	After  *string
	Before *string
	First  *int64
	Last   *int64
}
