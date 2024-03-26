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
	BulkMail Mail     `yaml:"bulk_mail"`
}

type Web struct {
	Address    string `yaml:"address"`
	Maintence  bool   `yaml:"maintence"`
	Newsletter bool   `yaml:"newsletter"`
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
		Str("address", configuration.Web.Address).
		Bool("maintence mode", configuration.Web.Maintence).
		Bool("newsletter", configuration.Web.Newsletter).
		Msg("web configuration")

	log.Info().
		Str("directory", configuration.Database.DataDir).
		Msg("database configuration")

	log.Info().
		Str("username", configuration.Mail.Username).
		Str("host", configuration.Mail.Host).
		Str("mailer", configuration.Mail.Mailer).
		Msg("mail configuration")

	log.Info().
		Str("username", configuration.BulkMail.Username).
		Str("host", configuration.BulkMail.Host).
		Str("mailer", configuration.BulkMail.Mailer).
		Msg("bulk mail configuration")

	return configuration, nil
}
