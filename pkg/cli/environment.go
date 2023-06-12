package cli

import (
	"os"
	"path/filepath"
)

type EnvSettings struct {
	CacheDir           string
	ConfigDir          string
	ConfigFilename     string
	ExecutableDir      string
	ModuleDir          string
	RepositoryCacheDir string
}

func New() *EnvSettings {
	homedirname, err := os.UserHomeDir()
	check(err)

	executable, err := os.Executable()
	check(err)
	executableDir := filepath.Dir(executable)

	env := &EnvSettings{
		CacheDir:           envOr("ALTIMA_CACHE_DIR", filepath.Join(homedirname, ".config", "altima", "cache")),
		ConfigDir:          envOr("ALTIMA_CONFIG_DIR", filepath.Join(homedirname, ".config", "altima")),
		ConfigFilename:     envOr("ALTIMA_CONFIG_FILENAME", "altima.toml"),
		ExecutableDir:      executableDir,
		ModuleDir:          envOr("ALTIMA_MODULE_DIR", filepath.Join(homedirname, ".config", "altima", "module")),
		RepositoryCacheDir: envOr("ALTIMA_REPOSITORY_CACHE_DIR", filepath.Join(homedirname, ".config", "altima", "cache", "repository")),
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
