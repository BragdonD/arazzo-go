package models_test

import (
	"encoding/json"
	"testing"

	v1 "github.com/bragdonD/arazzo-go/v1/models"
	"github.com/go-test/deep"
	"sigs.k8s.io/yaml"
)

// strPtr is a helper function to create string pointers.
func strPtr(s string) *string {
	return &s
}

// yamlEqual is a helper function to compare two YAML strings.
func yamlEqual(yaml1, yaml2 string) (bool, error) {
	var obj1, obj2 interface{}

	if err := yaml.Unmarshal([]byte(yaml1), &obj1); err != nil {
		return false, err
	}

	if err := yaml.Unmarshal([]byte(yaml2), &obj2); err != nil {
		return false, err
	}

	return deep.Equal(obj1, obj2) == nil, nil
}

// jsonEqual is a helper function to compare two JSON strings.
func jsonEqual(json1, json2 string) (bool, error) {
	var obj1, obj2 interface{}

	if err := json.Unmarshal([]byte(json1), &obj1); err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(json2), &obj2); err != nil {
		return false, err
	}

	return deep.Equal(obj1, obj2) == nil, nil
}

func TestSpec_UnmarshalYAML(t *testing.T) {
	yamlSpecData := `
arazzo: 1.0.0
info:
  title: "A pet purchasing workflow"
  summary: "This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls."
  description: "This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet."
  version: "1.0.1"
sourceDescriptions:
  - name: "petStoreDescription"
    url: "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml"
    type: "openapi"
workflows:
  - workflowId: "loginUserAndRetrievePet"
    summary: "Login User and then retrieve pets"
    description: "This workflow lays out the steps to login a user and then retrieve pets"
    inputs:
        type: object
        properties:
            username:
                type: string
            password:
                type: string
    steps:
        - stepId: loginStep
          description: This step demonstrates the user login step
          operationId: loginUser
          parameters:
              # parameters to inject into the loginUser operation (parameter name must be resolvable at the referenced operation and the value is determined using {expression} syntax)
              - name: username
                in: query
                value: $inputs.username
              - name: password
                in: query
                value: $inputs.password
          successCriteria:
              # assertions to determine step was successful
              - condition: $statusCode == 200
          outputs:
              # outputs from this step
              tokenExpires: $response.header.X-Expires-After
              rateLimit: $response.header.X-Rate-Limit
              sessionToken: $response.body
        - stepId: getPetStep
          description: Retrieve a pet by status from the GET pets endpoint
          operationPath: '{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get'
          parameters:
              - name: status
                in: query
                value: 'available'
              - name: Authorization
                in: header
                value: $steps.loginUser.outputs.sessionToken
          successCriteria:
              - condition: $statusCode == 200
          outputs:
              # outputs from this step
              availablePets: $response.body
    outputs:
      available: $steps.getPetStep.outputs.availablePets
`
	expectedSpec := v1.Spec{
		Arazzo: "1.0.0",
		Info: v1.Info{
			Title: "A pet purchasing workflow",
			Summary: strPtr(
				"This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
			),
			Description: strPtr(
				"This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
			),
			Version: "1.0.1",
		},
		SourcesDescriptions: []v1.SourceDescription{
			{
				Name: "petStoreDescription",
				Url:  "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
				Type: v1.SourceDescriptionTypeOpenAPI.ToPtr(),
			},
		},
		Workflows: []v1.Workflow{
			{
				WorkflowId: "loginUserAndRetrievePet",
				Summary: strPtr(
					"Login User and then retrieve pets",
				),
				Description: strPtr(
					"This workflow lays out the steps to login a user and then retrieve pets",
				),
				Inputs: map[string]interface{}{
					"type": "object",
					"properties": map[string]any{
						"username": map[string]any{"type": "string"},
						"password": map[string]any{"type": "string"},
					},
				},
				Steps: []v1.Step{
					{
						StepId: "loginStep",
						Description: strPtr(
							"This step demonstrates the user login step",
						),
						OperationId: strPtr("loginUser"),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "username",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.username",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "password",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.password",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"tokenExpires": "$response.header.X-Expires-After",
							"rateLimit":    "$response.header.X-Rate-Limit",
							"sessionToken": "$response.body",
						},
					},
					{
						StepId: "getPetStep",
						Description: strPtr(
							"Retrieve a pet by status from the GET pets endpoint",
						),
						OperationPath: strPtr(
							"{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
						),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "status",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "available",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "Authorization",
									In:    v1.ParameterLocationHeader.ToPtr(),
									Value: "$steps.loginUser.outputs.sessionToken",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"availablePets": "$response.body",
						},
					},
				},
				Outputs: map[string]any{
					"available": "$steps.getPetStep.outputs.availablePets",
				},
			},
		},
	}

	var spec v1.Spec

	// Unmarshal YAML data into the Spec struct
	if err := yaml.Unmarshal([]byte(yamlSpecData), &spec); err != nil {
		t.Fatalf("could not unmarshal the test's data: %v", err)
	}

	// Compare the expected and actual Spec structs
	if diff := deep.Equal(spec, expectedSpec); diff != nil {
		t.Fatalf(
			"expected and actual Spec structs are not equal: %v",
			diff,
		)
	}
}

