package config

import (
	"strings"

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
	DatabaseURL             string        `mapstructure:"mongodb_uri"`
	PulsarURL               string        `mapstructure:"pulsar_url"`
	PulsarCert              string        `mapstructure:"pulsar_cert"`
	CddServiceURL           string        `mapstructure:"cdd_service_url"`
	OnfidoServiceURL        string        `mapstructure:"onfido_service_url"`
	KYCServiceURL           string        `mapstructure:"kyc_service_url"`
	OnboardingServiceURL    string        `mapstructure:"onboarding_service_url"`
	VerifyServiceURL        string        `mapstructure:"verify_service_url"`
	AuthServiceURL          string        `mapstructure:"auth_service_url"`
	AccountServiceURL       string        `mapstructure:"account_service_url"`
	PayeeServiceURL         string        `mapstructure:"payee_service_url"`
	PersonServiceURL        string        `mapstructure:"person_service_url"`
	PaymentServiceURL       string        `mapstructure:"payment_service_url"`
	IdentityServiceURL      string        `mapstructure:"identity_service_url"`
	PricingServiceURL       string        `mapstructure:"pricing_service_url"`
	JWTSecrets              string        `json:"jwt_secrets" mapstructure:"JWT_SECRETS"`
	RedisURL                string        `json:"redis_url" mapstructure:"redis_url"`
	RedisPassword           string        `json:"redis_password" mapstructure:"redis_password"`
	RedisClient             *redis.Client `json:"redis_client"`
}

// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets() (*Secrets, error) {
	cfg := &Secrets{}
	// Load secrets
	err := secrets.UnmarshalMergedConfig(cfg, ".env", secrets.AllVaultSubPaths()...)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// PostProcess ...
func (s *Secrets) PostProcess() error {
	s.Database.URL = strings.Trim(s.Database.URL, "\n")
	if s.Service.Port == "" {
		s.Service.Port = "20002"
	}
	if s.Database.Name == "" {
		s.Database.Name = "roava"
	}
	if len(s.DatabaseURL) > 0 {
		s.Database.URL = s.DatabaseURL
	}
	if len(s.PulsarURL) > 0 {
		s.Pulsar.URL = s.PulsarURL
	}
	if len(s.PulsarCert) > 0 {
		s.Pulsar.TLSCert = s.PulsarCert
	}

	return nil
}
