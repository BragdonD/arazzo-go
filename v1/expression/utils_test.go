package expression

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"$workflows.foo.inputs.username", false},
		{"$workflows.foo.inputs.username#123", true},
	}

	for _, tt := range tests {
		_, err := Parse(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf(
				"Parse(%q) error = %v, wantErr %v",
				tt.input,
				err,
				tt.wantErr,
			)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"$workflows.foo.inputs.username", true},
		{"$workflows.foo.inputs.username#123", false},
	}

	for _, tt := range tests {
		if got := Validate(tt.input); got != tt.want {
			t.Errorf(
				"Validate(%q) = %v, want %v",
				tt.input,
				got,
				tt.want,
			)
		}
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr bool
	}{
		{
			"{$workflows.foo.inputs.username}",
			"$workflows.foo.inputs.username",
			false,
		},
		{"{content} extra", "", true},
		{"content}", "", true},
		{"{content", "", true},
	}

	for _, tt := range tests {
		got, err := Extract(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf(
				"Extract(%q) error = %v, wantErr %v",
				tt.input,
				err,
				tt.wantErr,
			)
			continue
		}
		if got != tt.want {
			t.Errorf(
				"Extract(%q) = %v, want %v",
				tt.input,
				got,
				tt.want,
			)
		}
	}
}