func TestSpec_MarshalYAML(t *testing.T) {
	expectedYAMLData := `arazzo: 1.0.0
info:
  title: "A pet purchasing workflow"
  summary: "This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls."
  description: "This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet."
  version: "1.0.1"
sourceDescriptions:
  - name: "petStoreDescription"
    url: "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml"
    type: "openapi"
workflows:
  - workflowId: "loginUserAndRetrievePet"
    summary: "Login User and then retrieve pets"
    description: "This workflow lays out the steps to login a user and then retrieve pets"
    inputs:
        type: object
        properties:
            username:
                type: string
            password:
                type: string
    steps:
        - stepId: loginStep
          description: This step demonstrates the user login step
          operationId: loginUser
          parameters:
              # parameters to inject into the loginUser operation (parameter name must be resolvable at the referenced operation and the value is determined using {expression} syntax)
              - name: username
                in: query
                value: $inputs.username
              - name: password
                in: query
                value: $inputs.password
          successCriteria:
              # assertions to determine step was successful
              - condition: $statusCode == 200
          outputs:
              # outputs from this step
              tokenExpires: $response.header.X-Expires-After
              rateLimit: $response.header.X-Rate-Limit
              sessionToken: $response.body
        - stepId: getPetStep
          description: Retrieve a pet by status from the GET pets endpoint
          operationPath: '{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get'
          parameters:
              - name: status
                in: query
                value: 'available'
              - name: Authorization
                in: header
                value: $steps.loginUser.outputs.sessionToken
          successCriteria:
              - condition: $statusCode == 200
          outputs:
              # outputs from this step
              availablePets: $response.body
    outputs:
      available: $steps.getPetStep.outputs.availablePets
`
	spec := v1.Spec{
		Arazzo: "1.0.0",
		Info: v1.Info{
			Title: "A pet purchasing workflow",
			Summary: strPtr(
				"This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
			),
			Description: strPtr(
				"This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
			),
			Version: "1.0.1",
		},
		SourcesDescriptions: []v1.SourceDescription{
			{
				Name: "petStoreDescription",
				Url:  "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
				Type: v1.SourceDescriptionTypeOpenAPI.ToPtr(),
			},
		},
		Workflows: []v1.Workflow{
			{
				WorkflowId: "loginUserAndRetrievePet",
				Summary: strPtr(
					"Login User and then retrieve pets",
				),
				Description: strPtr(
					"This workflow lays out the steps to login a user and then retrieve pets",
				),
				Inputs: map[string]interface{}{
					"type": "object",
					"properties": map[string]any{
						"username": map[string]any{"type": "string"},
						"password": map[string]any{"type": "string"},
					},
				},
				Steps: []v1.Step{
					{
						StepId: "loginStep",
						Description: strPtr(
							"This step demonstrates the user login step",
						),
						OperationId: strPtr("loginUser"),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "username",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.username",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "password",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.password",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"tokenExpires": "$response.header.X-Expires-After",
							"rateLimit":    "$response.header.X-Rate-Limit",
							"sessionToken": "$response.body",
						},
					},
					{
						StepId: "getPetStep",
						Description: strPtr(
							"Retrieve a pet by status from the GET pets endpoint",
						),
						OperationPath: strPtr(
							"{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
						),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "status",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "available",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "Authorization",
									In:    v1.ParameterLocationHeader.ToPtr(),
									Value: "$steps.loginUser.outputs.sessionToken",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"availablePets": "$response.body",
						},
					},
				},
				Outputs: map[string]any{
					"available": "$steps.getPetStep.outputs.availablePets",
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(spec)
	if err != nil {
		t.Fatalf("could not marshal the Spec struct: %v", err)
	}

	equal, err := yamlEqual(expectedYAMLData, string(yamlData))
	if err != nil {
		t.Fatalf("could not compare YAML strings: %v", err)
	}
	if !equal {
		t.Fatalf("expected and actual YAML strings are not equal")
	}
}

func TestSpec_UnmarshalJSON(t *testing.T) {
	jsonSpecData := `{
    "arazzo": "1.0.0",
    "info": {
        "title": "A pet purchasing workflow",
        "summary": "This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
        "description": "This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
        "version": "1.0.1"
    },
    "sourceDescriptions": [
        {
            "name": "petStoreDescription",
            "url": "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
            "type": "openapi"
        }
    ],
    "workflows": [
        {
            "workflowId": "loginUserAndRetrievePet",
            "summary": "Login User and then retrieve pets",
            "description": "This workflow lays out the steps to login a user and then retrieve pets",
            "inputs": {
                "type": "object",
                "properties": {
                    "username": {
                        "type": "string"
                    },
                    "password": {
                        "type": "string"
                    }
                }
            },
            "steps": [
                {
                    "stepId": "loginStep",
                    "description": "This step demonstrates the user login step",
                    "operationId": "loginUser",
                    "parameters": [
                        {
                            "name": "username",
                            "in": "query",
                            "value": "$inputs.username"
                        },
                        {
                            "name": "password",
                            "in": "query",
                            "value": "$inputs.password"
                        }
                    ],
                    "successCriteria": [
                        {
                            "condition": "$statusCode == 200"
                        }
                    ],
                    "outputs": {
                        "tokenExpires": "$response.header.X-Expires-After",
                        "rateLimit": "$response.header.X-Rate-Limit",
                        "sessionToken": "$response.body"
                    }
                },
                {
                    "stepId": "getPetStep",
                    "description": "Retrieve a pet by status from the GET pets endpoint",
                    "operationPath": "{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
                    "parameters": [
                        {
                            "name": "status",
                            "in": "query",
                            "value": "available"
                        },
                        {
                            "name": "Authorization",
                            "in": "header",
                            "value": "$steps.loginUser.outputs.sessionToken"
                        }
                    ],
                    "successCriteria": [
                        {
                            "condition": "$statusCode == 200"
                        }
                    ],
                    "outputs": {
                        "availablePets": "$response.body"
                    }
                }
            ],
            "outputs": {
                "available": "$steps.getPetStep.outputs.availablePets"
            }
        }
    ]
}`

	expectedSpec := v1.Spec{
		Arazzo: "1.0.0",
		Info: v1.Info{
			Title: "A pet purchasing workflow",
			Summary: strPtr(
				"This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
			),
			Description: strPtr(
				"This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
			),
			Version: "1.0.1",
		},
		SourcesDescriptions: []v1.SourceDescription{
			{
				Name: "petStoreDescription",
				Url:  "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
				Type: v1.SourceDescriptionTypeOpenAPI.ToPtr(),
			},
		},
		Workflows: []v1.Workflow{
			{
				WorkflowId: "loginUserAndRetrievePet",
				Summary: strPtr(
					"Login User and then retrieve pets",
				),
				Description: strPtr(
					"This workflow lays out the steps to login a user and then retrieve pets",
				),
				Inputs: map[string]interface{}{
					"type": "object",
					"properties": map[string]any{
						"username": map[string]any{"type": "string"},
						"password": map[string]any{"type": "string"},
					},
				},
				Steps: []v1.Step{
					{
						StepId: "loginStep",
						Description: strPtr(
							"This step demonstrates the user login step",
						),
						OperationId: strPtr("loginUser"),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "username",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.username",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "password",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.password",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"tokenExpires": "$response.header.X-Expires-After",
							"rateLimit":    "$response.header.X-Rate-Limit",
							"sessionToken": "$response.body",
						},
					},
					{
						StepId: "getPetStep",
						Description: strPtr(
							"Retrieve a pet by status from the GET pets endpoint",
						),
						OperationPath: strPtr(
							"{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
						),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "status",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "available",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "Authorization",
									In:    v1.ParameterLocationHeader.ToPtr(),
									Value: "$steps.loginUser.outputs.sessionToken",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"availablePets": "$response.body",
						},
					},
				},
				Outputs: map[string]any{
					"available": "$steps.getPetStep.outputs.availablePets",
				},
			},
		},
	}

	var spec v1.Spec

	// Unmarshal JSON data into the Spec struct
	if err := json.Unmarshal([]byte(jsonSpecData), &spec); err != nil {
		t.Fatalf("could not unmarshal the test's data: %v", err)
	}

	// Compare the expected and actual Spec structs
	if diff := deep.Equal(spec, expectedSpec); diff != nil {
		t.Fatalf(
			"expected and actual Spec structs are not equal: %v",
			diff,
		)
	}
}

