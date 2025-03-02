package v1

import (
	"encoding/json"
	"errors"

	"gopkg.in/yaml.v3"
)

// Spec is a struct that represents an [Arazzo 1.0.X] specification.
//
// This is the root object of the Arazzo Description.
//
// [Arazzo 1.0.X]: https://spec.openapis.org/arazzo/v1.0.0.html
type Spec struct {
	// Required. This string MUST be the version number of the
	// Arazzo Specification that the Arazzo Description uses.
	// The arazzo field MUST be used by tooling to interpret
	// the Arazzo Description.
	// Here, the value MUST match the pattern: "^1\.0\.\d(-.+)?$".
	Arazzo string `json:"arazzo"               yaml:"arazzo"`
	// Required. Provides metadata about the workflows contain
	// within the Arazzo Description. The metadata MAY be used
	// by tooling as required.
	Info Info `json:"info"                 yaml:"info"`
	// Required. A list of source descriptions (such as an OpenAPI
	// description) this Arazzo Description SHALL apply to. The
	// list MUST have at least one entry.
	SourcesDescriptions []SourceDescription `json:"sourceDescriptions"   yaml:"sourceDescriptions"`
	// Required. A list of workflows. The list MUST have at least
	// one entry.
	Workflows []Workflow `json:"workflows"            yaml:"workflows"`
	// An element to hold various schemas for the Arazzo Description.
	Components *Components `json:"components,omitempty" yaml:"components,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"          yaml:"-,omitempty"`
}

// Info is a struct that represents an Arazzo specification 1.0.X info
// object.
//
// The info object provides metadata about API workflows defined in
// this Arazzo document. The metadata MAY be used by the clients if
// needed.
type Info struct {
	// Required. A human readable title of the Arazzo Description.
	Title string `json:"title"                 yaml:"title"`
	// A short summary of the Arazzo Description.
	Summary *string `json:"summary,omitempty"     yaml:"summary,omitempty"`
	// A description of the purpose of the workflows defined.
	// CommonMark syntax MAY be used for rich text representation.
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required. The version identifier of the Arazzo document
	// (which is distinct from the Arazzo Specification version).
	Version string `json:"version"               yaml:"version"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"           yaml:"-,omitempty"`
}

// SourceDescription is a struct that represents an Arazzo
// specification 1.0.X source description object.
//
// A source description object describes a source description (such as
// an OpenAPI description) that will be referenced by one or more
// workflows described within an Arazzo Description.
//
// An object storing a map between named description keys and location
// URLs to the source descriptions (such as an OpenAPI description)
// this Arazzo Description SHALL apply to. Each source location string
// MUST be in the form of a URI-reference as defined by [RFC3986]
// [Section 4.1].
//
// [RFC3986]: https://tools.ietf.org/html/rfc3986
// [Section 4.1]: https://tools.ietf.org/html/rfc3986#section-4.1
type SourceDescription struct {
	// Required. A unique name for the source description.
	// SHOULD conform to the regular expression [A-Za-z0-9_\-]+.
	Name string `json:"name"           yaml:"name"`
	// Required. A URL to a source description to be used by a
	// workflow. If a relative reference is used, it MUST be in
	// the form of a URI-reference as defined by [RFC3986]
	// [Section 4.2].
	//
	// [RFC3986]: https://tools.ietf.org/html/rfc3986
	// [Section 4.2]: https://tools.ietf.org/html/rfc3986#section-4.2
	Url string `json:"url"            yaml:"url"`
	// The type of source description.
	Type *SourceDescriptionType `json:"type,omitempty" yaml:"type,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"    yaml:"-,omitempty"`
}

// SourceDescriptionType is a string that represents the type of
// source description available in an Arazzo specification 1.0.X.
type SourceDescriptionType string

const (
	// SourceDescriptionTypeOpenAPI represents an OpenAPI
	// documentation as source description.
	SourceDescriptionTypeOpenAPI SourceDescriptionType = "openapi"
	// SourceDescriptionTypeArazzo represents an Arazzo
	// documentation as source description.
	SourceDescriptionTypeArazzo SourceDescriptionType = "arazzo"
)

