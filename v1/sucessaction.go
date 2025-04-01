package v1

import "github.com/bragdonD/arazzo-go/v1/models"

// SuccessAction is a struct that represents an Arazzo specification
// 1.0.X success action object.
//
// A success action object is a single success action which describes
// an action to take upon success of a workflow step.
type SuccessAction struct {
	model      *models.SuccessAction
	name       string
	workflowId *string
	stepId     *string
	criteria   []*Criterion
}
