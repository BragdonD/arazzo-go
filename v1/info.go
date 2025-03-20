package v1

// Info is a struct that represents an Arazzo specification 1.0.X info
// object.
//
// The info object provides metadata about API workflows defined in
// this Arazzo document. The metadata MAY be used by the clients if
// needed.
type Info struct {
	// Required. A human readable title of the Arazzo Description.
	Title string `json:"title"`
	// A short summary of the Arazzo Description.
	Summary *string `json:"summary,omitempty"`
	// A description of the purpose of the workflows defined.
	// CommonMark syntax MAY be used for rich text representation.
	Description *string `json:"description,omitempty"`
	// Required. The version identifier of the Arazzo document
	// (which is distinct from the Arazzo Specification version).
	Version string `json:"version"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