func (s SourceDescriptionType) ToPtr() *SourceDescriptionType {
	return &s
}

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
	WorkflowId string `json:"workflowId"               yaml:"workflowId"`
	// A summary of the purpose or objective of the workflow.
	Summary *string `json:"summary,omitempty"        yaml:"summary,omitempty"`
	// A description of the workflow. [CommonMark] syntax MAY
	// be used for rich text representation.
	//
	// [CommonMark]: https://spec.commonmark.org/
	Description *string `json:"description,omitempty"    yaml:"description,omitempty"`
	// A JSON Schema 2020-12 object representing the input
	// parameters used by this workflow.
	Inputs map[string]any `json:"inputs,omitempty"         yaml:"inputs,omitempty"`
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
	DependsOn []string `json:"dependsOn,omitempty"      yaml:"dependsOn,omitempty"`
	// Required. An ordered list of steps where each step
	// represents a call to an API operation or to another
	// workflow.
	Steps []Step `json:"steps"                    yaml:"steps"`
	// A list of success actions that are applicable for all steps
	// described under this workflow. These success actions can be
	// overridden at the step level but cannot be removed there. If a
	// Reusable Object is provided, it MUST link to success actions
	// defined in the components/successActions of the current Arazzo
	// document. The list MUST NOT include duplicate success actions.
	SuccessActions []SuccessActionOrReusable `json:"successActions,omitempty" yaml:"successActions,omitempty"`
	// A list of failure actions that are applicable for all steps
	// described under this workflow. These failure actions can be
	// overridden at the step level but cannot be removed there. If a
	// Reusable Object is provided, it MUST link to failure actions
	// defined in the components/failureActions of the current Arazzo
	// document. The list MUST NOT include duplicate failure actions.
	FailureActions []FailureActionOrReusable `json:"failureActions,omitempty" yaml:"failureActions,omitempty"`
	// A map between a friendly name and a dynamic output value. The
	// name MUST use keys that match the regular expression:
	// ^[a-zA-Z0-9\.\-_]+$.
	Outputs map[string]any `json:"outputs,omitempty"        yaml:"outputs,omitempty"`
	// A list of parameters that are applicable for all steps
	// described under this workflow. These parameters can be
	// overridden at the step level but cannot be removed there. Each
	// parameter MUST be passed to an operation or workflow as
	// referenced by operationId, operationPath, or workflowId as
	// specified within each step. If a Reusable Object is provided,
	// it MUST link to a parameter defined in the
	// components/parameters of the current Arazzo document. The list
	// MUST NOT include duplicate parameters.
	Parameters []ParameterOrReusable `json:"parameters,omitempty"     yaml:"parameters,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"              yaml:"-,omitempty"`
}

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
	Description *string `json:"description,omitempty"     yaml:"description,omitempty"`
	// Required. Unique string to represent the step. The stepId
	// MUST be unique amongst all steps described in the workflow.
	// The stepId value is case-sensitive. SHOULD conform to the
	// regular expression [A-Za-z0-9_\-]+.
	StepId string `json:"stepId"                    yaml:"stepId"`
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
	OperationId *string `json:"operationId,omitempty"     yaml:"operationId,omitempty"`
	// A reference to a Source Description Object combined with a JSON
	// Pointer to reference an operation. This field is mutually
	// exclusive of the operationId and workflowId fields
	// respectively. The operation being referenced MUST be described
	// within one of the sourceDescriptions descriptions. A Runtime
	// Expression syntax MUST be used to identify the source
	// description document. If the referenced operation has an
	// operationId defined then the operationId SHOULD be preferred
	// over the operationPath.
	OperationPath *string `json:"operationPath,omitempty"   yaml:"operationPath,omitempty"`
	// The workflowId referencing an existing workflow within the
	// Arazzo Description. If the referenced workflow is contained
	// within an arazzo type sourceDescription, then the workflowId
	// MUST be specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<workflowId>) to avoid ambiguity or
	// potential clashes. The field is mutually exclusive of the
	// operationId and operationPath fields respectively.
	WorkflowId *string `json:"workflowId,omitempty"      yaml:"workflowId,omitempty"`
	// A list of parameters that MUST be passed to an operation or
	// workflow as referenced by operationId, operationPath, or
	// workflowId. If a parameter is already defined at the Workflow,
	// the new definition will override it but can never remove it. If
	// a Reusable Object is provided, it MUST link to a parameter
	// defined in the components/parameters of the current Arazzo
	// document. The list MUST NOT include duplicate parameters.
	Parameters []ParameterOrReusable `json:"parameters,omitempty"      yaml:"parameters,omitempty"`
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
	RequestBody *RequestBody `json:"requestBody,omitempty"     yaml:"requestBody,omitempty"`
	// A list of assertions to determine the success of the step. Each
	// assertion is described using a Criterion Object. All assertions
	// MUST be satisfied for the step to be deemed successful.
	SuccessCriteria []Criterion `json:"successCriteria,omitempty" yaml:"successCriteria,omitempty"`
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
	OnSuccess []SuccessActionOrReusable `json:"onSuccess,omitempty"       yaml:"onSuccess,omitempty"`
	// An array of failure action objects that specify what to do upon
	// step failure. If omitted, the default behavior is to break and
	// return. If multiple failure actions have similar criteria, the
	// first sequential action matching the criteria SHALL be the
	// action executed. If a failure action is already defined at the
	// Workflow, the new definition will override it but can never
	// remove it. If a Reusable Object is provided, it MUST link to a
	// failure action defined in the components of the current Arazzo
	// document. The list MUST NOT include duplicate failure actions
	OnFailure []FailureActionOrReusable `json:"onFailure,omitempty"       yaml:"onFailure,omitempty"`
	// A map between a friendly name and a dynamic output value
	// defined using a Runtime Expression. The name MUST use keys that
	// match the regular expression: ^[a-zA-Z0-9\.\-_]+$.
	Outputs map[string]any `json:"outputs,omitempty"         yaml:"outputs,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"               yaml:"-,omitempty"`
}

