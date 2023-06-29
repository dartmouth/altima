package repo

import (
	"altima/pkg/cli"
	"testing"
)

func TestSearch(t *testing.T) {
	settings := cli.New()

	url, err := Search("cow", "v0.0.1", settings.RepositoryCacheDir)

	if url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	url, err = Search("horse", "v0.0.2", settings.RepositoryCacheDir)

	if url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	url, err = Search("horse", "", settings.RepositoryCacheDir)

	if url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	url, err = Search("horse", "v0.x.y", settings.RepositoryCacheDir)

	if url != "" || err == nil {
		t.Error("Returned a false positive result!")
	}
}

func TestInstallModule(t *testing.T) {
	settings := cli.New()
	err := InstallModule("cow", "https://github.com/crossett/altima-modules/releases/download/v0.0.1/cow-v0.0.1.tgz", settings.ModulesDir)
	if err != nil {
		t.Error("Could not install module!")
	}
}
