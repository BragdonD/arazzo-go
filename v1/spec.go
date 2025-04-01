package v1

import "github.com/bragdonD/arazzo-go/v1/models"

type Spec struct {
	model      *models.Spec
	workflows  []*Workflow
	components *Components
	oaiDocs    []*OAIDocument
	// arazzo documents will be handled in a future version
	// arazzoDocs []*Spec // TODO: Find a way to break out of circular dependencies
}