// ParameterOrReusable allows a step to use either a [Parameter] or
// a [Reusable] object.
type ParameterOrReusable struct {
	Parameter *Parameter `json:",omitempty" yaml:",omitempty"`
	Reusable  *Reusable  `json:",omitempty" yaml:",omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (p *ParameterOrReusable) UnmarshalJSON(data []byte) error {
	var param Parameter
	if err := json.Unmarshal(data, &param); err == nil {
		p.Parameter = &param
		return nil
	}

	var reusable Reusable
	if err := json.Unmarshal(data, &reusable); err == nil {
		p.Reusable = &reusable
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (Parameter, Reusable)",
	)
}

// MarshalJSON implements json.Marshaler interface.
func (p ParameterOrReusable) MarshalJSON() ([]byte, error) {
	if p.Parameter != nil {
		return json.Marshal(p.Parameter)
	}
	if p.Reusable != nil {
		return json.Marshal(p.Reusable)
	}
	return nil, errors.New("no data to marshal")
}

// UnmarshalYAML implements yaml.Unmarshaler interface.
func (p *ParameterOrReusable) UnmarshalYAML(value *yaml.Node) error {
	key := value.Content[0].Value

	if key == "reference" {
		var reusable Reusable
		if err := value.Decode(&reusable); err == nil {
			p.Reusable = &reusable
			return nil
		}
	}

	var param Parameter
	if err := value.Decode(&param); err == nil {
		p.Parameter = &param
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (Parameter, Reusable)",
	)
}

// MarshalYAML implements yaml.Marshaler interface.
func (p ParameterOrReusable) MarshalYAML() (interface{}, error) {
	if p.Parameter != nil {
		return p.Parameter, nil
	}
	if p.Reusable != nil {
		return p.Reusable, nil
	}
	return nil, errors.New("no data to marshal")
}

// Parameter is a struct that represents an Arazzo specification 1.0.X
// parameter object.
//
// A paraemter describes a single step parameter. A unique parameter
// is defined by the combination of a name and in fields.
type Parameter struct {
	// Required. The name of the parameter. Parameter names are case
	// sensitive.
	Name string `json:"name"            yaml:"name"`
	// The location of the parameter. Possible values are "path",
	// "query", "header", or "cookie". When the step in context
	// specifies a workflowId, then all parameters map to workflow
	// inputs. In all other scenarios (e.g., a step specifies an
	// operationId), the in field MUST be specified.
	In *ParameterLocation `json:"in,omitempty"    yaml:"in,omitempty"`
	// Required. The value to pass in the parameter. The value can be
	// a constant or a Runtime Expression to be evaluated and passed
	// to the referenced operation or workflow.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"     yaml:"-,omitempty"`
}

