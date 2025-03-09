package v1

import (
	"regexp"
	"testing"
)

func TestABNFNameRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"validName123", true},
		{"invalid name", false},
		{"another_valid-name", true},
		{"invalid@name", false},
	}

	re := regexp.MustCompile(ABNFNameRegex)
	for _, test := range tests {
		match := re.MatchString(test.input)
		if match != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, match)
		}
	}
}

func TestABNFEscapedRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"~0", true},
		{"~1", true},
		{"~2", false},
		{"~", false},
	}
	re := regexp.MustCompile(ABNFEscapedRegex)
	for _, test := range tests {
		match := re.MatchString(test.input)
		if match != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, match)
		}
	}
}

func TestABNFUnescapedRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"/", false},
		{"~", false},
		{"\x7E", false},
		{"\x2E", true},
		{"\x30", true},
	}

	re := regexp.MustCompile(ABNFUnescapedRegex)
	for _, test := range tests {
		match := re.MatchString(test.input)
		if match != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, match)
		}
	}
}

func TestABNFTokenRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"token", true},
		{"token123", true},
		{"!#$%&'*+-.^_`|~", true},
		{"invalid token", false},
	}

	re := regexp.MustCompile(ABNFTokenRegex)

	for _, test := range tests {
		match := re.MatchString(test.input)
		if match != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, match)
		}
	}
}
