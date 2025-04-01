package v1

import (
	"github.com/bragdonD/arazzo-go/v1/models"
)

type Parameter struct {
	model *models.Parameter
	value *Value
}

func NewParameter(param *models.Parameter) *Parameter {
	return &Parameter{
		model: param,
		value: NewValue(param.Value),
	}
}

func (p *Parameter) GetModel() *models.Parameter {
	return p.model
}

func (p *Parameter) GetLocation() models.ParameterLocation {
	return *p.model.In
}

func (p *Parameter) GetName() string {
	return p.model.Name
}

func (p *Parameter) GetValue() *Value {
	return p.value
}
