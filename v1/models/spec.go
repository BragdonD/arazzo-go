package models

import (
	"fmt"
	"regexp"

	"sigs.k8s.io/yaml"
)

// Spec is a struct that represents an [Arazzo 1.0.X] specification.
//
// This is the root object of the Arazzo Description.
//
// [Arazzo 1.0.X]: https://spec.openapis.org/arazzo/v1.0.0.html
type Spec struct {
	// Required. This string MUST be the version number of the
	// Arazzo Specification that the Arazzo Description uses.
	// The arazzo field MUST be used by tooling to interpret
	// the Arazzo Description.
	// Here, the value MUST match the pattern: "^1\.0\.\d(-.+)?$".
	Arazzo string `json:"arazzo"`
	// Required. Provides metadata about the workflows contain
	// within the Arazzo Description. The metadata MAY be used
	// by tooling as required.
	Info Info `json:"info"`
	// Required. A list of source descriptions (such as an OpenAPI
	// description) this Arazzo Description SHALL apply to. The
	// list MUST have at least one entry.
	SourcesDescriptions []SourceDescription `json:"sourceDescriptions"`
	// Required. A list of workflows. The list MUST have at least
	// one entry.
	Workflows []Workflow `json:"workflows"`
	// An element to hold various schemas for the Arazzo Description.
	Components *Components `json:"components,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

const (
	VersionRegex = "^1\\.0\\.\\d(-.+)?$"
)

// ExtractSpecWithDocumentCheck extracts a Spec object from a YAML
// document and checks if the Arazzo version is valid.
func ExtractSpecWithDocumentCheck(doc []byte) (*Spec, error) {
	spec := &Spec{}

	if err := yaml.Unmarshal(doc, spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	versionRe, err := regexp.Compile(VersionRegex)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to compile version regex: %w",
			err,
		)
	}

	if !versionRe.MatchString(spec.Arazzo) {
		return nil, fmt.Errorf(
			"arazzo version must match regex %s",
			VersionRegex,
		)
	}

	return spec, nil
}
