package config

import (
	"testing"
)

var test_cfg = Config{
	"key":    "value",
	"number": 4,
	"table": map[string]any{
		"name":    "test",
		"enabled": false,
		"subtable": map[string]any{
			"inner_name": "test2",
		},
	},
}

func TestDeduceType(t *testing.T) {

	expect_pass := []string{"true", "false"}
	expect_fail := []string{"True", "T", "maybe", "1", "False", "F", "0"}

	for _, s := range expect_pass {
		shouldBeBoolean := DeduceType(s)
		_, ok := shouldBeBoolean.(bool)
		if !ok {
			t.Errorf("Could not deduce boolean from '%s'!", s)
		}
	}

	for _, s := range expect_fail {
		shouldNotBeBoolean := DeduceType(s)
		_, ok := shouldNotBeBoolean.(bool)
		if ok {
			t.Errorf("Incorrectly deduced boolean from '%s'!", s)
		}
	}
}

func TestReadConfig(t *testing.T) {

	cfg := ReadConfig("test_config.toml")

	// Test basic key/value pair
	if cfg["key"] != "value" {
		t.Errorf("Could not validate key 'key' and value 'value'!")
	}

	// Test table
	if len(cfg["modules"].(map[string]any)) != 2 {
		t.Errorf("Incorrect number of modules found!")
	}
	if cfg["user"].(map[string]any)["name"] != "Jack Doe" {
		t.Errorf("Incorrect 'name' read in table 'user'!")
	}

	// Test nested table
	if cfg["modules"].(map[string]any)["mycow"].(map[string]any)["name"] != "cow" {
		t.Errorf("Could not validate name 'cow' for module 'mycow'!")
	}

}

func TestWriteConfig(t *testing.T) {
	err := WriteConfig("test_config_out.toml", test_cfg)
	if err != nil {
		t.Errorf("Could not write config file!")
	}
}

func TestUpdateConfig(t *testing.T) {
	err := UpdateConfig("test_config_out.toml", "table.subtable.inner_name", "new_name")
	if err != nil {
		t.Errorf("Could not update config file!")
	}
}
