package models

import (
	"encoding/json"
	"errors"
)

// Parameter is a struct that represents an Arazzo specification 1.0.X
// parameter object.
//
// A parameter describes a single step parameter. A unique parameter
// is defined by the combination of a name and in fields.
type Parameter struct {
	// Required. The name of the parameter. Parameter names are case
	// sensitive.
	Name string `json:"name"`
	// The location of the parameter. Possible values are "path",
	// "query", "header", or "cookie". When the step in context
	// specifies a workflowId, then all parameters map to workflow
	// inputs. In all other scenarios (e.g., a step specifies an
	// operationId), the in field MUST be specified.
	In *ParameterLocation `json:"in,omitempty"`
	// Required. The value to pass in the parameter. The value can be
	// a constant or a Runtime Expression to be evaluated and passed
	// to the referenced operation or workflow.
	Value string `json:"value,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

// ParameterLocation is a string that represents the location of a
// parameter.
type ParameterLocation string

const (
	// ParameterLocationPath represents a path parameter. Used
	// together with OpenAPI style Path Templating, where the
	// parameter value is actually part of the operation’s URL. This
	// does not include the host or base path of the API. For example,
	// in /items/{itemId}, the path parameter is itemId.
	ParameterLocationPath ParameterLocation = "path"
	// ParameterLocationQuery represents a query parameter. Parameters
	// that are appended to the URL. For example, in /items?id=###,
	// the query parameter is id.
	ParameterLocationQuery ParameterLocation = "query"
	// ParameterLocationHeader represents a header parameter. Custom
	// headers that are expected as part of the request. Note that
	// [RFC9110] Name field names states field names (which includes
	// header) are case-insensitive.
	//
	// [RFC9110]: https://tools.ietf.org/html/rfc9110
	ParameterLocationHeader ParameterLocation = "header"
	// ParameterLocationCookie represents a cookie parameter. Used to
	// pass a specific cookie value to the source API.
	ParameterLocationCookie ParameterLocation = "cookie"
)

func (p ParameterLocation) ToPtr() *ParameterLocation {
	return &p
}

// ParameterOrReusable allows a step to use either a [Parameter] or
// a [Reusable] object.
type ParameterOrReusable struct {
	Parameter *Parameter `json:",omitempty"`
	Reusable  *Reusable  `json:",omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (p *ParameterOrReusable) UnmarshalJSON(data []byte) error {
	var param Parameter
	if err := json.Unmarshal(data, &param); err == nil {
		p.Parameter = &param
		return nil
	}

	var reusable Reusable
	if err := json.Unmarshal(data, &reusable); err == nil {
		p.Reusable = &reusable
		return nil
	}
	return errors.New(
		"data does not match any of the allowed types (Parameter, Reusable)",
	)
}

// MarshalJSON implements json.Marshaler interface.
func (p ParameterOrReusable) MarshalJSON() ([]byte, error) {
	if p.Parameter != nil {
		return json.Marshal(p.Parameter)
	}
	if p.Reusable != nil {
		return json.Marshal(p.Reusable)
	}
	return nil, errors.New("no data to marshal")
}

func (pr *ParameterOrReusable) ToParameter(components *Components) (*Parameter, error) {
	if pr.Parameter != nil {
		return pr.Parameter, nil
	}

	if pr.Reusable == nil {
		return nil, errors.New("no data to marshal")
	}

	return pr.Reusable.ToParameter(components)
}
