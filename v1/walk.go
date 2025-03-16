package v1

import "fmt"

// ResolveOperations checks the operation references in the workflow steps
// and resolves them to the actual operations. If the operation is not found
// an error is returned.
func (s *Spec) ResolveOperations() error {
	oaiDocs := []*OAIDocument{}

	for _, source := range s.SourcesDescriptions {
		if source.Type == SourceDescriptionTypeOpenAPI.ToPtr() {
			doc, err := LoadOpenAPISource(source.Url)
			if err != nil {
				return err
			}
			oaiDocs = append(oaiDocs, doc)
		}
	}

	operations := []*OAIOperation{}

	for _, doc := range oaiDocs {
		ops, err := ExtractOperationsFromOpenAPI(doc)
		if err != nil {
			return err
		}
		operations = append(operations, ops...)
	}

	for _, workflows := range s.Workflows {
		for _, step := range workflows.Steps {
			if step.OperationId != nil {
				// TODO: we need to check for duplicated operationId
				// across all the loaded documents
				_, err := GetOperationById(*step.OperationId, operations)
				if err != nil {
					return err
				}
				// TODO: we should probably build an index like
				// libopenapi does to represent the documents
				// loaded values
			} else if step.OperationPath != nil {
				// TODO: implement this part using the expression
				// parser and a jsonpointer library
				return fmt.Errorf("step %s has operationPath, not implemented", step.StepId)
			} else {
				return fmt.Errorf("step %s has no operationId or operationPath", step.StepId)
			}
		}
	}

	return nil
}
