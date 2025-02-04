package util

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func DeduceType(v string) any {
	// The TOML syntax with respect to types is identical to JSON
	// We can therefore leverage the JSON package to decode the strings

	// Is it boolean?
	var b bool
	err := json.Unmarshal([]byte(v), &b)
	if err == nil {
		return b
	}

	// Is it integer?
	var i int
	err = json.Unmarshal([]byte(v), &i)
	if err == nil {
		return i
	}

	// Is it float?
	var f float64
	err = json.Unmarshal([]byte(v), &f)
	if err == nil {
		return f
	}

	// Is it an array?
	var arr []any
	err = json.Unmarshal([]byte(v), &arr)
	if err == nil {
		return arr
	}

	// It probably was a string all along
	return v
}

func DownloadFile(filename string, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.Errorf("ERROR: File not found")
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UnpackFile(filename string, targetRoot string) error {
	dat, err := os.Open(filename)

	defer dat.Close()

	zr, err := gzip.NewReader(dat)
	if err != nil {
		return err
	}

	tr := tar.NewReader(zr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(targetRoot, hdr.Name)

		switch hdr.Typeflag {

		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			// Exclude Mac-specific indicator files
			if path.Base(target)[0:2] == "._" {
				continue
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			f.Close()
		}
	}

	return nil
}

func DeleteFile(filename string) error {
	return os.Remove(filename)
}


func GetName(s string) string {
	// The name of the module comes before the version or the alias
	pattern := regexp.MustCompile("^(.*?)(?:==|>|$)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}

func GetVersion(s string) string {
	// The version always follows `==` and may precede `>`
	pattern := regexp.MustCompile("==(.*?)(?:>|$)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}

func GetAlias(s string) string {
	// The alias always follows `>`
	pattern := regexp.MustCompile(">(.*)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}
