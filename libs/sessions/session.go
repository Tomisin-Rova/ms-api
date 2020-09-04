package sessions

import (
	"errors"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"ms.api/config"
)

type UnitOfValidity string

const (
	UnitOfValidityMinute UnitOfValidity = "MINUTE"
	UnitOfValidityHour   UnitOfValidity = "HOUR"
)

func (u UnitOfValidity) IsValid() bool {
	switch u {
	case UnitOfValidityHour, UnitOfValidityMinute:
		return true
	}
	return false
}

func (u UnitOfValidity) String() string {
	return string(u)
}

type AccountTypes string

func (e AccountTypes) String() string {
	return string(e)
}

type Session struct {
	Id             string         `json:"_id" bson:"_id"`
	Token          string         `json:"token" bson:"token"`
	AccountId      string         `json:"accountId" bson:"accountId" validate:"required"`
	TimeCreated    time.Time      `json:"timeCreated" bson:"timeCreated"`
	Validity       time.Duration  `json:"validity" bson:"validity" validate:"required"`
	UnitOfValidity UnitOfValidity `json:"unitOfValidity" bson:"unitOfValidity" validate:"required"`
	LastUsage      time.Time      `json:"lastUsage" bson:"lastUsage" `
}

type TokenPayload struct {
	Id     string `json:"_id"`
	Client string `json:"client"`
	jwt.Payload
}

// Generate Token.
func GenerateToken(id string) (string, error) {
	// Removed time. To handle Session and token invalidation on the server side.
	// Using inactivity token mechanism, db-centric source of truth.
	secret := jwt.NewHS256([]byte(config.GetSecrets().JWTSecrets))
	payload := &TokenPayload{
		Payload: jwt.Payload{
			Issuer:   "Roava.IO",
			Subject:  "Roava API Token",
			Audience: jwt.Audience{"https://roava.io"},
			IssuedAt: jwt.NumericDate(time.Now()),
			JWTID:    "RoavaGunjigalis",
		},
		//Client: client,
		Id:     id}

	token, err := jwt.Sign(payload, secret)
	if err != nil {
		logrus.Println(err.Error(), " Err generating JWT Token")
		return "", err
	}
	return string(token), nil
}

//
func VerifyAuthToken(token string) (*TokenPayload, error) {
	secret := jwt.NewHS256([]byte(config.GetSecrets().JWTSecrets))
	var payloadBody TokenPayload
	_, err := jwt.Verify([]byte(token), secret, &payloadBody)
	if err != nil {
		return nil, errors.New("Sorry, your authentication token is invalid! Please login again to continue. ")
	}
	return &payloadBody, nil
}
