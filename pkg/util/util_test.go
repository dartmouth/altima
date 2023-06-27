package util

import (
	"testing"
)

func TestDeduceType(t *testing.T) {

	// Boolean
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

	// Integer
	expect_pass = []string{"0", "23", "42"}
	expect_fail = []string{"0.", "2.3", "1e6"}

	for _, s := range expect_pass {
		shouldBeInt := DeduceType(s)
		_, ok := shouldBeInt.(int)
		if !ok {
			t.Errorf("Could not deduce integer from '%s'!", s)
		}
	}

	for _, s := range expect_fail {
		shouldNotBeInt := DeduceType(s)
		_, ok := shouldNotBeInt.(int)
		if ok {
			t.Errorf("Incorrectly deduced integer from '%s'!", s)
		}
	}

	// Float
	expect_pass = []string{"0.0", "2.3", "1e6"}
	expect_fail = []string{"0.", "23", "42"}

	for _, s := range expect_pass {
		shouldBeInt := DeduceType(s)
		_, ok := shouldBeInt.(float64)
		if !ok {
			t.Errorf("Could not deduce float from '%s'!", s)
		}
	}

	for _, s := range expect_fail {
		shouldNotBeInt := DeduceType(s)
		_, ok := shouldNotBeInt.(float64)
		if ok {
			t.Errorf("Incorrectly deduced float from '%s'!", s)
		}
	}

	// Array
	expect_pass = []string{"[ [ 1, 2 ], [3, 4, 5] ]", "[ 1, 2, 3 ]", "[ 0.1, 0.2, 0.5, 1, 2, 5 ]"}
	expect_fail = []string{"1, 2, 3", "(1, 2, 3)"}

	for _, s := range expect_pass {
		shouldBeArray := DeduceType(s)
		_, ok := shouldBeArray.([]any)
		if !ok {
			t.Errorf("Could not deduce array from '%s'!", s)
		}
	}

	for _, s := range expect_fail {
		shouldNotBeArray := DeduceType(s)
		_, ok := shouldNotBeArray.([]any)
		if ok {
			t.Errorf("Incorrectly deduced array from '%s'!", s)
		}
	}

}
