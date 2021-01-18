package config

import (
	"os"

	"ms.api/log"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

const (
	ServiceName = "kyc"
	Domain      = "io.roava"
)

type Secrets struct {
	CddServiceURL        string
	OnfidoServiceURL     string
	KYCServiceURL        string
	OnboardingServiceURL string
	VerifyServiceURL     string
	AuthServiceURL       string
	ProductServiceURL    string
	PayeeServiceURL      string
	PersonServiceURL     string
	PaymentServiceURL    string
	VaultAddress         string        `json:"vault_address"`
	VaultToken           string        `json:"vault_token"`
	VaultSecretsPath     string        `json:"vault_secrets_path"`
	JWTSecrets           string        `json:"jwt_secrets"`
	PulsarURL            string        `json:"pulsar_url"`
	Port                 string        `json:"port"`
	Environment          Environment   `json:"environment"`
	RedisURL             string        `json:"redis_url"`
	RedisPassword        string        `json:"redis_password"`
	RedisClient          *redis.Client `json:"redis_client"`
}

// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets() (*Secrets, error) {
	// Load env vars from .env file
	// TODO: Refactor to use zebra secrets library
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	ss := &Secrets{}
	if ss.Port = os.Getenv("PORT"); ss.Port == "" {
		ss.Port = "20002"
	}
	ss.PulsarURL = os.Getenv("PULSAR_URL")
	ss.RedisURL = os.Getenv("REDIS_URL")
	ss.RedisPassword = os.Getenv("REDIS_PASSWORD")
	ss.JWTSecrets = os.Getenv("JWT_SECRETS")
	ss.Environment = Environment(os.Getenv("ENVIRONMENT"))
	ss.VaultAddress = os.Getenv("VAULT_ADDRESS")
	ss.VaultToken = os.Getenv("VAULT_TOKEN")
	ss.VaultSecretsPath = os.Getenv("VAULT_SECRETS_PATH")
	ss.OnfidoServiceURL = os.Getenv("ONFIDO_SERVICE_URL")
	ss.KYCServiceURL = os.Getenv("KYC_SERVICE_URL")
	ss.OnboardingServiceURL = os.Getenv("ONBOARDING_SERVICE_URL")
	ss.VerifyServiceURL = os.Getenv("VERIFY_SERVICE_URL")
	ss.AuthServiceURL = os.Getenv("AUTH_SERVICE_URL")
	ss.CddServiceURL = os.Getenv("CDD_SERVICE_URL")
	ss.ProductServiceURL = os.Getenv("PRODUCT_SERVICE_URL")
	ss.PersonServiceURL = os.Getenv("PERSON_SERVICE_URL")
	ss.PaymentServiceURL = os.Getenv("PAYMENT_SERVICE_URL")
	ss.PayeeServiceURL = os.Getenv("PAYEE_SERVICE_URL")
	if err := ss.Environment.IsValid(); err != nil {
		log.Error("Error in environment variables: %v", err)
	}
	return ss, nil
}
