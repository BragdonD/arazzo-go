package v1

import (
	"fmt"

	"github.com/bragdonD/arazzo-go/v1/models"
	jsonpointergo "github.com/bragdond/jsonpointer-go"
)

// PayloadReplacement is a struct that represents an Arazzo
// specification 1.0.X payload replacement object.
//
// A payload replacement object describes a location within a payload
// (e.g., a request body) and a value to set within the location.
type PayloadReplacement struct {
	model  *models.PayloadReplacement
	target *jsonpointergo.JSONPointer
	value  *Value
}

func NewPayloadReplacement(pr *models.PayloadReplacement) (*PayloadReplacement, error) {
	pointer, err := jsonpointergo.NewJSONPointer(pr.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to create new json"+
			" pointer: %v", err)
	}
	return &PayloadReplacement{
		model:  pr,
		target: pointer,
		value:  NewValue(pr.Value),
	}, nil
}
