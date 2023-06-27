package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config = map[string]any

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

func ReadConfig(filename string) Config {
	doc, err := os.ReadFile(filename)
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

func WriteConfig(filename string, cfg Config) error {

	toml_string, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, toml_string, 0666)

	if err != nil {
		panic(err)
	}

	return err
}

func UpdateConfig(filename string, key string, val any) error {
	cfg := ReadConfig(filename)

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

	WriteConfig(filename, cfg)

	return nil
}
