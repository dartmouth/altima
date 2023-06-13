package repo

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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
