package v1

import (
	"errors"
	"log/slog"
	"net/url"

	"github.com/bragdonD/arazzo-go"
	"github.com/bragdonD/arazzo-go/v1/models"
	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

// IndexConfig is a configuration structure for Index that
// provides an extensible set of granular options. The first is the
// ability to allow or disallow local or remote files to be consulted.
type IndexConfig struct {
	// allowRemoteLookup allows the index to look for files not found
	// locally in the base URL. The default value is false to prevent
	// remote exploitation if the remote reference is malicious.
	allowRemoteLookup bool
	// allowLocalLookup allows the index to look for files not found
	// locally in the base URL. The default value is false to prevent
	// local exploitation if the local reference is malicious.
	// This is to prevent log4j-like attacks.
	allowLocalLookup bool
	// logger is a logger that will be used for logging errors and
	// warnings. If not set, the default logger will be used, set to
	// the Error level.
	logger *slog.Logger
	// spec is a pointer to the Spec struct that contains the Arazzo
	// specification model.
	spec *models.Spec
}

// Index is a struct that holds all the values extracted from
// an Spec object.
type Index struct {
	// config is the configuration for the index.
	config *IndexConfig
	// baseURL will be the root from which relative references will be
	// resolved if they cannot be found locally.
	//
	// For example:
	//  - url: ./somefile.yaml
	//
	// The file may not be found locally if it has been retrieved from
	// a remote server. So, by defining a BaseURL, the reference will
	// try to be resolved from the remote server.
	//
	// If our baseURL is set to https://github.com/BragdonD/arazzo-go
	// then our url will try to be resolved from:
	//   - url: https://github.com/BragdonD/arazzo-go/somefile.yaml
	baseURL *url.URL
	// filename is the name of the Arazzo Specification file. Usually
	// named <name>.arazzo.yaml or <name>.arazzo.json. Where <name> is
	// optional.
	filename string
	// spec is the actual arazzo specification model.
	spec *models.Spec
	// openapiDocuments is a list of OpenAPI documents that are
	// referenced by the Arazzo specification.
	openapiDocuments []OAIDocument
	// arazzoDocuments is a list of Arazzo documents that are
	// referenced by the Arazzo specification.
	arazzoDocuments []*Document
	// workflows is a list of workflows that are referenced by the
	// Arazzo specification.
	workflows []*WorkflowIndex
}

type WorkflowIndex struct {
	// model is the actual workflow model.
	model *models.Workflow
	// index is a reference to the index parent of this workflow.
	index *Index
	// dependencies is a list of workflows to run before this
	// workflow.
	dependencies []*WorkflowIndex
	// parameters is a list of resolved parameters for this workflow.
	// In the model, parameters are defined as either a parameter
	// or a reusable object. This parameters are applied to the
	// steps in the workflow.
	parameters []*models.Parameter
	// steps is a list of resolved steps for this workflow.
	steps []*StepIndex
}

type StepIndex struct {
	// model is the actual step model.
	model *models.Step
	// openapiOperation is a reference to the OpenAPI operation
	// that this step is associated with.
	openapiOperation *OAIOperation
	// workflowIndex is a reference to the workflow index parent
	// of this step.
	workflowIndex *WorkflowIndex
	// parameters is a list of resolved parameters for this step.
	// In the model, parameters are defined as either a parameter
	// or a reusable object.
	parameters []*models.Parameter
	// TODO: probably need some sort of way to represent the actions
	// to perform on success or failure.
}

func buildIndexFromConfig(config *IndexConfig) (*Index, error) {
	idx := &Index{}
	idx.config = config

	options := []arazzo.LoaderOption{}
	if config.allowRemoteLookup {
		options = append(options, arazzo.AllowRemoteLookup())
	}
	if config.allowLocalLookup {
		options = append(options, arazzo.AllowLocalLookup())
	}
	// TODO: think of a way to set the http client for the loader.
	loader := arazzo.NewLoader(options...)

	// Load all the sources to populate the index.
	for _, source := range config.spec.SourcesDescriptions {
		data, err := loader.LoadFile(source.Url)
		if err != nil {
			return nil, err
		}
		if *source.Type == models.SourceDescriptionTypeOpenAPI {
			doc, err := libopenapi.NewDocument(data)
			if err != nil {
				return nil, err
			}

			docValidator, validatorErrs := validator.NewValidator(doc)
			if validatorErrs != nil {
				return nil, errors.Join(validatorErrs...)
			}

			valid, errs := docValidator.ValidateDocument()
			if !valid {
				var err error
				for _, e := range errs {
					err = errors.Join(err, e)
				}
				return nil, err
			}

			idx.openapiDocuments = append(idx.openapiDocuments, doc)
		} else if *source.Type == models.SourceDescriptionTypeArazzo {
			doc, err := NewDocument(data)
			if err != nil {
				return nil, err
			}

			// TODO: validate the Arazzo document.
			idx.arazzoDocuments = append(idx.arazzoDocuments, doc)
		}
	}

	// Build the workflows.
	for _, workflow := range config.spec.Workflows {
		parameterOrReusables := workflow.Parameters
		parameters := []*models.Parameter{}
		for _, por := range parameterOrReusables {
			param, err := por.ToParameter(config.spec.Components)
			if err != nil {
				return nil, err
			}
			parameters = append(parameters, param)
		}

		wfi := &WorkflowIndex{
			model:        &workflow,
			index:        idx,
			dependencies: []*WorkflowIndex{}, // this needs to be populated later.
			parameters:   parameters,
			steps:        []*StepIndex{}, // this needs to be populated later.
		}
		idx.workflows = append(idx.workflows, wfi)
	}

	return idx, nil
}
