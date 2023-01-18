package controller

import (
	"time"

	"github.com/michaelquigley/cf"
	"github.com/openziti/zrok/controller/store"
	"github.com/pkg/errors"
)

const ConfigVersion = 1

type Config struct {
	V            int
	Account      *AccountConfig
	Admin        *AdminConfig
	Endpoint     *EndpointConfig
	Email        *EmailConfig
	Influx       *InfluxConfig
	Limits       *LimitsConfig
	Maintenance  *MaintenanceConfig
	Metrics      *MetricsConfig
	Registration *RegistrationConfig
	Store        *store.Config
	Ziti         *ZitiConfig
}

type AccountConfig struct {
	ForgotPasswordUrlTemplate string
}

type AdminConfig struct {
	Secrets []string `cf:"+secret"`
}

type EndpointConfig struct {
	Host string
	Port int
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string `cf:"+secret"`
}

type RegistrationConfig struct {
	EmailFrom               string
	RegistrationUrlTemplate string
	TokenStrategy           string
}

type ZitiConfig struct {
	ApiEndpoint string
	Username    string
	Password    string `cf:"+secret"`
}

type MetricsConfig struct {
	ServiceName string
}

type InfluxConfig struct {
	Url    string
	Bucket string
	Org    string
	Token  string `cf:"+secret"`
}

type MaintenanceConfig struct {
	Registration *RegistrationMaintenanceConfig
}

type RegistrationMaintenanceConfig struct {
	ExpirationTimeout time.Duration
	CheckFrequency    time.Duration
	BatchLimit        int
}

const Unlimited = -1

type LimitsConfig struct {
	Environments int
	Shares       int
}

func DefaultConfig() *Config {
	return &Config{
		Limits: &LimitsConfig{
			Environments: Unlimited,
			Shares:       Unlimited,
		},
		Metrics: &MetricsConfig{
			ServiceName: "metrics",
		},
		Maintenance: &MaintenanceConfig{
			Registration: &RegistrationMaintenanceConfig{
				ExpirationTimeout: time.Hour * 24,
				CheckFrequency:    time.Hour,
				BatchLimit:        500,
			},
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()
	if err := cf.BindYaml(cfg, path, cf.DefaultOptions()); err != nil {
		return nil, errors.Wrapf(err, "error loading controller config '%v'", path)
	}
	if cfg.V != ConfigVersion {
		return nil, errors.Errorf("expecting configuration version '%v', your configuration is version '%v'; please see zrok.io for changelog and configuration documentation", ConfigVersion, cfg.V)
	}
	return cfg, nil
}
