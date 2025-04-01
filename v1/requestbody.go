package v1

import "github.com/bragdonD/arazzo-go/v1/models"

// RequestBody is a struct that represents an Arazzo specification
// 1.0.X request body object.
//
// A request body object describes the Content-Type and content to be
// passed by a step to an operation.
type RequestBody struct {
	model        *models.RequestBody
	payload      string
	replacements []*PayloadReplacement
}
