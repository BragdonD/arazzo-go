package v1

import (
	"encoding/json"
	"errors"
)

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
	Context *string `json:"context,omitempty"`
	// Required. The condition to apply. Conditions can be simple
	// (e.g. $statusCode == 200 which applies an operator on a value
	// obtained from a runtime expression), or a regex, or a JSONPath
	// expression. For regex or JSONPath, the type and context MUST be
	// specified.
	Condition string `json:"condition"`
	// The type of condition to be applied. If specified, the options
	// allowed are simple, regex, jsonpath or xpath. If omitted, then
	// the condition is assumed to be simple, which at most combines
	// literals, operators and Runtime Expressions. If jsonpath, then
	// the expression MUST conform to JSONPath. If xpath the
	// expression MUST conform to XML Path Language 3.1. Should other
	// variants of JSONPath or XPath be required, then a Criterion
	// Expression Type Object MUST be specified.
	Type *CriterionTypeOrCriterionExpressionType `json:"type,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

// CriterionTypeOrCriterionExpressionType allows a criterion to use
// either a [CriterionType] or a [CriterionExpressionType] object.
type CriterionTypeOrCriterionExpressionType struct {
	CriterionType           *CriterionType           `json:",omitempty"`
	CriterionExpressionType *CriterionExpressionType `json:",omitempty"`
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
	Type CriterionExpressionTypeType `json:"type"`
	// Required. A short hand string representing the version of the
	// expression type being used. The allowed values for JSONPath are
	// draft-goessner-dispatch-jsonpath-00. The allowed values for
	// XPath are xpath-30, xpath-20, or xpath-10.
	Version string `json:"version"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
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
