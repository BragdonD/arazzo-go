package v1

import (
	"encoding/json"
	"errors"
)

// SuccessAction is a struct that represents an Arazzo specification
// 1.0.X success action object.
//
// A success action object is a single success action which describes
// an action to take upon success of a workflow step.
type SuccessAction struct {
	// Required. The name of the success action. Names are case
	// sensitive.
	Name string `json:"name"`
	// Required. The type of action to take. Possible values are "end"
	// or "goto".
	Type SuccessActionType `json:"type"`
	// The workflowId referencing an existing workflow within the
	// Arazzo Description to transfer to upon success of the step.
	// This field is only relevant when the type field value is
	// "goto". If the referenced workflow is contained within an
	// arazzo type sourceDescription, then the workflowId MUST be
	// specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<workflowId>) to avoid ambiguity or
	// potential clashes. This field is mutually exclusive to stepId.
	WorkflowId *string `json:"workflowId,omitempty"`
	// The stepId to transfer to upon success of the step. This field
	// is only relevant when the type field value is "goto". The
	// referenced stepId MUST be within the current workflow. This
	// field is mutually exclusive to workflowId.
	StepId *string `json:"stepId,omitempty"`
	// A list of assertions to determine if this action SHALL be
	// executed. Each assertion is described using a Criterion Object.
	// All criteria assertions MUST be satisfied for the action to be
	// executed.
	Criteria []Criterion `json:"criteria,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

// SuccessActionType is a string that represents the type of success
// action available in an Arazzo specification 1.0.X.
type SuccessActionType string

const (
	// SuccessActionTypeEnd represents an end success action. The
	// workflow ends, and context returns to the caller with
	// applicable outputs.
	SuccessActionTypeEnd SuccessActionType = "end"
	// SuccessActionTypeGoto represents a goto success action. A
	// one-way transfer of workflow control to the specified label
	// (either a workflowId or stepId)
	SuccessActionTypeGoto SuccessActionType = "goto"
)

func (s SuccessActionType) ToPtr() *SuccessActionType {
	return &s
}

// SuccessActionOrReusable allows a step to use either a
// [SuccessAction] or a [Reusable] object.
type SuccessActionOrReusable struct {
	SuccessAction *SuccessAction `json:",omitempty"`
	Reusable      *Reusable      `json:",omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *SuccessActionOrReusable) UnmarshalJSON(data []byte) error {
	var successAction SuccessAction
	if err := json.Unmarshal(data, &successAction); err == nil {
		s.SuccessAction = &successAction
		return nil
	}

	var reusable Reusable
	if err := json.Unmarshal(data, &reusable); err == nil {
		s.Reusable = &reusable
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (SuccessAction, Reusable)",
	)
}

// MarshalJSON implements json.Marshaler interface.
func (s SuccessActionOrReusable) MarshalJSON() ([]byte, error) {
	if s.SuccessAction != nil {
		return json.Marshal(s.SuccessAction)
	}
	if s.Reusable != nil {
		return json.Marshal(s.Reusable)
	}
	return nil, errors.New("no data to marshal")
}
