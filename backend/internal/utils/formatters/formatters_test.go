package formatters

import (
	"testing"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"example", "Example"},
		{"Example", "Example"},
		{"", ""},
		{"a", "A"},
		{"A", "A"},
		{"hello world", "Hello world"},
		{"1hello", "1hello"},
		{" hello", " hello"},
	}

	for _, test := range tests {
		result := CapitalizeFirstLetter(test.input)
		if result != test.expected {
			t.Errorf("capitalizeFirstLetter(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestStringToUint(t *testing.T) {
	tests := []struct {
		input      string
		expected   uint
		isExpected bool
	}{
		{"1", 1, true},
		{"10", 10, true},
		{"0", 0, true},
		{"A", 0, true},
		{"", 0, true},
		{"-1", 1, false},
		{"Two", 2, false},
	}

	for _, test := range tests {
		result := StringToUint(test.input)
		if (result != test.expected) == test.isExpected {
			t.Errorf("StringToUint(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestUintToString(t *testing.T) {
	tests := []struct {
		input      uint
		expected   string
		isExpected bool
	}{
		{1, "1", true},
		{10, "10", true},
		{0, "0", true},
		{999, "Nine hundred ninety nine", false},
	}

	for _, test := range tests {
		result := UintToString(test.input)
		if (result != test.expected) == test.isExpected {
			t.Errorf("UintToString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
