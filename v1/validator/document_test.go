package validator

import (
	"encoding/json"
	"os"
	"testing"

	v1 "github.com/bragdonD/arazzo-go/v1"
	"github.com/stretchr/testify/assert"
)

func TestValidateArazzoDocument(t *testing.T) {
	petstore, err := os.ReadFile("../test_specs/petstore.arazzo.json")
	if err != nil {
		t.Fatalf("failed to read petstore spec: %v", err)
	}

	var spec v1.Spec
	if err := json.Unmarshal([]byte(petstore), &spec); err != nil {
		t.Fatalf("could not unmarshal the test's data: %v", err)
	}

	valid, errs := ValidateArazzoDocument(&spec)

	assert.True(t, valid)
	assert.Empty(t, errs)
}
