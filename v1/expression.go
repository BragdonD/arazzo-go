package v1

import (
	"fmt"
	"regexp"
	"strings"
)

type RuntimeExpressionLexerTokenType int

const (
	StepURLToken RuntimeExpressionLexerTokenType = iota
	StepMethodToken
	StepStatusCodeToken
	StepRequestToken
	StepResponseToken
	WorkflowInputsToken
	WorkflowOutputsToken
	WorkflowStepsToken
	DocumentWorkflowsToken
	DocumentSourceDescriptionsToken
	DocumentComponentsToken
	ComponentsInputsToken
	ComponentsParametersToken
	ComponentsSuccessActionsToken
	ComponentsFailureActionsToken
	HeaderToken
	QueryToken
	PathToken
	BodyToken
	JSONPointerStartToken
	JSONPointerReferenceToken
	NameToken
	Token
)

var LexerTokenValue = map[RuntimeExpressionLexerTokenType]string{
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
	DocumentComponentsToken:         ABNFExpressionComponents,
	ComponentsInputsToken:           ABNFExpressionComponentsInputs,
	ComponentsParametersToken:       ABNFExpressionComponentsParameters,
	ComponentsSuccessActionsToken:   ABNFExpressionComponentsSuccessActions,
	ComponentsFailureActionsToken:   ABNFExpressionComponentsFailureActions,
	HeaderToken:                     ABNFExpressionHeader,
	QueryToken:                      ABNFExpressionQuery,
	PathToken:                       ABNFExpressionPath,
	BodyToken:                       ABNFExpressionBody,
	JSONPointerStartToken:           ABNFExpressionJSONPointer,
}

type LexerToken struct {
	Type     RuntimeExpressionLexerTokenType
	Value    string
	Position int
}

type UnknownTokenError struct {
	Value    string
	Position int
}

func (e *UnknownTokenError) Error() string {
	return fmt.Sprintf(
		"arazzo-go: lexer: unknown character at pos: %d, its value is: %s",
		e.Position,
		e.Value,
	)
}

func tokenize(input string) ([]LexerToken, error) {
	tokens := []LexerToken{}
	position := 0

	nameRe := regexp.MustCompile(ABNFNameRegex)
	tokenRe := regexp.MustCompile(ABNFTokenRegex)
	JSONPointerRe := regexp.MustCompile(ABNFJSONPointerReferenceTokenRegex)

lexerStart:
	for position < len(input) {
		for tokenType, tokenValue := range LexerTokenValue {
			if strings.HasPrefix(input[position:], tokenValue) {
				tokens = append(tokens, LexerToken{
					Type:     tokenType,
					Value:    tokenValue,
					Position: position,
				})
				position += len(tokenValue)
				continue lexerStart
			}
		}

		name := nameRe.FindString(input[position:])
		if name != "" {
			tokens = append(tokens, LexerToken{
				Type:     NameToken,
				Value:    name,
				Position: position,
			})
			position += len(name)
			continue
		}

		token := tokenRe.FindString(input[position:])
		if token != "" {
			tokens = append(tokens, LexerToken{
				Type:     Token,
				Value:    token,
				Position: position,
			})
			position += len(token)
			continue
		}

		jsonPointer := JSONPointerRe.FindString(input[position:])
		if jsonPointer != "" {
			tokens = append(tokens, LexerToken{
				Type:     JSONPointerReferenceToken,
				Value:    jsonPointer,
				Position: position,
			})
			position += len(jsonPointer)
			continue
		}

		return nil, &UnknownTokenError{
			Value:    input[position:],
			Position: position,
		}
	}

	return tokens, nil
}

// var ExpressionTokenName = map[ExpressionTokenType]string{
// 	StepURLToken:                    "StepURLToken",
// 	StepMethodToken:                 "StepMethodToken",
// 	StepStatusCodeToken:             "StepStatusCodeToken",
// 	StepRequestToken:                "StepRequestToken",
// 	StepResponseToken:               "StepResponseToken",
// 	WorkflowInputsToken:             "WorkflowInputsToken",
// 	WorkflowOutputsToken:            "WorkflowOutputsToken",
// 	DocumentWorkflowsToken:          "DocumentWorkflowsToken",
// 	DocumentSourceDescriptionsToken: "DocumentSourceDescriptionsToken",
// 	DocumentComponentsToken:         "DocumentComponentsToken",
// 	ComponentsInputsToken:           "ComponentsInputsToken",
// 	ComponentsParametersToken:       "ComponentsParametersToken",
// 	ComponentsSuccessActionsToken:   "ComponentsSuccessActionsToken",
// 	ComponentsFailureActionsToken:   "ComponentsFailureActionsToken",
// 	HeaderToken:                     "HeaderToken",
// 	QueryToken:                      "QueryToken",
// 	PathToken:                       "PathToken",
// 	BodyToken:                       "BodyToken",
// }

// type ParserToken struct {
// 	Type  ExpressionTokenType
// 	Value string
// }
