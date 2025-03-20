package models

// Workflow is a struct that represents an Arazzo specification 1.0.X
// workflow object.
//
// A workflow describes the steps to be taken across one or more APIs
// to achieve an objective. The workflow object MAY define inputs
// needed in order to execute workflow steps, where the defined steps
// represent a call to an API operation or another workflow, and a set
// of outputs.
type Workflow struct {
	// Required. Unique string to represent the workflow. The
	// id MUST be unique amongst all workflows described in
	// the Arazzo Description. The workflowId value is
	// case-sensitive. SHOULD conform to the regular
	// expression [A-Za-z0-9_\-]+.
	WorkflowId string `json:"workflowId"`
	// A summary of the purpose or objective of the workflow.
	Summary *string `json:"summary,omitempty"`
	// A description of the workflow. [CommonMark] syntax MAY
	// be used for rich text representation.
	//
	// [CommonMark]: https://spec.commonmark.org/
	Description *string `json:"description,omitempty"`
	// A JSON Schema 2020-12 object representing the input
	// parameters used by this workflow.
	Inputs map[string]any `json:"inputs,omitempty"`
	// A list of workflows that MUST be completed before this
	// workflow can be processed. Each value provided MUST be
	// a workflowId. If the workflow depended on is defined
	// within the current Workflow Document, then specify the
	// workflowId of the relevant local workflow. If the
	// workflow is defined in a separate Arazzo Document then
	// the workflow MUST be defined in the sourceDescriptions
	// and the workflowId MUST be specified using a [Runtime
	// Expression] (e.g., $sourceDescriptions.<name>.<workflowId>)
	// to avoid ambiguity or potential clashes.
	//
	// [Runtime Expression]:
	// https://spec.openapis.org/arazzo/v1.0.0.html#runtime-expressions
	DependsOn []string `json:"dependsOn,omitempty"`
	// Required. An ordered list of steps where each step
	// represents a call to an API operation or to another
	// workflow.
	Steps []Step `json:"steps"`
	// A list of success actions that are applicable for all steps
	// described under this workflow. These success actions can be
	// overridden at the step level but cannot be removed there. If a
	// Reusable Object is provided, it MUST link to success actions
	// defined in the components/successActions of the current Arazzo
	// document. The list MUST NOT include duplicate success actions.
	SuccessActions []SuccessActionOrReusable `json:"successActions,omitempty"`
	// A list of failure actions that are applicable for all steps
	// described under this workflow. These failure actions can be
	// overridden at the step level but cannot be removed there. If a
	// Reusable Object is provided, it MUST link to failure actions
	// defined in the components/failureActions of the current Arazzo
	// document. The list MUST NOT include duplicate failure actions.
	FailureActions []FailureActionOrReusable `json:"failureActions,omitempty"`
	// A map between a friendly name and a dynamic output value. The
	// name MUST use keys that match the regular expression:
	// ^[a-zA-Z0-9\.\-_]+$.
	Outputs map[string]any `json:"outputs,omitempty"`
	// A list of parameters that are applicable for all steps
	// described under this workflow. These parameters can be
	// overridden at the step level but cannot be removed there. Each
	// parameter MUST be passed to an operation or workflow as
	// referenced by operationId, operationPath, or workflowId as
	// specified within each step. If a Reusable Object is provided,
	// it MUST link to a parameter defined in the
	// components/parameters of the current Arazzo document. The list
	// MUST NOT include duplicate parameters.
	Parameters []ParameterOrReusable `json:"parameters,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
