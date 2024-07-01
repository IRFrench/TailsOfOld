package cfg

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
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

func LoadConfig(filePath string) (Configuration, error) {
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
