package sessions

import (
	"errors"
	"strings"
	"time"

	"ms.api/utils"
)

func (sm *Session) UpdateLastUsage() time.Time {
	now := time.Now()
	sm.LastUsage = now
	var expiration time.Time
	switch sm.UnitOfValidity {
	case UnitOfValidityHour:
		expiration = now.Add(time.Hour * sm.Validity)
		break
	case UnitOfValidityMinute:
		expiration = now.Add(time.Minute * sm.Validity)
		break
	}
	return expiration
}

func (sm *Session) AssertValidity() error {
	now := time.Now().Unix()
	lastUsage := sm.LastUsage
	uV := sm.UnitOfValidity
	var validity int64 = 0
	switch uV {
	case UnitOfValidityHour:
		validity = lastUsage.Add(time.Hour * sm.Validity).Unix()
		break
	case UnitOfValidityMinute:
		validity = lastUsage.Add(time.Minute * sm.Validity).Unix()
		break
	}

	if now > validity {
		return errors.New("sorry, session has expired. Please login again to continue")
	}
	return nil
}

func NewSession(AccountId string, Validity time.Duration, unitOfValidity UnitOfValidity) (*Session, error) {
	Id, _ := utils.GenerateUUID(24)
	uv := strings.ToUpper(unitOfValidity.String()) // ensuring caps.
	token, err := GenerateToken(AccountId)
	if err != nil {
		return nil, err
	}
	return &Session{
		Id:             Id,
		Token:          token,
		AccountId:      AccountId,
		Validity:       Validity,
		LastUsage:      time.Now(),
		UnitOfValidity: UnitOfValidity(uv),
		TimeCreated:    time.Now(),
	}, nil
}
