package repo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

func DownloadIndexFile(name string, url string, cacheDir string) error {
	res, err := http.Get(url + "/index.yaml")
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.Errorf("ERROR: Index not found")
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(cacheDir, name+".yaml"), content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Check the cached list of modules for the specified version and return the URL
func Search(name string, version string, cacheDir string) (string, error) {

	indexFiles, _ := filepath.Glob(cacheDir + "/*.yaml")

	for _, indexFile := range indexFiles {

		data, err := os.ReadFile(indexFile)
		if err != nil {
			fmt.Println(fmt.Errorf("Error reading index file %q: %q", indexFile, err))
		}

		type Index struct {
			ApiVersion string
			Modules    map[string][]map[string]string
		}
		var index Index
		err = yaml.Unmarshal([]byte(data), &index)
		if err != nil {
			fmt.Println(fmt.Errorf("Error reading index file %q: %q", indexFile, err))
		}

		for listedName, versions := range index.Modules {
			if listedName != name {
				continue
			}
			for _, listedVersion := range versions {
				if listedVersion["version"] == version {
					return listedVersion["url"], nil
				}
			}
		}
	}

	return "", fmt.Errorf("Could not find module %q in version %q in any index!", name, version)
}

// DownloadIndexFile fetches the index from a repository.
// func (r *ChartRepository) DownloadIndexFile() error {

// indexURL, err := ResolveReferenceURL(r.Config.URL, "index.yaml")
// if err != nil {
// 	return "", err
// }

// resp, err := r.Client.Get(indexURL,
// 	getter.WithURL(r.Config.URL),
// 	getter.WithInsecureSkipVerifyTLS(r.Config.InsecureSkipTLSverify),
// 	getter.WithTLSClientConfig(r.Config.CertFile, r.Config.KeyFile, r.Config.CAFile),
// 	getter.WithBasicAuth(r.Config.Username, r.Config.Password),
// 	getter.WithPassCredentialsAll(r.Config.PassCredentialsAll),
// )
// if err != nil {
// 	return "", err
// }

// index, err := io.ReadAll(resp)
// if err != nil {
// 	return "", err
// }

// indexFile, err := loadIndex(index, r.Config.URL)
// if err != nil {
// 	return "", err
// }

// // Create the chart list file in the cache directory
// var charts strings.Builder
// for name := range indexFile.Entries {
// 	fmt.Fprintln(&charts, name)
// }
// chartsFile := filepath.Join(r.CachePath, helmpath.CacheChartsFile(r.Config.Name))
// os.MkdirAll(filepath.Dir(chartsFile), 0755)
// os.WriteFile(chartsFile, []byte(charts.String()), 0644)

// // Create the index file in the cache directory
// fname := filepath.Join(r.CachePath, helmpath.CacheIndexFile(r.Config.Name))
// os.MkdirAll(filepath.Dir(fname), 0755)
// return fname, os.WriteFile(fname, index, 0644)
// }
