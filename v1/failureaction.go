package v1

import "github.com/bragdonD/arazzo-go/v1/models"

// FailureAction is a struct that represents an Arazzo specification
// 1.0.X failure action object.
//
// A failure action object is single failure action which describes an
// action to take upon failure of a workflow step.
type FailureAction struct {
	model      *models.FailureAction
	name       string
	workflowId *string
	stepId     *string
	criteria   []*Criterion
}
