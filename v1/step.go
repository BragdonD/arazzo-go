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

func NewStep(model *models.Step, parent *Workflow) (*Step, error) {
	step := &Step{
		model:           model,
		parent:          parent,
		id:              model.StepId,
		parameters:      []*Parameter{},
		requestBody:     nil,
		successCriteria: []*Criterion{},
		onSuccess:       []*SuccessAction{},
		onFailure:       []*FailureAction{},
		outputs:         map[string]any{},
	}

	// Step's parameters can come from three sources:
	//  - the step itself, in this case the parameter MUST NOT
	//    be duplicated.
	//  - its parent workflow, in this case the parameter CAN
	//    be duplicated by the step and it will take precedence.
	//  - the components object, in this case the parameter is
	//    store in the components object and the step stores a
	//    reference to it. This parameter SHOULD NOT be duplicated
	//    in the step.
	for _, paramOrReusable := range model.Parameters {
		var param *models.Parameter
		var err error
		if paramOrReusable.Parameter != nil {
			param = paramOrReusable.Parameter
		} else if paramOrReusable.Reusable != nil {
			// TODO: Change this to take a reference from the Components
			// object store in the workflow's parent spec.
			param, err = paramOrReusable.Reusable.ToParameter(parent.GetParent().GetComponents().GetModel())
		}
		if err != nil {
			return nil, err
		}
		if param != nil {
			parameter := NewParameter(param)
			step.parameters = append(step.parameters, parameter)
		}
	}

	return step, nil
}

// checkParameters verifies that parameters are not duplicated
func (step *Step) checkParameters() error {

}

func (s *Step) GetModel() *models.Step {
	return s.model
}

func (s *Step) GetParent() *Workflow {
	return s.parent
}
