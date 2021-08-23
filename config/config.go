package config

import (
	"github.com/go-redis/redis/v7"
	"github.com/roava/zebra/secrets"
	"github.com/roava/zebra/secrets/config"
)

const (
	ServiceName = "kyc"
	Domain      = "io.roava"
)

// Ensure interface implementation
var _ config.SecretGroup = &Secrets{}

// Secrets model
type Secrets struct {
	config.DecoratedSecrets `mapstructure:",squash"`
	CddServiceURL           string        `mapstructure:"CDD_SERVICE_URL"`
	OnfidoServiceURL        string        `mapstructure:"ONFIDO_SERVICE_URL"`
	KYCServiceURL           string        `mapstructure:"KYC_SERVICE_URL"`
	OnboardingServiceURL    string        `mapstructure:"ONBOARDING_SERVICE_URL"`
	VerifyServiceURL        string        `mapstructure:"VERIFY_SERVICE_URL"`
	AuthServiceURL          string        `mapstructure:"AUTH_SERVICE_URL"`
	AccountServiceURL       string        `mapstructure:"ACCOUNT_SERVICE_URL"`
	PayeeServiceURL         string        `mapstructure:"PAYEE_SERVICE_URL"`
	PersonServiceURL        string        `mapstructure:"PERSON_SERVICE_URL"`
	PaymentServiceURL       string        `mapstructure:"PAYMENT_SERVICE_URL"`
	IdentityServiceURL      string        `mapstructure:"IDENTITY_SERVICE_URL"`
	PricingServiceURL       string        `mapstructure:"PRICING_SERVICE_URL"`
	JWTSecrets              string        `json:"jwt_secrets" mapstructure:"JWT_SECRETS"`
	RedisURL                string        `json:"redis_url" mapstructure:"REDIS_URL"`
	RedisPassword           string        `json:"redis_password" mapstructure:"REDIS_PASSWORD"`
	RedisClient             *redis.Client `json:"redis_client"`
}

// PostProcess secrets post process logic
func (s Secrets) PostProcess() error { return nil }

// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets() (*Secrets, error) {
	// Load secrets service
	secretsService := secrets.New(".env")
	err := secretsService.LoadFromEnv()
	if err != nil {
		return nil, err
	}
	// Set secret vars
	var _secrets Secrets
	err = secretsService.Unmarshal(&_secrets)
	if err != nil {
		return nil, err
	}
	if _secrets.Service.Port == "" {
		_secrets.Service.Port = "20002"
	}
	if _secrets.Database.Name == "" {
		_secrets.Database.Name = "roava"
	}

	return &_secrets, nil
}
