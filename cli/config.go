package cli

import "fmt"

// Version info
type Version struct {
	Branch string `env:"SOURCE_BRANCH"`
	Commit string `env:"SOURCE_COMMIT"`
	Image  string `env:"IMAGE_NAME"`
}

type Config struct {
	// Env must be local, development, test or production
	Env    string `env:"PROJECTS_ENV" envDefault:"production"`
	Host   string `env:"PROJECTS_HOST"`
	Port   int    `env:"PROJECTS_PORT" envDefault:"3000"`
	APIKey string `env:"PROJECTS_API_KEY"`
	Repo   struct {
		Conn string `env:"CONN,notEmpty"`
	} `envPrefix:"PROJECTS_REPO_"`
	Search struct {
		Conn      string `env:"CONN,notEmpty"`
		Index     string `env:"INDEX,notEmpty"`
		Retention int    `env:"RETENTION" envDefault:"5"`
	} `envPrefix:"PROJECTS_SEARCH_"`
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
