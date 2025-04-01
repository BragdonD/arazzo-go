package expression

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Lexer struct {
	input string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

// RuntimeExpressionLexerTokenType represents the type of tokens
// recognized by the lexer.
type LexerTokenType = int

// Token types corresponding to specific runtime expressions as per
// Arazzo ABNF syntax.
const (
	// Token corresponding to "$url" expression.
	StepURLToken LexerTokenType = iota
	// Token corresponding to "$method" expression.
	StepMethodToken
	// Token corresponding to "$statusCode" expression.
	StepStatusCodeToken
	// Token corresponding to "$request." expression.
	StepRequestToken
	// Token corresponding to "$response." expression.
	StepResponseToken
	// Token corresponding to "$inputs." expression.
	WorkflowInputsToken
	// Token corresponding to "$outputs." expression.
	WorkflowOutputsToken
	// Token corresponding to "$steps." expression.
	WorkflowStepsToken
	// Token corresponding to "$workflows." expression.
	DocumentWorkflowsToken
	// Token corresponding to "$sourceDescriptions." expression.
	DocumentSourceDescriptionsToken
	// Token corresponding to "$components.inputs." expression.
	ComponentsInputsToken
	// Token corresponding to "$components.parameters." expression.
	ComponentsParametersToken
	// Token corresponding to "$components.successActions."
	// expression.
	ComponentsSuccessActionsToken
	// Token corresponding to "$components.failureActions."
	// expression.
	ComponentsFailureActionsToken
	// Token corresponding to "$components." expression.
	DocumentComponentsToken
	// Token corresponding to "header." expression.
	HeaderToken
	// Token corresponding to "query." expression.
	QueryToken
	// Token corresponding to "path." expression.
	PathToken
	// Token corresponding to "body" expression.
	BodyToken
	// Token corresponding to "#" expression.
	JSONPointerStartToken
	// Token corresponding to JSON Pointer reference.
	JSONPointerReferenceToken
	// Token corresponding to Name token as defined in ABNF syntax.
	NameToken
	// Token corresponding to Generic token as defined in ABNF syntax.
	Token
	// NameOrToken represents a token that can be either a name or a
	// generic token. This is necessary because the regex for names
	// and tokens share some common patterns.
	NameOrToken
)

// LexerTokenValue maps each token type to its corresponding string
// representation.
var LexerTokenValue = map[LexerTokenType]string{
	StepURLToken:                    ABNFExpressionURL,
	StepMethodToken:                 ABNFExpressionMethod,
	StepStatusCodeToken:             ABNFExpressionStatusCode,
	StepRequestToken:                ABNFExpressionRequest,
	StepResponseToken:               ABNFExpressionResponse,
	WorkflowInputsToken:             ABNFExpressionInputs,
	WorkflowOutputsToken:            ABNFExpressionOutputs,
	WorkflowStepsToken:              ABNFExpressionSteps,
	DocumentWorkflowsToken:          ABNFExpressionWorkflows,
	DocumentSourceDescriptionsToken: ABNFExpressionSourceDescriptions,
	ComponentsInputsToken:           ABNFExpressionComponentsInputs,
	ComponentsParametersToken:       ABNFExpressionComponentsParameters,
	ComponentsSuccessActionsToken:   ABNFExpressionComponentsSuccessActions,
	ComponentsFailureActionsToken:   ABNFExpressionComponentsFailureActions,
	DocumentComponentsToken:         ABNFExpressionComponents,
	HeaderToken:                     ABNFExpressionHeader,
	QueryToken:                      ABNFExpressionQuery,
	PathToken:                       ABNFExpressionPath,
	BodyToken:                       ABNFExpressionBody,
	JSONPointerStartToken:           ABNFExpressionJSONPointer,
}

// LexerToken represents a token identified by the lexer, including
// its type, value, and position in the input string.
type LexerToken struct {
	Type     LexerTokenType
	Value    string
	Position int
}

// UnknownTokenError is an error type returned when the lexer
// encounters an unrecognized token.
type UnknownTokenError struct {
	Value    string
	Position int
}

// Error returns a formatted error message indicating the position and
// value of the unknown token.
func (e *UnknownTokenError) Error() string {
	return fmt.Sprintf(
		"arazzo-go: lexer: unknown character at pos: %d, its value is: %s",
		e.Position,
		e.Value,
	)
}

// tokenize splits the input string into a slice of LexerTokens based
// on the Arazzo runtime expression syntax.
func (l *Lexer) Tokenize() ([]LexerToken, error) {
	tokens := []LexerToken{}
	position := 0

	// Compile regular expressions for matching names, tokens, and
	// JSON Pointer references.
	nameRe := regexp.MustCompile(ABNFNameRegex)
	tokenRe := regexp.MustCompile(ABNFTokenRegex)
	JSONPointerRe := regexp.MustCompile(
		ABNFJSONPointerReferenceTokenRegex,
	)

	tokenTypes := make([]LexerTokenType, len(LexerTokenValue))
	for k, _ := range LexerTokenValue {
		tokenTypes[k] = k
	}
	sort.Ints(tokenTypes)

	// lexerStart is the main loop for tokenizing the input string. It
	// checks for known token prefixes and processes them accordingly.
	// If a token is matched, the loop continues to check for the next
	// token.
lexerStart:
	for position < len(l.input) {
		// Check for known token prefixes in the input string.
		for _, tokenType := range tokenTypes {
			tokenValue := LexerTokenValue[tokenType]
			if strings.HasPrefix(l.input[position:], tokenValue) {
				tokens = append(tokens, LexerToken{
					Type:     tokenType,
					Value:    tokenValue,
					Position: position,
				})
				position += len(tokenValue)
				continue lexerStart
			}
		}

		// Match 'token' as defined in the ABNF syntax.
		token := tokenRe.FindString(l.input[position:])
		if token != "" {
			name := nameRe.FindString(l.input[position:])
			if name != "" {
				tokens = append(tokens, LexerToken{
					Type:     NameOrToken,
					Value:    name,
					Position: position,
				})
				position += len(name)
				continue
			}

			tokens = append(tokens, LexerToken{
				Type:     Token,
				Value:    token,
				Position: position,
			})
			position += len(token)
			continue
		}

		// Match 'name' as defined in the ABNF syntax.
		name := nameRe.FindString(l.input[position:])
		if name != "" {
			tokens = append(tokens, LexerToken{
				Type:     NameToken,
				Value:    name,
				Position: position,
			})
			position += len(name)
			continue
		}

		// Match JSON Pointer references as defined in the ABNF
		// syntax.
		jsonPointer := JSONPointerRe.FindString(l.input[position:])
		if jsonPointer != "" {
			tokens = append(tokens, LexerToken{
				Type:     JSONPointerReferenceToken,
				Value:    jsonPointer,
				Position: position,
			})
			position += len(jsonPointer)
			continue
		}

		// Return an error if an unknown token is encountered.
		return nil, &UnknownTokenError{
			Value:    l.input[position:],
			Position: position,
		}
	}

	return tokens, nil
}
