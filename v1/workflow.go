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

func (w *Workflow) GetParent() *Spec {
	return w.parent
}
