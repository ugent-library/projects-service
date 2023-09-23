package cli

import "fmt"

type Config struct {
	// Version info
	Version struct {
		Branch string `env:"SOURCE_BRANCH"`
		Commit string `env:"SOURCE_COMMIT"`
		Image  string `env:"IMAGE_NAME"`
	}
	// Env must be local, development, test or production
	Env    string `env:"PROJECTS_ENV" envDefault:"production"`
	Host   string `env:"PROJECTS_HOST"`
	Port   int    `env:"PROJECTS_PORT" envDefault:"3000"`
	APIKey string `env:"PROJECTS_API_KEY"`
	Nats   struct {
		URL      string `env:"URL,notEmpty"`
		Stream   string `env:"STREAM,notEmpty"`
		Consumer string `env:"CONSUMER,notEmpty"`
	} `envPrefix:"PROJECTS_NATS_"`
	Repo struct {
		Conn   string `env:"CONN,notEmpty"`
		Secret string `env:"SECRET,notEmpty"`
	} `envPrefix:"PROJECTS_REPO_"`
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
