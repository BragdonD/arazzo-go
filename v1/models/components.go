package models

// Components is a struct that represents an Arazzo specification
// 1.0.X components object.
//
// A component object holds a set of reusable
// objects for different aspects of the Arazzo Specification. All
// objects defined within the components object will have no effect on
// the Arazzo Description unless they are explicitly referenced from
// properties outside the components object.
//
// Components are scoped to the Arazzo document they are defined in.
// For example, if a step defined in Arazzo document “A”
// references a workflow defined in Arazzo document “B”, the
// components in “A” are not considered when evaluating the
// workflow referenced in “B”.
type Components struct {
	// An object to hold reusable JSON Schema objects to be referenced
	// from workflow inputs.
	Inputs map[string]any `json:"inputs,omitempty"`
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter `json:"parameters,omitempty"`
	// An object to hold reusable Success Action Objects.
	SuccessActions map[string]SuccessAction `json:"successActions,omitempty"`
	// An object to hold reusable Failure Action Objects.
	FailureActions map[string]FailureAction `json:"failureActions,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
