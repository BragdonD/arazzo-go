package expression

import "fmt"

// Parse parses the input string and returns an expression.
// If the input string is not a valid expression, it returns an error.
func Parse(input string) (Expr, error) {
	lexer := NewLexer(input)
	tokens, err := lexer.Tokenize()
	if err != nil {
		return nil, fmt.Errorf("failed to tokenize input: %w", err)
	}

	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse tokens: %w", err)
	}

	return expr, nil
}

// Validate checks if the input string is a valid expression.
// It does not check if the expression is contextually correct.
//
// For example the expression "$workflows.foo.inputs.username"
// is a valid expression but it may not be valid in the context of the
// Arazzo document. The workflow named "foo" may not exist. The input
// "username" may not exist.
func Validate(input string) bool {
	_, err := Parse(input)
	return err == nil
}

// Extract extracts the content enclosed within curly braces {} from
// the input string. It ensures that the braces are properly matched
// and returns an error if they are not.
func Extract(input string) (string, error) {
	openBraces := 0
	result := ""

	for _, char := range input {
		switch char {
		case '{':
			openBraces++
		case '}':
			if openBraces == 0 {
				return "", fmt.Errorf("mismatched closing brace")
			}
			openBraces--
		default:
			if openBraces > 0 {
				result += string(char)
			} else {
				return "", fmt.Errorf("unexpected character outside of braces")
			}
		}
	}

	if openBraces != 0 {
		return "", fmt.Errorf("mismatched opening brace")
	}

	return result, nil
}
