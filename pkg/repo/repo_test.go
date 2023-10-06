package repo

import (
	"altima/pkg/cli"
	"testing"
)

func TestSearch(t *testing.T) {
	settings := cli.New()

	module, err := Search(Module{Name: "cow", Version: "v0.0.1"}, settings.RepositoryCacheDir)

	if module.Url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	module, err = Search(Module{Name: "horse", Version: "v0.0.2"}, settings.RepositoryCacheDir)

	if module.Url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	module, err = Search(Module{Name: "horse"}, settings.RepositoryCacheDir)

	if module.Url == "" || err != nil {
		t.Error("Failed to find module in index!")
	}

	module, err = Search(Module{Name: "horse", Version: "v0.x.y"}, settings.RepositoryCacheDir)

	if module.Url != "" || err == nil {
		t.Error("Returned a false positive result!")
	}
}

func TestInstallModule(t *testing.T) {
	settings := cli.New()
	testModuleWithAlias := Module{Name: "cow", Alias: "mycow", Url: "https://github.com/crossett/altima-modules/releases/download/v0.0.1/cow-v0.0.1.tgz"}
	err := InstallModule(testModuleWithAlias, settings.ModulesDir)
	if err != nil {
		t.Error("Could not install module!")
	}
	testModule := Module{Name: "cow", Url: "https://github.com/crossett/altima-modules/releases/download/v0.0.1/cow-v0.0.1.tgz"}
	err = InstallModule(testModule, settings.ModulesDir)
	if err != nil {
		t.Error("Could not install module!")
	}
}

func TestUninstallModule(t *testing.T){
	settings := cli.New()
	testModuleWithAlias := Module{Name: "cow", Alias: "mycow", Url: "https://github.com/crossett/altima-modules/releases/download/v0.0.1/cow-v0.0.1.tgz"}
	err := InstallModule(testModuleWithAlias, settings.ModulesDir)
	if err != nil {
		t.Error("Could not install module!")
	}
	err = UninstallModule(testModuleWithAlias, settings.ModulesDir)
	if err != nil {
		t.Error("Could not uninstall module!")
	}

	testModule := Module{Name: "cow", Url: "https://github.com/crossett/altima-modules/releases/download/v0.0.1/cow-v0.0.1.tgz"}
	err = InstallModule(testModule, settings.ModulesDir)
	if err != nil {
		t.Error("Could not install module!")
	}

	err = UninstallModule(testModule, settings.ModulesDir)
	if err != nil {
		t.Error("Could not uninstall module!")
	}
}