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

func NewComponents(model *models.Components) *Components {
	components := &Components{
		model:          model,
		parameters:     map[string]*Parameter{},
		successActions: map[string]*SuccessAction{},
		failureActions: map[string]*FailureAction{},
	}

	for _, param := range model.Parameters {
		parameter := NewParameter(&param)
		components.parameters[param.Name] = parameter
	}

	// TODO: Add success actions
	// TODO: Add failure actions
	return components
}

func (c *Components) GetModel() *models.Components {
	return c.model
}

func (c *Components) GetParameters() map[string]*Parameter {
	return c.parameters
}

func (c *Components) GetParameter(name string) (*Parameter, bool) {
	param, ok := c.parameters[name]
	return param, ok
}
