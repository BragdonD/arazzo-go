package v1

import "github.com/bragdonD/arazzo-go/v1/models"

// Workflow is a struct that represents an Arazzo specification 1.0.X
// workflow object.
type Workflow struct {
	model          *models.Workflow
	parent         *Spec
	id             string
	dependencies   []*Workflow
	steps          []*Step
	successActions []*SuccessAction
	failureActions []*FailureAction
	outputs        map[string]any
	parameters     []*Parameter
}

func NewWorkflow(model *models.Workflow, parent *Spec) (*Workflow, error) {
	workflow := &Workflow{
		model:          model,
		parent:         parent,
		id:             model.WorkflowId,
		dependencies:   []*Workflow{},
		steps:          []*Step{},
		successActions: []*SuccessAction{},
		failureActions: []*FailureAction{},
		outputs:        map[string]any{},
		parameters:     []*Parameter{},
	}

	for _, paramOrReusable := range model.Parameters {
		var param *models.Parameter
		var err error
		if paramOrReusable.Parameter != nil {
			param = paramOrReusable.Parameter
		} else if paramOrReusable.Reusable != nil {
			param, ok := parent.GetComponents().GetParameter(paramOrReusable.Reusable.Reference)
		}
		if err != nil {
			return nil, err
		}
		if param != nil {
			parameter := NewParameter(param)
			workflow.parameters = append(workflow.parameters, parameter)
		}
	}

	for _, step := range model.Steps {
		stepObj, err := NewStep(&step, workflow)
		if err != nil {
			return nil, err
		}
		workflow.steps = append(workflow.steps, stepObj)
	}

	return workflow, nil
}

func (w *Workflow) GetParent() *Spec {
	return w.parent
}

func (w *Workflow) ResolveDependencies() error {
	return nil
}
