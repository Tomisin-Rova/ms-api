package sessions

import (
	"errors"
	redisConnector "ms.api/cache/redis"
	"strings"

	"ms.api/config"
)

var (
	InvalidTokenString  = errors.New("sorry, please you must provide a token string")
	SessTokenNotFound   = errors.New("sorry, a session for this token is not found")
	UnitOfValidityError = errors.New("sorry, your unit of validity must be HOUR or MINUTE")
)
var sessionCache = redisConnector.NewCache(config.GetSecrets().RedisClient, 14400, "AUTH_SESSION_")

func GetSessionByToken(token string) (*Session, error) {
	if strings.TrimSpace(token) == "" {
		return nil, InvalidTokenString
	}
	// Verify token before proceed.
	_, err := VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}

	var session Session
	if err := sessionCache.Get(token, &session); err != nil {
		return nil, SessTokenNotFound
	}

	if err := session.AssertValidity(); err != nil {
		// Destroy it in case it still exists in the store,
		// We don't need this error, because it'll never pass the third line at this point.
		_ = DestroySession(session.Token)
		return nil, err
	}

	return &session, nil
}

func CreateSession(payload Session) (*Session, error) {
	//if err := utils.Validator(payload); err != nil {
	//	return nil, err
	//}

	if !payload.UnitOfValidity.IsValid() {
		return nil, UnitOfValidityError
	}
	s, e := NewSession(payload.AccountId, payload.Validity, payload.UnitOfValidity)
	if e != nil {
		return nil, e
	}

	// TODO: Save to redis store with expiration for 4hours and keep extending the expiration for each usage, and not db.
	if err := sessionCache.Set(s.Token, s); err != nil {
		return nil, err
	}

	return s, nil
}

func ExtendSession(token string) error {
	s, err := GetSessionByToken(token)

	if err != nil {
		return err
	}

	_ = s.UpdateLastUsage() // Actual refresh. ( returns time not, error. )
	//TODO: Save extend the token by expiration time returned above.

	return sessionCache.Update(s.Token, s)
}

func DestroySession(token string) error {
	// Delete the session from cache.
	sessionCache.Delete(token)
	return nil
}
