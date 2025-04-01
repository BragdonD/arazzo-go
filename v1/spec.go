package v1

import "github.com/bragdonD/arazzo-go/v1/models"

type Spec struct {
	model      *models.Spec
	url        string
	workflows  []*Workflow
	components *Components
	oaiDocs    []*OAIDocument
	// arazzo documents will be handled in a future version
	// arazzoDocs []*Spec // TODO: Find a way to break out of circular dependencies
}

func NewSpec(model *models.Spec, url string) (*Spec, error) {
	spec := &Spec{
		model:      model,
		url:        url,
		workflows:  []*Workflow{},
		components: NewComponents(model.Components),
		oaiDocs:    []*OAIDocument{},
	}

	for _, source := range model.SourcesDescriptions {
		if *source.Type == models.SourceDescriptionTypeOpenAPI {
			doc, err := NewOAIDocument(source.Url)
			if err != nil {
				return nil, err
			}
			spec.oaiDocs = append(spec.oaiDocs, doc)
		}
		// TODO: Handle arazzo source types
	}

	return spec, nil
}

func (s *Spec) GetComponents() *Components {
	return s.components
}
