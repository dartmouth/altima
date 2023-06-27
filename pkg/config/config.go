package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config = map[string]any

var Filepath = "altima.toml"

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

func ReadConfig() Config {
	return readConfig(Filepath)
}

func readConfig(filepath string) Config {
	doc, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	cfg := make(Config)
	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func WriteConfig(cfg Config) error {

	return writeConfig(Filepath, cfg)
}

func writeConfig(filepath string, cfg Config) error {
	toml_string, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath, toml_string, 0666)

	if err != nil {
		panic(err)
	}

	return err
}

func UpdateConfig(key string, val any) error {
	return updateConfig(Filepath, key, val)
}

func updateConfig(filepath string, key string, val any) error {
	cfg := readConfig(filepath)

	// The key may be a nested key (dot notation)
	parts := strings.Split(key, ".")
	// Get a pointer to the outermost map
	current_map := &cfg
	// Now follow the sequence of keys until the last one (exclusive)
	for _, part := range parts[:len(parts)-1] {
		// The value of the current map must be another, inner map
		v, ok := (*current_map)[part].(Config)
		if !ok {
			return errors.New("Could not follow nested keys!")
		}
		current_map = &v
	}
	// The current map is the innermost map
	// Check if the key exists first
	_, ok := (*current_map)[parts[len(parts)-1]]
	if !ok {
		return errors.New("Could not find key!")
	}
	// If it exists, set it to the requested value
	(*current_map)[parts[len(parts)-1]] = val

	writeConfig(filepath, cfg)

	return nil
}

func Enable(module string) error {
	return enable(Filepath, module)
}

func enable(filepath string, module string) error {
	err := updateConfig(filepath, "modules."+module+".enabled", true)

	if err != nil {
		err = fmt.Errorf("Module %q not found!", module)
	}
	return err
}
