package config

import (
	"strings"

	"github.com/roava/zebra/secrets"
	"github.com/roava/zebra/secrets/config"
)

const (
	Domain = "io.roava"

	ServiceName      = "api"
	LocalEnvironment = "local"
	DevEnvironment   = "dev"
)

// Ensure interface implementation
var _ config.SecretGroup = &Secrets{}

// Secrets model
type Secrets struct {
	config.DecoratedSecrets `mapstructure:",squash"`
	DatabaseURL             string `mapstructure:"mongodb_uri"`
	PulsarURL               string `mapstructure:"pulsar_url"`
	PulsarCert              string `mapstructure:"pulsar_cert"`
	OnboardingServiceURL    string `mapstructure:"onboarding_service_url"`
	VerificationServiceURL  string `mapstructure:"verification_service_url"`
	AuthServiceURL          string `mapstructure:"auth_service_url"`
	AccountServiceURL       string `mapstructure:"account_service_url"`
	CustomerServiceURL      string `mapstructure:"customer_service_url"`
	PaymentServiceURL       string `mapstructure:"payment_service_url"`
	PricingServiceURL       string `mapstructure:"pricing_service_url"`
}

// LoadSecrets loads up Secrets from the vault server.
// If environment it's local Secrets are loaded from local.yml file.
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

	port := s.Service.Port
	if port == "" {
		port = "8080"
	}
	if s.Database.Name == "" {
		s.Database.Name = "roava"
	}
	s.Service.Port = port
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
