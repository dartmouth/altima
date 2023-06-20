/*
Copyright © 2023 Simon Stone <simon.stone@dartmouth.edu>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

var configFilePath = "altima.toml"

type Config = map[string]any

// configureCmd represents the enable command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Changes a value in the configuration",
	Long: `This command changes a value in the configuration.

	If the specified key does not exist in the file, an error is raised.
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		val, err := deduceType(args[1])
		if err != nil {
			panic(err)
		}

		err = updateConfig(key, val)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deduceType(v string) (any, error) {
	// The TOML syntax with respect to types is identical to JSON
	// We can therefore leverage the JSON package to decode the strings

	// Is it boolean?
	var b bool
	err := json.Unmarshal([]byte(v), &b)
	if err == nil {
		return b, nil
	}

	// Is it integer?
	var i int
	err = json.Unmarshal([]byte(v), &i)
	if err == nil {
		return i, nil
	}

	// Is it float?
	var f float64
	err = json.Unmarshal([]byte(v), &f)
	if err == nil {
		return f, nil
	}

	// Is it an array?
	var arr []any
	err = json.Unmarshal([]byte(v), &arr)
	if err == nil {
		return arr, nil
	}

	// It probably was a string all along
	return v, nil
}

func readConfig(filename string) Config {
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

func writeConfig(filename string, cfg Config) {

	toml_string, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(configFilePath, toml_string, 0666)

	if err != nil {
		panic(err)
	}
}

func updateConfig(key string, val any) error {
	cfg := readConfig(configFilePath)

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

	writeConfig(configFilePath, cfg)

	return nil
}
