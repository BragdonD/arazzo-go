package expression

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"$method", "$method", false},
		{"$request.header.accept", "($request. (header. accept))", false},
		{"$request.path.id", "($request. (path. id))", false},
		{"$request.body#/user/uuid", "($request. (body # /user/uuid))", false},
		{"$url", "$url", false},
		{"$response.body#/status", "($response. (body # /status))", false},
		{"$response.header.Server", "($response. (header. Server))", false},
		{"$inputs.username", "($inputs. username)", false},
		{"$workflows.foo.inputs.username", "($workflows. foo.inputs.username)", false},
		{"$steps.someStepId.outputs.pets", "($steps. someStepId.outputs.pets)", false},
		{"$outputs.bar", "($outputs. bar)", false},
		{"$workflows.foo.outputs.bar", "($workflows. foo.outputs.bar)", false},
		{"$components.parameters.foo", "($components.parameters. foo)", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			lexer := NewLexer(test.input)
			tokens, err := lexer.Tokenize()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			parser := NewParser(tokens)
			expr, err := parser.Parse()

			if (err != nil) != test.err {
				t.Fatalf("expected error: %v, got: %v", test.err, err)
			}

			if err != nil {
				return
			}

			printer := &ASTPrinter{}
			str := printer.Stringify(expr)
			if str != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, str)
			}
		})
	}
}
