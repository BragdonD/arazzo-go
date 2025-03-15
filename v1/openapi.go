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

const (
	FileScheme = "file"
)

type HTTPMethod string

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

type OAIOperation struct {
	ServerUrl string
	Path      string
	Method    HTTPMethod
	Operation *oai31.Operation
}

func LoadOpenAPISource(oaiUrl string) (*libopenapi.DocumentModel[oai31.Document], error) {
	oaiContent := []byte{}
	parsedURL, err := url.Parse(oaiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	// An URL is considered local if:
	// - It has no scheme (e.g. "myfile.txt", "/home/user/myfile.txt")
	// - It has the "file" scheme (e.g. "file:///home/user/file.txt")
	if parsedURL.Scheme == "" || parsedURL.Scheme == FileScheme {
		path := filepath.FromSlash(parsedURL.Path)
		oaiContent, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
	} else {

	}

	openapiDoc, err := libopenapi.NewDocument(oaiContent)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize new openapi doc: %v", err)
	}

	model, errs := openapiDoc.BuildV3Model()

	return model, errors.Join(errs...)
}

func ExtractOperationsFromOpenAPI(oaiDoc *libopenapi.DocumentModel[oai31.Document]) ([]OAIOperation, error) {
	operations := []OAIOperation{}

	if oaiDoc == nil {
		return operations, fmt.Errorf("openapi document is nil")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := orderedmap.Iterate(ctx, oaiDoc.Model.Paths.PathItems)

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
				operations = append(operations, OAIOperation{
					ServerUrl: oaiDoc.Model.Servers[0].URL,
					Path:      path,
					Method:    method,
					Operation: operation,
				})
			}
		}
	}

	return operations, nil
}

func GetOperationById(operationId string, oaiDoc *libopenapi.DocumentModel[oai31.Document]) {

}
