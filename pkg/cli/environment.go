package cli

import (
	"log"
	"os"
	"path/filepath"
)

type EnvSettings struct {
	ConfigDir            string
	RepositoryConfigFile string
	RepositoryCacheDir   string
}

func New() *EnvSettings {
	homedirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	env := &EnvSettings{
		ConfigDir:            envOr("ALTIMA_CONFIG_DIR", filepath.Join(homedirname, ".config", "altima")),
		RepositoryConfigFile: envOr("ALTIMA_REPOSITORY_CONFIG_FILE", "repositories.toml"),
		RepositoryCacheDir:   envOr("ALTIMA_REPOSITORY_CACHE_DIR", filepath.Join(homedirname, ".config", "altima", "cache", "repository")),
	}
	return env
}

func envOr(name, def string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	return def
}
