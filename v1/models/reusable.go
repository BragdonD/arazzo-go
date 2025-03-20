package models

import (
	"errors"

	"github.com/bragdonD/arazzo-go/v1/expression"
)

// Reusable is a struct that represents an Arazzo specification 1.0.X
// reusable object.
//
// A reusable object is simple object to allow referencing of objects
// contained within the Components Object. It can be used from
// locations within steps or workflows in the Arazzo description.
type Reusable struct {
	// Required. A Runtime Expression used to reference the desired
	// object.
	Reference string `json:"reference"`
	// Sets a value of the referenced parameter. This is only
	// applicable for parameter object references.
	Value string `json:"value,omitempty"`
}

// reusableToParameterVisitor is a struct that helps in resolving
// references from Reusable objects to Parameter objects within
// the Components object.
type reusableToParameterVisitor struct {
	// err stores any encountered error during traversal.
	err error
	// parameters holds the resolved parameter if successful.
	parameter *Parameter
}

// VisitExpressionWithNameNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitExpressionWithNameNode(expr *expression.ExpressionWithNameNode) any {
	if expr.Value != expression.ABNFExpressionComponentsParameters {
		r.err = errors.New("expected $components.parameters.<name>")
		r.parameter = nil
		return nil
	}

	r.parameter = &Parameter{
		Name: expr.Name.Value,
	}
	return nil
}

// VisitSingleExpressionNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitSingleExpressionNode(*expression.SingleExpressionNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitExpressionWithSourceNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitExpressionWithSourceNode(*expression.ExpressionWithSourceNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitHeaderReferenceNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitHeaderReferenceNode(*expression.HeaderReferenceNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitQueryReferenceNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitQueryReferenceNode(*expression.QueryReferenceNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitPathReferenceNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitPathReferenceNode(*expression.PathReferenceNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitBodyReferenceNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitBodyReferenceNode(*expression.BodyReferenceNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitNameNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitNameNode(*expression.NameNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitTokenNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitTokenNode(*expression.TokenNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// VisitJSONPointerNode implements the Visitor interface for
// expression.
func (r *reusableToParameterVisitor) VisitJSONPointerNode(*expression.JSONPointerNode) any {
	r.err = errors.New("expected $components.parameters.<name>")
	r.parameter = nil
	return nil
}

// ToParameter resolves a reusable object reference to a Parameter object
// within the Components object.
//
// This function parses the reference expression, verifies that it conforms
// to the expected format, and retrieves the corresponding parameter from
// the Components object. If a value is set in the reusable object, it is
// assigned to the parameter.
func (r *Reusable) ToParameter(components *Components) (*Parameter, error) {
	if components == nil {
		return nil, errors.New("components is nil")
	}
	if r.Reference == "" {
		return nil, errors.New("reference is empty")
	}

	// Extract the expression from the reference.
	exprStr := r.Reference
	var err error
	if exprStr[0] == '{' {
		exprStr, err = expression.Extract(exprStr)
		if err != nil {
			return nil, err
		}
	}
	expr, err := expression.Parse(exprStr)
	if err != nil {
		return nil, err
	}

	// Resolve the parameter from the expression.
	visitor := &reusableToParameterVisitor{}
	if expr.Accept(visitor); visitor.err != nil {
		return nil, visitor.err
	}
	parameter, ok := components.Parameters[visitor.parameter.Name]
	if !ok {
		return nil, errors.New("parameter not found")
	}

	// Set the value if it is present.
	if r.Value != "" {
		parameter.Value = r.Value
	}
	return &parameter, nil
}
