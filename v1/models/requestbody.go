package models

// RequestBody is a struct that represents an Arazzo specification
// 1.0.X request body object.
//
// A request body object describes the Content-Type and content to be
// passed by a step to an operation.
type RequestBody struct {
	// The Content-Type for the request content. If omitted then refer
	// to Content-Type specified at the targeted operation to
	// understand serialization requirements.
	ContentType *string `json:"contentType,omitempty"`
	// A value representing the request body payload. The value can be
	// a literal value or can contain Runtime Expressions which MUST
	// be evaluated prior to calling the referenced operation. To
	// represent examples of media types that cannot be naturally
	// represented in JSON or YAML, use a string value to contain the
	// example, escaping where necessary.
	Payload any `json:"payload,omitempty"`
	// A list of locations and values to set within a payload.
	Replacements []PayloadReplacement `json:"replacements,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
