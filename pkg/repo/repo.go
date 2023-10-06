package repo

import (
	"altima/pkg/util"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Module struct {
	Name    string
	Version string
	Alias   string
	Repo    string
	Url     string
}

func DownloadIndexFile(name string, url string, cacheDir string) error {
	err := util.DownloadFile(filepath.Join(cacheDir, name+".yaml"), url+"/index.yaml")
	if err != nil {
		return err
	}

	return nil
}

// Check the cached list of modules for the specified version and return the Module info
// If `version` is an empty string, the version listed last in the index is used
func Search(module Module, cacheDir string) (Module, error) {
	indexFiles, _ := filepath.Glob(cacheDir + "/*.yaml")

	for _, indexFile := range indexFiles {

		data, err := os.ReadFile(indexFile)
		if err != nil {
			fmt.Println(fmt.Errorf("Error reading index file %q: %q", indexFile, err))
			continue
		}

		repoName := filepath.Base(indexFile)[:len(filepath.Base(indexFile))-5]

		// Read index File
		type Index struct {
			ApiVersion string
			Modules    map[string][]map[string]string
		}
		var index Index
		err = yaml.Unmarshal([]byte(data), &index)
		if err != nil {
			fmt.Println(fmt.Errorf("Error reading index file %q: %q", indexFile, err))
			continue
		}

		// Find URL for correct version of module
		for listedName, versions := range index.Modules {
			if listedName != module.Name {
				continue
			}
			// If the version was not explicitly specified, return the url of the last version
			if module.Version == "" {
				module.Version = versions[len(versions)-1]["version"]
				module.Url = versions[len(versions)-1]["url"]
				module.Repo = repoName
				return module, nil
			}
			// Otherwise look for an exact match
			for _, listedVersion := range versions {
				if listedVersion["version"] == module.Version {
					module.Url = listedVersion["url"]
					module.Repo = repoName
					return module, nil
				}
			}
		}
	}

	return module, fmt.Errorf("Could not find module %q in version %q in any index!", module.Name, module.Version)
}

func InstallModule(module Module, rootDir string) error {
	archiveFile := filepath.Join(rootDir, path.Base(module.Url))
	installName := module.Name
	if module.Alias != "" {
		installName = module.Alias
	}

	moduleRootFolder := filepath.Join(rootDir, installName)
	err := util.DownloadFile(archiveFile, module.Url)
	if err != nil {
		return err
	}

	err = util.UnpackFile(archiveFile, moduleRootFolder)
	if err != nil {
		return err
	}

	err = util.DeleteFile(archiveFile)
	if err != nil {
		return err
	}

	return nil
}

func UninstallModule(module Module, rootDir string) error {
	installName := module.Name
	if module.Alias != "" {
		installName = module.Alias
	}

	moduleRootFolder := filepath.Join(rootDir, installName)
	err := os.RemoveAll(moduleRootFolder)

	return err
}