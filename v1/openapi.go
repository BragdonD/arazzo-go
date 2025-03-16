package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pb33f/libopenapi"
	oai31 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
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
type OAIDocument struct {
	URL   string
	Model *libopenapi.DocumentModel[oai31.Document]
}

// LoadOpenAPISource loads an OpenAPI document from a given URL or
// file path.
func LoadOpenAPISource(oaiUrl string) (*OAIDocument, error) {
	oaiContent := []byte{}
	parsedURL, err := url.Parse(oaiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	// Check if the URL represents a local file.
	// A local file is identified by:
	// - Having no scheme (e.g., "myfile.yaml",
	// "/home/user/myfile.yaml")
	// - Using the "file" scheme (e.g., "file:///home/user/file.yaml")
	if parsedURL.Scheme == "" || parsedURL.Scheme == FileScheme {
		path := filepath.FromSlash(parsedURL.Path)
		oaiContent, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
	} else {
		// TODO: Implement remote file download functionality.
		return nil, fmt.Errorf("remote file download not implemented")
	}

	// Parse the OpenAPI document content.
	openapiDoc, err := libopenapi.NewDocument(oaiContent)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialize new openapi doc: %v",
			err,
		)
	}
	model, errs := openapiDoc.BuildV3Model()

	return &OAIDocument{
		URL:   oaiUrl,
		Model: model,
	}, errors.Join(errs...)
}

// ExtractOperationsFromOpenAPI extracts API operations from an
// OpenAPI document.
func ExtractOperationsFromOpenAPI(
	oaiDoc *OAIDocument,
) ([]*OAIOperation, error) {
	operations := []*OAIOperation{}

	if oaiDoc == nil {
		return operations, fmt.Errorf("openapi document is nil")
	}

	// Create a cancellable context for iterating over OpenAPI paths.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Iterate through the OpenAPI paths to extract operations.
	c := orderedmap.Iterate(ctx, oaiDoc.Model.Model.Paths.PathItems)

	for pair := range c {
		path := pair.Key()
		pathItem := pair.Value()

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
func GetOperationById(
	operationId string,
	operations []*OAIOperation,
) (*OAIOperation, error) {
	for _, operation := range operations {
		if operation.Operation.OperationId == operationId {
			return operation, nil
		}
	}

	return nil, fmt.Errorf("operation not found")
}
