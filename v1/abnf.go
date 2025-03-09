package v1

const (
	ABNFExpressionURL        = "$url"
	ABNFExpressionMethod     = "$method"
	ABNFExpressionStatusCode = "$statusCode"
)

const (
	ABNFExpressionRequest  = "$request."
	ABNFExpressionResponse = "$response."
	ABNFExpressionHeader   = "header."
	ABNFExpressionQuery    = "query."
	ABNFExpressionPath     = "path."
	ABNFExpressionBody     = "body"
)

const (
	ABNFExpressionJSONPointer = "#"
)

const (
	ABNFExpressionInputs             = "$inputs."
	ABNFExpressionOutputs            = "$outputs."
	ABNFExpressionSteps              = "$steps."
	ABNFExpressionWorkflows          = "$workflows."
	ABNFExpressionSourceDescriptions = "$sourceDescriptions."
	ABNFExpressionComponents         = "$components."
)

const (
	ABNFExpressionComponentsInputs         = "$components.inputs."
	ABNFExpressionComponentsParameters     = "$components.parameters."
	ABNFExpressionComponentsSuccessActions = "$components.successActions."
	ABNFExpressionComponentsFailureActions = "$components.failureActions."
)

const (
	ABNFNameRegex       = `^[a-zA-Z0-9_.-]+$`
	ABNFEscapedRegex    = `~[01]`
	ABNFUnescapedRegex  = `[\x00-\x2E\x30-\x7D\x7F-\x{10FFFF}]`
	ABNFTokenCharsRegex = "!#$%&'*+-.^_`|~0-9a-zA-Z"
	ABNFTokenRegex      = `^[` + ABNFTokenCharsRegex + `]+$`
)
