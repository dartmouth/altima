package cli

import (
	"os"
	"path/filepath"
)

type EnvSettings struct {
	ConfigDir            string
	RepositoryConfigFile string
	RepositoryCacheDir   string
	ExecutableDir        string
}

func New() *EnvSettings {
	homedirname, err := os.UserHomeDir()
	check(err)

	executable, err := os.Executable()
	check(err)
	executableDir := filepath.Dir(executable)

	env := &EnvSettings{
		ExecutableDir:        executableDir,
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
