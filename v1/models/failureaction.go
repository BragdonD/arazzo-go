package models

import (
	"encoding/json"
	"errors"
)

// FailureAction is a struct that represents an Arazzo specification
// 1.0.X failure action object.
//
// A failure action object is single failure action which describes an
// action to take upon failure of a workflow step.
type FailureAction struct {
	// Required. The name of the failure action. Names are case
	// sensitive.
	Name string `json:"name"`
	// Required. The type of action to take. Possible values are
	// "end", "retry", or "goto".
	Type FailureActionType `json:"type"`
	// The workflowId referencing an existing workflow within the
	// Arazzo Description to transfer to upon failure of the step.
	// This field is only relevant when the type field value is "goto"
	// or "retry". If the referenced workflow is contained within an
	// arazzo type sourceDescription, then the workflowId MUST be
	// specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<workflowId>) to avoid ambiguity or
	// potential clashes. This field is mutually exclusive to stepId.
	// When used with "retry", context transfers back upon completion
	// of the specified workflow.
	WorkflowId *string `json:"workflowId,omitempty"`
	// The stepId to transfer to upon failure of the step. This field
	// is only relevant when the type field value is "goto" or
	// "retry". The referenced stepId MUST be within the current
	// workflow. This field is mutually exclusive to workflowId. When
	// used with "retry", context transfers back upon completion of
	// the specified step.
	StepId *string `json:"stepId,omitempty"`
	// A non-negative decimal indicating the seconds to delay after
	// the step failure before another attempt SHALL be made. Note: if
	// an HTTP Retry-After response header was returned to a step from
	// a targeted operation, then it SHOULD overrule this particular
	// field value. This field only applies when the type field value
	// is "retry".
	RetryDelay *float64 `json:"retryDelay,omitempty"`
	// A non-negative integer indicating how many attempts to retry
	// the step MAY be attempted before failing the overall step. If
	// not specified then a single retry SHALL be attempted. This
	// field only applies when the type field value is "retry". The
	// retryLimit MUST be exhausted prior to executing subsequent
	// failure actions.
	RetryLimit *int `json:"retryLimit,omitempty"`
	// A list of assertions to determine if this action SHALL be
	// executed. Each assertion is described using a Criterion Object.
	// All criteria assertions MUST be satisfied for the action to be
	// executed
	Criteria []Criterion `json:"criteria,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

// FailureActionType is a string that represents the type of failure
// action available in an Arazzo specification 1.0.X.
type FailureActionType string

const (
	// FailureActionTypeEnd represents an end failure action. The
	// workflow ends, and context returns to the caller with
	// applicable outputs.
	FailureActionTypeEnd FailureActionType = "end"
	// FailureActionTypeRetry represents a retry failure action. The
	// current step will be retried. The retry will be constrained by
	// the retryAfter and retryLimit fields. If a stepId or workflowId
	// are specified, then the reference is executed and the context
	// is returned, after which the current step is retried.
	FailureActionTypeRetry FailureActionType = "retry"
	// FailureActionTypeGoto represents a goto failure action. A
	// one-way transfer of workflow control to the specified label
	// (either a workflowId or stepId).
	FailureActionTypeGoto FailureActionType = "goto"
)

// FailureActionOrReusable allows a step to use either a
// [FailureAction] or a [Reusable] object.
type FailureActionOrReusable struct {
	FailureAction *FailureAction `json:",omitempty"`
	Reusable      *Reusable      `json:",omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (f *FailureActionOrReusable) UnmarshalJSON(data []byte) error {
	var failureAction FailureAction
	if err := json.Unmarshal(data, &failureAction); err == nil {
		f.FailureAction = &failureAction
		return nil
	}

	var reusable Reusable
	if err := json.Unmarshal(data, &reusable); err == nil {
		f.Reusable = &reusable
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (FailureAction, Reusable)",
	)
}

// MarshalJSON implements json.Marshaler interface.
func (f FailureActionOrReusable) MarshalJSON() ([]byte, error) {
	if f.FailureAction != nil {
		return json.Marshal(f.FailureAction)
	}
	if f.Reusable != nil {
		return json.Marshal(f.Reusable)
	}
	return nil, errors.New("no data to marshal")
}
