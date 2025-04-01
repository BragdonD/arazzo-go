package v1

import (
	"fmt"

	"github.com/bragdonD/arazzo-go/v1/expression"
)

type Value struct {
	// val is either a runtime expression or a constant
	// expression
	val     any
	isConst bool
}

func NewValue(input any) *Value {
	value := &Value{
		isConst: false,
	}
	str, ok := input.(string)
	// If the value is not a string then it is a constant
	// value. Else, it can either be a runtime expression
	// or a constant string.
	if !ok {
		value.isConst = true
		value.val = input
	}
	var err error
	if len(str) > 0 && str[0] == '{' {
		str, err = expression.Extract(str)
		if err != nil {
			value.isConst = true
			value.val = input
		}
	}

	_, err = expression.Parse(str)
	if err != nil {
		value.isConst = true
		value.val = input
	}

	return value
}

func (v *Value) IsConstant() bool {
	return v.isConst
}

func (v *Value) Evaluate(spec *Spec) (any, error) {
	if v.isConst {
		return v.val, nil
	}
	str, ok := v.val.(string)
	if !ok {
		return nil, fmt.Errorf("the value being evaluated" +
			" is not a constant and is not a string to be" +
			" evaluated as a runtime expression")
	}
	var err error
	if len(str) > 0 && str[0] == '{' {
		str, err = expression.Extract(str)
		if err != nil {
			return nil, fmt.Errorf("failed to extract the"+
				" runtime expression: %v", err)
		}
	}

	expr, err := expression.Parse(str)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the"+
			" runtime expression: %v", err)
	}

	return ResolveRuntimeExpression(expr, spec)
}
