package config

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"ms.api/log"
	"ms.api/utils"
	"os"
	"sync"
)

const (
	ServiceName = "kyc"
	Domain      = "io.roava"
)

var _ = godotenv.Load()

type Secrets struct {
	OnfidoServiceURL    string
	KYCServiceURL    string
	OnboardingServiceURL    string
	VaultAddress     string        `json:"vault_address"`
	VaultToken       string        `json:"vault_token"`
	VaultSecretsPath string        `json:"vault_secrets_path"`
	JWTSecrets       string        `json:"jwt_secrets"`
	PulsarURL        string        `json:"pulsar_url"`
	Port             string        `json:"port"`
	Environment      Environment   `json:"environment"`
	RedisURL         string        `json:"redis_url"`
	RedisPassword    string        `json:"redis_password"`
	RedisClient      *redis.Client `json:"redis_client"`
	wg               *sync.WaitGroup
	mu               *sync.Mutex
}

var secrets Secrets
var EventRoot = fmt.Sprintf("%s.%s", Domain, ServiceName)

/**
This loads up Secrets from the .env file once.
If an env file is present, Secrets will be loaded, else it'll be ignored.
*/
func init() {
	secrets.wg = &sync.WaitGroup{}
	secrets.mu = &sync.Mutex{}

	if secrets.Port = os.Getenv("PORT"); secrets.Port == "" {
		secrets.Port = "20002"
	}
	secrets.PulsarURL = os.Getenv("PULSAR_URL")
	secrets.RedisURL = os.Getenv("REDIS_URL")
	secrets.RedisPassword = os.Getenv("REDIS_PASSWORD")
	secrets.JWTSecrets = os.Getenv("JWT_SECRETS")
	secrets.Environment = Environment(os.Getenv("ENVIRONMENT"))
	secrets.VaultAddress = os.Getenv("VAULT_ADDRESS")
	secrets.VaultToken = os.Getenv("VAULT_TOKEN")
	secrets.VaultSecretsPath = os.Getenv("VAULT_SECRETS_PATH")
	secrets.OnfidoServiceURL = os.Getenv("ONFIDO_SERVICE")
	secrets.KYCServiceURL = os.Getenv("KYC_SERVICE")
	secrets.OnboardingServiceURL = os.Getenv("ONBOARDING_SERVICE")
	if err := secrets.Environment.IsValid(); err != nil {
		log.Error("Error in environment variables: %v", err)
	}
}

// Get Secrets is used to get value from the Secrets runtime.
func GetSecrets() Secrets {
	return secrets
}

// Watch Secrets does management of hot update on Secrets from vault and any secret store provided.
func WatchSecrets() {
	secrets.mu.Lock()
	defer secrets.mu.Unlock()

	data, err := connectVault(secrets.VaultAddress, secrets.VaultToken, secrets.VaultSecretsPath)
	if err != nil {
		logrus.Print("There was an error parsing secrets from vault: ", err)
		return
	}

	var _s Secrets
	if err := utils.Pack(data, &_s); err != nil {
		logrus.Print("There was an error parsing secrets from vault: ", err)
		return
	}
	secrets.PulsarURL = _s.PulsarURL
	secrets.JWTSecrets = _s.JWTSecrets
	secrets.wg.Add(1)
	go secrets.connectRedis()
	//go secrets.connectEventStore()
	secrets.wg.Wait()
}

func (s *Secrets) connectRedis() {
	defer s.wg.Done()
	s.RedisClient = redis.NewClient(&redis.Options{
		Addr: s.RedisURL,
		//DB:       0,
		Password: s.RedisPassword,
	})

	if s.RedisClient == nil {
		logrus.Fatal("Redis client invalid.")
	}
}

func connectVault(address, token, path string) (utils.JSON, error) {
	config := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		logrus.Print("There was an error connecting to vault: ", err)
		return nil, err
	}
	client.SetToken(token)
	s, err := client.Logical().Read(path)
	if err != nil {
		logrus.Print("There was an error reading secrets from vault: ", err)
		return nil, err
	}

	return s.Data, nil
}
