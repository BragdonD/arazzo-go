package v1

import "github.com/bragdonD/arazzo-go/v1/models"

// Step is a struct that represents an Arazzo specification 1.0.X step
// object.
//
// A step describes a single workflow step which MAY be a call to an
// API operation (OpenAPI Operation Object) or another Workflow
// Object.
type Step struct {
	model           *models.Step
	parent          *Workflow
	id              string
	operation       *OAIOperation
	opWorkflow      *Workflow
	parameters      []*Parameter
	requestBody     *RequestBody
	successCriteria []*Criterion
	onSuccess       []*SuccessAction
	onFailure       []*FailureAction
	outputs         map[string]any
}

func NewStep(model *models.Step, parent *Workflow) {
	step := new(Step)
	step.model = model
	step.parent = parent

	//spec := parent.GetParent()

}

func (s *Step) GetModel() *models.Step {
	return s.model
}

func (s *Step) GetParent() *Workflow {
	return s.parent
}
