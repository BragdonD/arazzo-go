package v1

import (
	"github.com/bragdonD/arazzo-go/v1/models"
)

type Document struct {
	Version string
	Model   models.Spec
	Index   *Index
}

// NewDocument initializes a new Document instance.
func NewDocument(specByteArray []byte) (*Document, error) {
	spec, err := models.ExtractSpecWithDocumentCheck(specByteArray)
	if err != nil {
		return nil, err
	}

	Document := &Document{
		Version: spec.Arazzo,
		Model:   *spec,
		Index:   nil,
	}

	return Document, nil
}
