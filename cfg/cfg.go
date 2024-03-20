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
	Mail     Mail     `yaml:"mail"`
}

type Web struct {
	Address   string `yaml:"address"`
	Maintence bool   `yaml:"maintence"`
}

type Database struct {
	DataDir string `yaml:"data_dir"`
}

type Mail struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Mailer   string `yaml:"mailer"`
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
		Str("web address", configuration.Web.Address).
		Bool("maintence", configuration.Web.Maintence).
		Str("database dir", configuration.Database.DataDir).
		Msg("configuration loaded")

	return configuration, nil
}
