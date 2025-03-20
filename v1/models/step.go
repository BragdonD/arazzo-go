package models

// Step is a struct that represents an Arazzo specification 1.0.X step
// object.
//
// A step describes a single workflow step which MAY be a call to an
// API operation (OpenAPI Operation Object) or another Workflow
// Object.
type Step struct {
	// A description of the step. [CommonMark] syntax MAY be used
	// for rich text representation.
	//
	// [CommonMark]: https://spec.commonmark.org/
	Description *string `json:"description,omitempty"`
	// Required. Unique string to represent the step. The stepId
	// MUST be unique amongst all steps described in the workflow.
	// The stepId value is case-sensitive. SHOULD conform to the
	// regular expression [A-Za-z0-9_\-]+.
	StepId string `json:"stepId"`
	// The name of an existing, resolvable operation, as defined
	// with a unique operationId and existing within one of the
	// sourceDescriptions. The referenced operation will be
	// invoked by this workflow step. If multiple (non arazzo
	// type) sourceDescriptions are defined, then the operationId
	// MUST be specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<operationId>) to avoid
	// ambiguity or potential clashes. This field is mutually
	// exclusive of the operationPath and workflowId fields
	// respectively.
	OperationId *string `json:"operationId,omitempty"`
	// A reference to a Source Description Object combined with a JSON
	// Pointer to reference an operation. This field is mutually
	// exclusive of the operationId and workflowId fields
	// respectively. The operation being referenced MUST be described
	// within one of the sourceDescriptions descriptions. A Runtime
	// Expression syntax MUST be used to identify the source
	// description document. If the referenced operation has an
	// operationId defined then the operationId SHOULD be preferred
	// over the operationPath.
	OperationPath *string `json:"operationPath,omitempty"`
	// The workflowId referencing an existing workflow within the
	// Arazzo Description. If the referenced workflow is contained
	// within an arazzo type sourceDescription, then the workflowId
	// MUST be specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<workflowId>) to avoid ambiguity or
	// potential clashes. The field is mutually exclusive of the
	// operationId and operationPath fields respectively.
	WorkflowId *string `json:"workflowId,omitempty"`
	// A list of parameters that MUST be passed to an operation or
	// workflow as referenced by operationId, operationPath, or
	// workflowId. If a parameter is already defined at the Workflow,
	// the new definition will override it but can never remove it. If
	// a Reusable Object is provided, it MUST link to a parameter
	// defined in the components/parameters of the current Arazzo
	// document. The list MUST NOT include duplicate parameters.
	Parameters []ParameterOrReusable `json:"parameters,omitempty"`
	// The request body to pass to an operation as referenced by
	// operationId or operationPath. The requestBody is fully
	// supported in HTTP methods where the HTTP 1.1 specification
	// [RFC9110] [Section 9.3] explicitly defines semantics for
	// “content” like request bodies, such as within POST, PUT,
	// and PATCH methods. For methods where the HTTP specification
	// provides less clarity—such as GET, HEAD, and DELETE—the use
	// of requestBody is permitted but does not have well-defined
	// semantics. In these cases, its use SHOULD be avoided if
	// possible.
	//
	// [RFC9110]: https://tools.ietf.org/html/rfc9110
	// [Section 9.3]: https://tools.ietf.org/html/rfc9110#section-9.3
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	// A list of assertions to determine the success of the step. Each
	// assertion is described using a Criterion Object. All assertions
	// MUST be satisfied for the step to be deemed successful.
	SuccessCriteria []Criterion `json:"successCriteria,omitempty"`
	// An array of success action objects that specify what to do upon
	// step success. If omitted, the next sequential step shall be
	// executed as the default behavior. If multiple success actions
	// have similar criteria, the first sequential action matching the
	// criteria SHALL be the action executed. If a success action is
	// already defined at the Workflow, the new definition will
	// override it but can never remove it. If a Reusable Object is
	// provided, it MUST link to a success action defined in the
	// components of the current Arazzo document. The list MUST NOT
	// include duplicate success actions.
	OnSuccess []SuccessActionOrReusable `json:"onSuccess,omitempty"`
	// An array of failure action objects that specify what to do upon
	// step failure. If omitted, the default behavior is to break and
	// return. If multiple failure actions have similar criteria, the
	// first sequential action matching the criteria SHALL be the
	// action executed. If a failure action is already defined at the
	// Workflow, the new definition will override it but can never
	// remove it. If a Reusable Object is provided, it MUST link to a
	// failure action defined in the components of the current Arazzo
	// document. The list MUST NOT include duplicate failure actions
	OnFailure []FailureActionOrReusable `json:"onFailure,omitempty"`
	// A map between a friendly name and a dynamic output value
	// defined using a Runtime Expression. The name MUST use keys that
	// match the regular expression: ^[a-zA-Z0-9\.\-_]+$.
	Outputs map[string]any `json:"outputs,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}
