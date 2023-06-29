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
}
