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
