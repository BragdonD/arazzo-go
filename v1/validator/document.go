package validator

import (
	"encoding/json"
	"strings"

	arazzo "github.com/bragdonD/arazzo-go/v1"
	"github.com/bragdonD/arazzo-go/v1/validator/helpers"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

// ValidateArazzoDocument will validate an Arazzo [Spec] against the
// Arazzo 1.0 schemas (depending on version). It will return true if
// the document is valid, false if it is not and a slice of
// [ValidationError] pointers.
func ValidateArazzoDocument(doc *arazzo.Spec) (bool, []error) {
	compiler := jsonschema.NewCompiler()
	compiler.UseLoader(helpers.NewCompilerLoader())

	jsch, err := compiler.Compile("./schemas/schemav1_0.json")
	if err != nil {
		return false, []error{err}
	}

	loadedSchema, err := json.Marshal(doc)
	if err != nil {
		return false, []error{err}
	}

	decodedSchema, _ := jsonschema.UnmarshalJSON(strings.NewReader(string(loadedSchema)))

	scErrs := jsch.Validate(decodedSchema)

	if scErrs != nil {
		return false, []error{scErrs}
	}
	return true, nil
}
