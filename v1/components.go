package v1

import "github.com/bragdonD/arazzo-go/v1/models"

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
	model *models.Components
	// TODO: ADD a field for the input schema using jsonschema
	parameters     map[string]*Parameter
	successActions map[string]*SuccessAction
	failureActions map[string]*FailureAction
}
