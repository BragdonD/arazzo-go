package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pb33f/libopenapi"
	oai31 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// FileScheme represents the "file" URL scheme for local file access.
const (
	FileScheme = "file"
)

// HTTPMethod defines a type for HTTP methods.
type HTTPMethod string

// Constants for different HTTP methods, matching the net/http
// package.
const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

// OAIOperation represents an OpenAPI operation with its associated
// path and HTTP method.
type OAIOperation struct {
	Path      string
	Method    HTTPMethod
	Operation *oai31.Operation
}

// OAIDocument holds an OpenAPI document model and its source URL.
type OAIDocument = libopenapi.Document

// ExtractOperationsFromOpenAPI extracts API operations from an
// OpenAPI document.
func ExtractOperationsFromOpenAPI(
	oaiDoc OAIDocument,
) ([]*OAIOperation, error) {
	v3Doc, errs := oaiDoc.BuildV3Model()
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	operations := []*OAIOperation{}

	if oaiDoc == nil {
		return operations, fmt.Errorf("openapi document is nil")
	}

	for path, pathItem := range v3Doc.Model.Paths.PathItems.FromNewest() {
		if pathItem == nil {
			continue
		}
		methods := map[HTTPMethod]*oai31.Operation{
			MethodGet:     pathItem.Get,
			MethodHead:    pathItem.Head,
			MethodPost:    pathItem.Post,
			MethodPut:     pathItem.Put,
			MethodPatch:   pathItem.Patch,
			MethodDelete:  pathItem.Delete,
			MethodOptions: pathItem.Options,
			MethodTrace:   pathItem.Trace,
		}

		for method, operation := range methods {
			if operation != nil {
				operations = append(operations, &OAIOperation{
					Path:      path,
					Method:    method,
					Operation: operation,
				})
			}
		}
	}

	return operations, nil
}

// GetOperationById searches for an OpenAPI operation by its
// OperationId.
// func (idx *Index) GetOperationById(
// 	operationId string,
// ) (*OAIOperation, error) {
// 	for _, operation := range idx.operations {
// 		if operation.Operation.OperationId == operationId {
// 			return operation, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("operation not found")
// }
