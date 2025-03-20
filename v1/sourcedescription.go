package v1

// SourceDescription is a struct that represents an Arazzo
// specification 1.0.X source description object.
//
// A source description object describes a source description (such as
// an OpenAPI description) that will be referenced by one or more
// workflows described within an Arazzo Description.
//
// An object storing a map between named description keys and location
// URLs to the source descriptions (such as an OpenAPI description)
// this Arazzo Description SHALL apply to. Each source location string
// MUST be in the form of a URI-reference as defined by [RFC3986]
// [Section 4.1].
//
// [RFC3986]: https://tools.ietf.org/html/rfc3986
// [Section 4.1]: https://tools.ietf.org/html/rfc3986#section-4.1
type SourceDescription struct {
	// Required. A unique name for the source description.
	// SHOULD conform to the regular expression [A-Za-z0-9_\-]+.
	Name string `json:"name"`
	// Required. A URL to a source description to be used by a
	// workflow. If a relative reference is used, it MUST be in
	// the form of a URI-reference as defined by [RFC3986]
	// [Section 4.2].
	//
	// [RFC3986]: https://tools.ietf.org/html/rfc3986
	// [Section 4.2]: https://tools.ietf.org/html/rfc3986#section-4.2
	Url string `json:"url"`
	// The type of source description.
	Type *SourceDescriptionType `json:"type,omitempty"`
	// Allows extensions to the Arazzo Specification. The field name
	// MUST begin with x-, for example, x-internal-id. Field names
	// beginning x-oai-, x-oas-, and x-arazzo are reserved for uses
	// defined by the OpenAPI Initiative. The value MAY be null, a
	// primitive, an array or an object.
	Extensions map[string]any `json:"-,omitempty"`
}

// SourceDescriptionType is a string that represents the type of
// source description available in an Arazzo specification 1.0.X.
type SourceDescriptionType string

const (
	// SourceDescriptionTypeOpenAPI represents an OpenAPI
	// documentation as source description.
	SourceDescriptionTypeOpenAPI SourceDescriptionType = "openapi"
	// SourceDescriptionTypeArazzo represents an Arazzo
	// documentation as source description.
	SourceDescriptionTypeArazzo SourceDescriptionType = "arazzo"
)

func (s SourceDescriptionType) ToPtr() *SourceDescriptionType {
	return &s
}
