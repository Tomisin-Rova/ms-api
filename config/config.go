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

type Secrets struct {
	OnfidoServiceURL     string
	KYCServiceURL        string
	OnboardingServiceURL string
	VerifyServiceURL     string
	AuthServiceURL       string
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
	wg                   *sync.WaitGroup
	mu                   *sync.Mutex
}

var _secrets Secrets
var EventRoot = fmt.Sprintf("%s.%s", Domain, ServiceName)

// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets() (*Secrets, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	ss := &Secrets{}
	ss.wg = &sync.WaitGroup{}
	ss.mu = &sync.Mutex{}

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
	ss.OnfidoServiceURL = os.Getenv("ONFIDO_SERVICE")
	ss.KYCServiceURL = os.Getenv("KYC_SERVICE")
	ss.OnboardingServiceURL = os.Getenv("ONBOARDING_SERVICE_URL")
	ss.VerifyServiceURL = os.Getenv("VERIFY_SERVICE_URL")
	ss.AuthServiceURL = os.Getenv("AUTH_SERVICE_URL")
	if err := ss.Environment.IsValid(); err != nil {
		log.Error("Error in environment variables: %v", err)
	}
	return ss, nil
}

// Get Secrets is used to get value from the Secrets runtime.
func GetSecrets() Secrets {
	return _secrets
}

// Watch Secrets does management of hot update on Secrets from vault and any secret store provided.
func WatchSecrets() {
	_secrets.mu.Lock()
	defer _secrets.mu.Unlock()

	data, err := connectVault(_secrets.VaultAddress, _secrets.VaultToken, _secrets.VaultSecretsPath)
	if err != nil {
		logrus.Print("There was an error parsing _secrets from vault: ", err)
		return
	}

	var _s Secrets
	if err := utils.Pack(data, &_s); err != nil {
		logrus.Print("There was an error parsing _secrets from vault: ", err)
		return
	}
	_secrets.PulsarURL = _s.PulsarURL
	_secrets.JWTSecrets = _s.JWTSecrets
	_secrets.wg.Add(1)
	go _secrets.connectRedis()
	//go _secrets.connectEventStore()
	_secrets.wg.Wait()
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
		logrus.Print("There was an error reading _secrets from vault: ", err)
		return nil, err
	}

	return s.Data, nil
}
