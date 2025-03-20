package models

// PayloadReplacement is a struct that represents an Arazzo
// specification 1.0.X payload replacement object.
//
// A payload replacement object describes a location within a payload
// (e.g., a request body) and a value to set within the location.
type PayloadReplacement struct {
	// Required. A [JSON Pointer] or [XPath Expression] which MUST be
	// resolved against the request body. Used to identify the
	// location to inject the value.
	//
	// [JSON Pointer]: https://tools.ietf.org/html/rfc6901
	// [XPath Expression]:
	// https://www.w3.org/TR/xpath-31/#id-expressions
	Target string `json:"target"`
	// Required. The value set within the target location. The value
	// can be a constant or a Runtime Expression to be evaluated and
	// passed to the referenced operation or workflow.
	Value any `json:"value"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
