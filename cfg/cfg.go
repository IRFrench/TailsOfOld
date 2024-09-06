package cfg

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const (
	dbEnvironment          = "DB"
	webEnvironment         = "WEB"
	maintenanceEnvironment = "MAINTENANCE"
)

var (
	ErrFailedToReadConfigFile = errors.New("failed to read configuration file")
	ErrUnableToParseConfig    = errors.New("failed to parse configuration file")
)

type Configuration struct {
	Web      Web      `yaml:"web"`
	Database Database `yaml:"database"`
}

type Web struct {
	Address   string `yaml:"address"`
	Maintence bool   `yaml:"maintence"`
}

type Database struct {
	DataDir string `yaml:"data_dir"`
}

func LoadConfigFromFile(filePath string) (Configuration, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return Configuration{}, ErrFailedToReadConfigFile
	}

	configuration := Configuration{}

	if err = yaml.Unmarshal(configFile, &configuration); err != nil {
		return configuration, ErrUnableToParseConfig
	}

	log.Info().
		Str("address", configuration.Web.Address).
		Bool("maintence mode", configuration.Web.Maintence).
		Msg("web configuration")

	log.Info().
		Str("directory", configuration.Database.DataDir).
		Msg("database configuration")

	return configuration, nil
}

func LoadConfigFromEnvironment() (Configuration, error) {
	newConfig := Configuration{}
	var ok bool

	newConfig.Web.Address, ok = os.LookupEnv(webEnvironment)
	if !ok {
		return Configuration{}, fmt.Errorf("missing environment: %v", webEnvironment)
	}

	newConfig.Database.DataDir, ok = os.LookupEnv(dbEnvironment)
	if !ok {
		return Configuration{}, fmt.Errorf("missing environment: %v", dbEnvironment)
	}

	_, newConfig.Web.Maintence = os.LookupEnv(maintenanceEnvironment)

	return newConfig, nil
}
