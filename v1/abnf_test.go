package v1

import (
	"regexp"
	"testing"
)

func TestABNFNameRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"validName123", "validName123"},
		{"another_valid-name", "another_valid-name"},
		{"name_with_underscore", "name_with_underscore"},
		{"name-with-dash", "name-with-dash"},
		{"name.with.dot", "name.with.dot"},
		{
			"name with space",
			"name",
		}, // Extracts only valid prefix
		{
			"123valid_name!",
			"123valid_name",
		}, // Extracts valid part before `!`
		{
			"!invalidStart",
			"",
		}, // No valid match at the beginning
		{
			"",
			"",
		}, // No match for empty string
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_.-]+`)
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			match := re.FindString(test.input)
			if match != test.expected {
				t.Errorf(
					"For input '%s', expected '%s' but got '%s'",
					test.input,
					test.expected,
					match,
				)
			}
		})
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
			t.Errorf(
				"For input '%s', expected %v but got %v",
				test.input,
				test.expected,
				match,
			)
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
			t.Errorf(
				"For input '%s', expected %v but got %v",
				test.input,
				test.expected,
				match,
			)
		}
	}
}

func TestABNFTokenRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"token", "token"},
		{"token123", "token123"},
		{"!#$%&'*+-.^_`|~", "!#$%&'*+-.^_`|~"},
		{
			"invalid token",
			"invalid",
		}, // Extracts valid part before space
		{
			"123token!",
			"123token!",
		}, // Full match since all chars are valid
		{
			"   leadingSpace",
			"",
		}, // No match because it starts with a space
		{
			"trailingSpace ",
			"trailingSpace",
		}, // Extracts valid part before space
		{
			"",
			"",
		}, // Empty string should not match
	}

	re := regexp.MustCompile(ABNFTokenRegex)

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			match := re.FindString(test.input)
			if match != test.expected {
				t.Errorf(
					"For input '%s', expected '%s' but got '%s'",
					test.input,
					test.expected,
					match,
				)
			}
		})
	}
}
