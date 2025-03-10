package v1

const (
	// Represents the URL of the request.
	ABNFExpressionURL = "$url"
	// Represents the HTTP method of the request.
	ABNFExpressionMethod = "$method"
	// Represents the HTTP status code of the response.
	ABNFExpressionStatusCode = "$statusCode"
)

const (
	// Prefix for accessing request-specific data.
	ABNFExpressionRequest = "$request."
	// Prefix for accessing response-specific data.
	ABNFExpressionResponse = "$response."
	// Prefix for accessing headers in request or response.
	ABNFExpressionHeader = "header."
	// Prefix for accessing query parameters in the request.
	ABNFExpressionQuery = "query."
	// Prefix for accessing path parameters in the request.
	ABNFExpressionPath = "path."
	// Represents the body of the request or response.
	ABNFExpressionBody = "body"
)

// Prefix for JSON Pointer syntax used in body references.
const ABNFExpressionJSONPointer = "#"

const (
	// Prefix for accessing workflow input parameters.
	ABNFExpressionInputs = "$inputs."
	// Prefix for accessing workflow output parameters.
	ABNFExpressionOutputs = "$outputs."
	// Prefix for accessing outputs of specific steps within a
	// workflow.
	ABNFExpressionSteps = "$steps."
	// Prefix for accessing workflows within the Arazzo document.
	ABNFExpressionWorkflows = "$workflows."
	// Prefix for accessing source descriptions within the Arazzo
	// document.
	ABNFExpressionSourceDescriptions = "$sourceDescriptions."
	// Prefix for accessing components within the Arazzo document.
	ABNFExpressionComponents = "$components."
)

const (
	// Prefix for accessing input components within the Arazzo
	// document.
	ABNFExpressionComponentsInputs = "$components.inputs."
	// Prefix for accessing parameter components within the Arazzo
	// document.
	ABNFExpressionComponentsParameters = "$components.parameters."
	// Prefix for accessing success action components within the
	// Arazzo document.
	ABNFExpressionComponentsSuccessActions = "$components.successActions."
	// Prefix for accessing failure action components within the
	// Arazzo document.
	ABNFExpressionComponentsFailureActions = "$components.failureActions."
)

const (
	// Regular expression pattern for valid 'name' as per ABNF.
	ABNFNameRegex = `^[a-zA-Z0-9_.-]+`
	// Regular expression pattern for escaped characters in JSON
	// Pointer.
	ABNFEscapedRegex = `~[01]`
	// Regular expression pattern for unescaped characters in JSON
	// Pointer.
	ABNFUnescapedRegex = `[\x00-\x2E\x30-\x7D\x7F-\x{10FFFF}]`
	// Characters allowed in 'token' as per ABNF.
	ABNFTokenChars = "!#$%&'*+-.^_`|~0-9a-zA-Z"
	// Regular expression pattern for valid 'token' as per ABNF.
	ABNFTokenRegex = `^[` + ABNFTokenChars + `]+`
	// Regular expression pattern for JSON Pointer reference tokens.
	ABNFJSONPointerReferenceTokenRegex = "^(?:/(?:" + ABNFUnescapedRegex + "|" + ABNFEscapedRegex + ")+)+$"
)
