package v1

// Reusable is a struct that represents an Arazzo specification 1.0.X
// reusable object.
//
// A reusable object is simple object to allow referencing of objects
// contained within the Components Object. It can be used from
// locations within steps or workflows in the Arazzo description.
type Reusable struct {
	// Required. A Runtime Expression used to reference the desired
	// object.
	Reference string `json:"reference"`
	// Sets a value of the referenced parameter. This is only
	// applicable for parameter object references.
	Value string `json:"value,omitempty"`
}