func TestSpec_MarshalJSON(t *testing.T) {
	expectedJSONData := `{
    "arazzo": "1.0.0",
    "info": {
        "title": "A pet purchasing workflow",
        "summary": "This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
        "description": "This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
        "version": "1.0.1"
    },
    "sourceDescriptions": [
        {
            "name": "petStoreDescription",
            "url": "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
            "type": "openapi"
        }
    ],
    "workflows": [
        {
            "workflowId": "loginUserAndRetrievePet",
            "summary": "Login User and then retrieve pets",
            "description": "This workflow lays out the steps to login a user and then retrieve pets",
            "inputs": {
                "type": "object",
                "properties": {
                    "username": {
                        "type": "string"
                    },
                    "password": {
                        "type": "string"
                    }
                }
            },
            "steps": [
                {
                    "stepId": "loginStep",
                    "description": "This step demonstrates the user login step",
                    "operationId": "loginUser",
                    "parameters": [
                        {
                            "name": "username",
                            "in": "query",
                            "value": "$inputs.username"
                        },
                        {
                            "name": "password",
                            "in": "query",
                            "value": "$inputs.password"
                        }
                    ],
                    "successCriteria": [
                        {
                            "condition": "$statusCode == 200"
                        }
                    ],
                    "outputs": {
                        "tokenExpires": "$response.header.X-Expires-After",
                        "rateLimit": "$response.header.X-Rate-Limit",
                        "sessionToken": "$response.body"
                    }
                },
                {
                    "stepId": "getPetStep",
                    "description": "Retrieve a pet by status from the GET pets endpoint",
                    "operationPath": "{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
                    "parameters": [
                        {
                            "name": "status",
                            "in": "query",
                            "value": "available"
                        },
                        {
                            "name": "Authorization",
                            "in": "header",
                            "value": "$steps.loginUser.outputs.sessionToken"
                        }
                    ],
                    "successCriteria": [
                        {
                            "condition": "$statusCode == 200"
                        }
                    ],
                    "outputs": {
                        "availablePets": "$response.body"
                    }
                }
            ],
            "outputs": {
                "available": "$steps.getPetStep.outputs.availablePets"
            }
        }
    ]
}`
	spec := v1.Spec{
		Arazzo: "1.0.0",
		Info: v1.Info{
			Title: "A pet purchasing workflow",
			Summary: strPtr(
				"This Arazzo Description showcases the workflow for how to purchase a pet through a sequence of API calls.",
			),
			Description: strPtr(
				"This Arazzo Description walks you through the workflow and steps of searching for, selecting, and purchasing an available pet.",
			),
			Version: "1.0.1",
		},
		SourcesDescriptions: []v1.SourceDescription{
			{
				Name: "petStoreDescription",
				Url:  "https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml",
				Type: v1.SourceDescriptionTypeOpenAPI.ToPtr(),
			},
		},
		Workflows: []v1.Workflow{
			{
				WorkflowId: "loginUserAndRetrievePet",
				Summary: strPtr(
					"Login User and then retrieve pets",
				),
				Description: strPtr(
					"This workflow lays out the steps to login a user and then retrieve pets",
				),
				Inputs: map[string]interface{}{
					"type": "object",
					"properties": map[string]any{
						"username": map[string]any{"type": "string"},
						"password": map[string]any{"type": "string"},
					},
				},
				Steps: []v1.Step{
					{
						StepId: "loginStep",
						Description: strPtr(
							"This step demonstrates the user login step",
						),
						OperationId: strPtr("loginUser"),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "username",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.username",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "password",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "$inputs.password",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"tokenExpires": "$response.header.X-Expires-After",
							"rateLimit":    "$response.header.X-Rate-Limit",
							"sessionToken": "$response.body",
						},
					},
					{
						StepId: "getPetStep",
						Description: strPtr(
							"Retrieve a pet by status from the GET pets endpoint",
						),
						OperationPath: strPtr(
							"{$sourceDescriptions.petstoreDescription.url}#/paths/~1pet~1findByStatus/get",
						),
						Parameters: []v1.ParameterOrReusable{
							{
								Parameter: &v1.Parameter{
									Name:  "status",
									In:    v1.ParameterLocationQuery.ToPtr(),
									Value: "available",
								},
							},
							{
								Parameter: &v1.Parameter{
									Name:  "Authorization",
									In:    v1.ParameterLocationHeader.ToPtr(),
									Value: "$steps.loginUser.outputs.sessionToken",
								},
							},
						},
						SuccessCriteria: []v1.Criterion{
							{Condition: "$statusCode == 200"},
						},
						Outputs: map[string]any{
							"availablePets": "$response.body",
						},
					},
				},
				Outputs: map[string]any{
					"available": "$steps.getPetStep.outputs.availablePets",
				},
			},
		},
	}

	jsonData, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("could not marshal the Spec struct: %v", err)
	}

	equal, err := jsonEqual(expectedJSONData, string(jsonData))
	if err != nil {
		t.Fatalf("could not compare JSON strings: %v", err)
	}
	if !equal {
		t.Fatalf("expected and actual JSON strings are not equal")
	}
}
