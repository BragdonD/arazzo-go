package v1

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []LexerToken
		error    bool
	}{
		{
			input: ABNFExpressionURL,
			expected: []LexerToken{
				{Type: StepURLToken, Value: ABNFExpressionURL, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionMethod,
			expected: []LexerToken{
				{Type: StepMethodToken, Value: ABNFExpressionMethod, Position: 0},
			},
			error: false,
		},
		{
			input: "any",
			expected: []LexerToken{
				{Type: NameToken, Value: "any", Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionURL + ABNFExpressionMethod,
			expected: []LexerToken{
				{Type: StepURLToken, Value: ABNFExpressionURL, Position: 0},
				{Type: StepMethodToken, Value: ABNFExpressionMethod, Position: len(ABNFExpressionURL)},
			},
			error: false,
		},
		{
			input: ABNFExpressionStatusCode,
			expected: []LexerToken{
				{Type: StepStatusCodeToken, Value: ABNFExpressionStatusCode, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionRequest,
			expected: []LexerToken{
				{Type: StepRequestToken, Value: ABNFExpressionRequest, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionResponse,
			expected: []LexerToken{
				{Type: StepResponseToken, Value: ABNFExpressionResponse, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionInputs,
			expected: []LexerToken{
				{Type: WorkflowInputsToken, Value: ABNFExpressionInputs, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionOutputs,
			expected: []LexerToken{
				{Type: WorkflowOutputsToken, Value: ABNFExpressionOutputs, Position: 0},
			},
			error: false,
		},
		{
			input: ABNFExpressionHeader + ABNFExpressionQuery,
			expected: []LexerToken{
				{Type: HeaderToken, Value: ABNFExpressionHeader, Position: 0},
				{Type: QueryToken, Value: ABNFExpressionQuery, Position: len(ABNFExpressionHeader)},
			},
			error: false,
		},
		{
			input: ABNFExpressionPath + ABNFExpressionBody,
			expected: []LexerToken{
				{Type: PathToken, Value: ABNFExpressionPath, Position: 0},
				{Type: BodyToken, Value: ABNFExpressionBody, Position: len(ABNFExpressionPath)},
			},
			error: false,
		},
		{
			input: ABNFExpressionJSONPointer,
			expected: []LexerToken{
				{Type: JSONPointerStartToken, Value: ABNFExpressionJSONPointer, Position: 0},
			},
			error: false,
		},
		{
			input: "/foo/bar",
			expected: []LexerToken{
				{Type: JSONPointerReferenceToken, Value: "/foo/bar", Position: 0},
			},
			error: false,
		},
	}

	for _, test := range tests {
		tokens, err := tokenize(test.input)
		if test.error {
			if err == nil {
				t.Errorf("Expected error for input %q, got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", test.input, err)
			}
			if len(tokens) != len(test.expected) {
				t.Errorf("Token count mismatch for input %q: expected %d, got %d", test.input, len(test.expected), len(tokens))
				continue
			}
			for i, token := range tokens {
				if token != test.expected[i] {
					t.Errorf("Token mismatch at index %d for input %q: expected %+v, got %+v", i, test.input, test.expected[i], token)
				}
			}
		}
	}
}
