package v1

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pb33f/libopenapi"
	oai31 "github.com/pb33f/libopenapi/datamodel/high/v3"
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

func GetOperationById(operationId string, oaiDoc *libopenapi.DocumentModel[oai31.Document]) {

}
