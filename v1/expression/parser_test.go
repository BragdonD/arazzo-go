package expression

import (
	"testing"
)

func TestParser(t *testing.T) {
	tokens, err := tokenize("$request.header.x")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	printer := &ASTPrinter{}
	str := printer.Stringify(expr)
	t.Log(str)
}