// ParameterLocation is a string that represents the location of a
// parameter.
type ParameterLocation string

const (
	// ParameterLocationPath represents a path parameter. Used
	// together with OpenAPI style Path Templating, where the
	// parameter value is actually part of the operation’s URL. This
	// does not include the host or base path of the API. For example,
	// in /items/{itemId}, the path parameter is itemId.
	ParameterLocationPath ParameterLocation = "path"
	// ParameterLocationQuery represents a query parameter. Parameters
	// that are appended to the URL. For example, in /items?id=###,
	// the query parameter is id.
	ParameterLocationQuery ParameterLocation = "query"
	// ParameterLocationHeader represents a header parameter. Custom
	// headers that are expected as part of the request. Note that
	// [RFC9110] Name field names states field names (which includes
	// header) are case-insensitive.
	//
	// [RFC9110]: https://tools.ietf.org/html/rfc9110
	ParameterLocationHeader ParameterLocation = "header"
	// ParameterLocationCookie represents a cookie parameter. Used to
	// pass a specific cookie value to the source API.
	ParameterLocationCookie ParameterLocation = "cookie"
)

func (p ParameterLocation) ToPtr() *ParameterLocation {
	return &p
}

// Reusable is a struct that represents an Arazzo specification 1.0.X
// reusable object.
//
// A reusable object is simple object to allow referencing of objects
// contained within the Components Object. It can be used from
// locations within steps or workflows in the Arazzo description.
type Reusable struct {
	// Required. A Runtime Expression used to reference the desired
	// object.
	Reference string `json:"reference"       yaml:"reference"`
	// Sets a value of the referenced parameter. This is only
	// applicable for parameter object references.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

// RequestBody is a struct that represents an Arazzo specification
// 1.0.X request body object.
//
// A request body object describes the Content-Type and content to be
// passed by a step to an operation.
type RequestBody struct {
	// The Content-Type for the request content. If omitted then refer
	// to Content-Type specified at the targeted operation to
	// understand serialization requirements.
	ContentType *string `json:"contentType,omitempty"  yaml:"contentType,omitempty"`
	// A value representing the request body payload. The value can be
	// a literal value or can contain Runtime Expressions which MUST
	// be evaluated prior to calling the referenced operation. To
	// represent examples of media types that cannot be naturally
	// represented in JSON or YAML, use a string value to contain the
	// example, escaping where necessary.
	Payload any `json:"payload,omitempty"      yaml:"payload,omitempty"`
	// A list of locations and values to set within a payload.
	Replacements []PayloadReplacement `json:"replacements,omitempty" yaml:"replacements,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"            yaml:"-,omitempty"`
}

// PayloadReplacement is a struct that represents an Arazzo
// specification 1.0.X payload replacement object.
//
// A payload replacement object describes a location within a payload
// (e.g., a request body) and a value to set within the location.
type PayloadReplacement struct {
	// Required. A [JSON Pointer] or [XPath Expression] which MUST be
	// resolved against the request body. Used to identify the
	// location to inject the value.
	//
	// [JSON Pointer]: https://tools.ietf.org/html/rfc6901
	// [XPath Expression]:
	// https://www.w3.org/TR/xpath-31/#id-expressions
	Target string `json:"target"      yaml:"target"`
	// Required. The value set within the target location. The value
	// can be a constant or a Runtime Expression to be evaluated and
	// passed to the referenced operation or workflow.
	Value any `json:"value"       yaml:"value"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty" yaml:"-,omitempty"`
}

// Criterion is a struct that represents an Arazzo specification 1.0.X
// criterion object.
//
// A criterion object is An object used to specify the context,
// conditions, and condition types that can be used to prove or
// satisfy assertions specified in [Step] Object successCriteria,
// [SuccessAction] Object criteria, and [FailureAction] Object
// criteria.
type Criterion struct {
	// A Runtime Expression used to set the context for the condition
	// to be applied on. If type is specified, then the context MUST
	// be provided (e.g. $response.body would set the context that a
	// JSONPath query expression could be applied to).
	Context *string `json:"context,omitempty" yaml:"context,omitempty"`
	// Required. The condition to apply. Conditions can be simple
	// (e.g. $statusCode == 200 which applies an operator on a value
	// obtained from a runtime expression), or a regex, or a JSONPath
	// expression. For regex or JSONPath, the type and context MUST be
	// specified.
	Condition string `json:"condition"         yaml:"condition"`
	// The type of condition to be applied. If specified, the options
	// allowed are simple, regex, jsonpath or xpath. If omitted, then
	// the condition is assumed to be simple, which at most combines
	// literals, operators and Runtime Expressions. If jsonpath, then
	// the expression MUST conform to JSONPath. If xpath the
	// expression MUST conform to XML Path Language 3.1. Should other
	// variants of JSONPath or XPath be required, then a Criterion
	// Expression Type Object MUST be specified.
	Type *CriterionTypeOrCriterionExpressionType `json:"type,omitempty"    yaml:"type,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"       yaml:"-,omitempty"`
}

// CriterionTypeOrCriterionExpressionType allows a criterion to use
// either a [CriterionType] or a [CriterionExpressionType] object.
type CriterionTypeOrCriterionExpressionType struct {
	CriterionType           *CriterionType           `json:",omitempty" yaml:",omitempty"`
	CriterionExpressionType *CriterionExpressionType `json:",omitempty" yaml:",omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (c *CriterionTypeOrCriterionExpressionType) UnmarshalJSON(
	data []byte,
) error {
	var criterionType CriterionType
	if err := json.Unmarshal(data, &criterionType); err == nil {
		c.CriterionType = &criterionType
		return nil
	}

	var criterionExpressionType CriterionExpressionType
	if err := json.Unmarshal(data, &criterionExpressionType); err == nil {
		c.CriterionExpressionType = &criterionExpressionType
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (CriterionType, CriterionExpressionType)",
	)
}

// MarshalJSON implements json.Marshaler interface.
func (c CriterionTypeOrCriterionExpressionType) MarshalJSON() ([]byte, error) {
	if c.CriterionType != nil {
		return json.Marshal(c.CriterionType)
	}
	if c.CriterionExpressionType != nil {
		return json.Marshal(c.CriterionExpressionType)
	}
	return nil, errors.New("no data to marshal")
}

// UnmarshalYAML implements yaml.Unmarshaler interface.
func (c *CriterionTypeOrCriterionExpressionType) UnmarshalYAML(
	value *yaml.Node,
) error {
	key := value.Content[0].Value

	if key == "type" {
		var criterionType CriterionType
		if err := value.Decode(&criterionType); err == nil {
			c.CriterionType = &criterionType
			return nil
		}
	}

	var criterionExpressionType CriterionExpressionType
	if err := value.Decode(&criterionExpressionType); err == nil {
		c.CriterionExpressionType = &criterionExpressionType
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (CriterionType, CriterionExpressionType)",
	)
}

// MarshalYAML implements yaml.Marshaler interface.
func (c CriterionTypeOrCriterionExpressionType) MarshalYAML() (interface{}, error) {
	if c.CriterionType != nil {
		return c.CriterionType, nil
	}
	if c.CriterionExpressionType != nil {
		return c.CriterionExpressionType, nil
	}
	return nil, errors.New("no data to marshal")
}

// CriterionType is a string that represents the type of criterion
// available in an Arazzo specification 1.0.X.
type CriterionType string

const (
	// CriterionTypeSimple represents a simple criterion (e.g.
	// $statusCode == 200 which applies an operator on a value
	// obtained from a runtime expression).
	CriterionTypeSimple CriterionType = "simple"
	// CriterionTypeRegex represents a regex criterion.
	CriterionTypeRegex CriterionType = "regex"
	// CriterionTypeJsonPath represents a JSONPath criterion.
	CriterionTypeJsonPath CriterionType = "jsonpath"
	// CriterionTypeXPath represents an XPath criterion.
	CriterionTypeXPath CriterionType = "xpath"
)

func (c CriterionType) ToPtr() *CriterionType {
	return &c
}

// CriterionExpressionType is a struct that represents an Arazzo
// specification 1.0.X criterion expression type object. A criterion
// expression type object is an object used to describe the type and
// version of an expression used within a Criterion Object. If this
// object is not defined, then the following defaults apply:
//   - JSONPath as described by [RFC9535]
//   - XPath as described by [XML Path Language 3.1]
//
// [RFC9535]: https://tools.ietf.org/html/rfc9535
// [XML Path Language 3.1]: https://www.w3.org/TR/xpath-31/
type CriterionExpressionType struct {
	// Required. The type of condition to be applied. The options
	// allowed are jsonpath or xpath.
	Type CriterionExpressionTypeType `json:"type"        yaml:"type"`
	// Required. A short hand string representing the version of the
	// expression type being used. The allowed values for JSONPath are
	// draft-goessner-dispatch-jsonpath-00. The allowed values for
	// XPath are xpath-30, xpath-20, or xpath-10.
	Version string `json:"version"     yaml:"version"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty" yaml:"-,omitempty"`
}

// CriterionExpressionTypeType is a string that represents the type of
// criterion expression type available in an Arazzo specification
// 1.0.X.
type CriterionExpressionTypeType string

const (
	// CriterionExpressionTypeTypeJsonPath represents a JSONPath
	// criterion expression type.
	CriterionExpressionTypeTypeJsonPath CriterionExpressionTypeType = "jsonpath"
	// CriterionExpressionTypeTypeXPath represents an XPath criterion
	// expression type.
	CriterionExpressionTypeTypeXPath CriterionExpressionTypeType = "xpath"
)

// SuccessAction is a struct that represents an Arazzo specification
// 1.0.X success action object.
//
// A success action object is a single success action which describes
// an action to take upon success of a workflow step.
type SuccessAction struct {
	// Required. The name of the success action. Names are case
	// sensitive.
	Name string `json:"name"                 yaml:"name"`
	// Required. The type of action to take. Possible values are "end"
	// or "goto".
	Type SuccessActionType `json:"type"                 yaml:"type"`
	// The workflowId referencing an existing workflow within the
	// Arazzo Description to transfer to upon success of the step.
	// This field is only relevant when the type field value is
	// "goto". If the referenced workflow is contained within an
	// arazzo type sourceDescription, then the workflowId MUST be
	// specified using a Runtime Expression (e.g.,
	// $sourceDescriptions.<name>.<workflowId>) to avoid ambiguity or
	// potential clashes. This field is mutually exclusive to stepId.
	WorkflowId *string `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	// The stepId to transfer to upon success of the step. This field
	// is only relevant when the type field value is "goto". The
	// referenced stepId MUST be within the current workflow. This
	// field is mutually exclusive to workflowId.
	StepId *string `json:"stepId,omitempty"     yaml:"stepId,omitempty"`
	// A list of assertions to determine if this action SHALL be
	// executed. Each assertion is described using a Criterion Object.
	// All criteria assertions MUST be satisfied for the action to be
	// executed.
	Criteria []Criterion `json:"criteria,omitempty"   yaml:"criteria,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"          yaml:"-,omitempty"`
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
	SuccessAction *SuccessAction `json:",omitempty" yaml:",omitempty"`
	Reusable      *Reusable      `json:",omitempty" yaml:",omitempty"`
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

// UnmarshalYAML implements yaml.Unmarshaler interface.
func (s *SuccessActionOrReusable) UnmarshalYAML(
	value *yaml.Node,
) error {
	key := value.Content[0].Value

	if key == "reference" {
		var reusable Reusable
		if err := value.Decode(&reusable); err == nil {
			s.Reusable = &reusable
			return nil
		}
	}

	var successAction SuccessAction
	if err := value.Decode(&successAction); err == nil {
		s.SuccessAction = &successAction
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (SuccessAction, Reusable)",
	)
}

// MarshalYAML implements yaml.Marshaler interface.
func (s SuccessActionOrReusable) MarshalYAML() (interface{}, error) {
	if s.SuccessAction != nil {
		return s.SuccessAction, nil
	}
	if s.Reusable != nil {
		return s.Reusable, nil
	}
	return nil, errors.New("no data to marshal")
}

// FailureAction is a struct that represents an Arazzo specification
// 1.0.X failure action object.
//
// A failure action object is single failure action which describes an
// action to take upon failure of a workflow step.
type FailureAction struct {
	// Required. The name of the failure action. Names are case
	// sensitive.
	Name string `json:"name"                 yaml:"name"`
	// Required. The type of action to take. Possible values are
	// "end", "retry", or "goto".
	Type FailureActionType `json:"type"                 yaml:"type"`
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
	WorkflowId *string `json:"workflowId,omitempty" yaml:"workflowId,omitempty"`
	// The stepId to transfer to upon failure of the step. This field
	// is only relevant when the type field value is "goto" or
	// "retry". The referenced stepId MUST be within the current
	// workflow. This field is mutually exclusive to workflowId. When
	// used with "retry", context transfers back upon completion of
	// the specified step.
	StepId *string `json:"stepId,omitempty"     yaml:"stepId,omitempty"`
	// A non-negative decimal indicating the seconds to delay after
	// the step failure before another attempt SHALL be made. Note: if
	// an HTTP Retry-After response header was returned to a step from
	// a targeted operation, then it SHOULD overrule this particular
	// field value. This field only applies when the type field value
	// is "retry".
	RetryDelay *float64 `json:"retryDelay,omitempty" yaml:"retryDelay,omitempty"`
	// A non-negative integer indicating how many attempts to retry
	// the step MAY be attempted before failing the overall step. If
	// not specified then a single retry SHALL be attempted. This
	// field only applies when the type field value is "retry". The
	// retryLimit MUST be exhausted prior to executing subsequent
	// failure actions.
	RetryLimit *int `json:"retryLimit,omitempty" yaml:"retryLimit,omitempty"`
	// A list of assertions to determine if this action SHALL be
	// executed. Each assertion is described using a Criterion Object.
	// All criteria assertions MUST be satisfied for the action to be
	// executed
	Criteria []Criterion `json:"criteria,omitempty"   yaml:"criteria,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"          yaml:"-,omitempty"`
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
	FailureAction *FailureAction `json:",omitempty" yaml:",omitempty"`
	Reusable      *Reusable      `json:",omitempty" yaml:",omitempty"`
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

// UnmarshalYAML implements yaml.Unmarshaler interface.
func (f *FailureActionOrReusable) UnmarshalYAML(
	value *yaml.Node,
) error {
	key := value.Content[0].Value

	if key == "reference" {
		var reusable Reusable
		if err := value.Decode(&reusable); err == nil {
			f.Reusable = &reusable
			return nil
		}
	}

	var failureAction FailureAction
	if err := value.Decode(&failureAction); err == nil {
		f.FailureAction = &failureAction
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (FailureAction, Reusable)",
	)
}

// MarshalYAML implements yaml.Marshaler interface.
func (f FailureActionOrReusable) MarshalYAML() (interface{}, error) {
	if f.FailureAction != nil {
		return f.FailureAction, nil
	}
	if f.Reusable != nil {
		return f.Reusable, nil
	}
	return nil, errors.New("no data to marshal")
}

// Components is a struct that represents an Arazzo specification
// 1.0.X components object.
//
// A component object holds a set of reusable
// objects for different aspects of the Arazzo Specification. All
// objects defined within the components object will have no effect on
// the Arazzo Description unless they are explicitly referenced from
// properties outside the components object.
//
// Components are scoped to the Arazzo document they are defined in.
// For example, if a step defined in Arazzo document “A”
// references a workflow defined in Arazzo document “B”, the
// components in “A” are not considered when evaluating the
// workflow referenced in “B”.
type Components struct {
	// An object to hold reusable JSON Schema objects to be referenced
	// from workflow inputs.
	Inputs map[string]any `json:"inputs,omitempty"         yaml:"inputs,omitempty"`
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter `json:"parameters,omitempty"     yaml:"parameters,omitempty"`
	// An object to hold reusable Success Action Objects.
	SuccessActions map[string]SuccessAction `json:"successActions,omitempty" yaml:"successActions,omitempty"`
	// An object to hold reusable Failure Action Objects.
	FailureActions map[string]FailureAction `json:"failureActions,omitempty" yaml:"failureActions,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"              yaml:"-,omitempty"`
}
