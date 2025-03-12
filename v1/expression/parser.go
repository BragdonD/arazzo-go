package expression

import "fmt"

// Parser represents a parser for Arazzo runtime expression.
type Parser struct {
	// List of tokens to parse.
	tokens []LexerToken
	// Current position in the token list.
	current int
}

// NewParser creates a new Parser instance.
func NewParser(tokens []LexerToken) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// Parse parses the tokens and returns an expression.
func (p *Parser) Parse() (Expr, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

// expression parses an expression based on the current token.
func (p *Parser) expression() (Expr, error) {
	if p.match(StepURLToken, StepMethodToken, StepStatusCodeToken) {
		return p.singleExpression()
	}
	if p.match(StepRequestToken, StepResponseToken) {
		value := p.previous().Value
		expr, err := p.expressionWithSource()
		if err != nil {
			return nil, err
		}
		return &ExpressionWithSourceNode{
			Value:  value,
			Source: expr,
		}, nil
	}
	if p.match(WorkflowInputsToken, WorkflowOutputsToken,
		WorkflowStepsToken, DocumentWorkflowsToken,
		DocumentComponentsToken, DocumentSourceDescriptionsToken,
		ComponentsFailureActionsToken, ComponentsInputsToken,
		ComponentsParametersToken, ComponentsSuccessActionsToken,
	) {
		value := p.previous().Value
		expr, err := p.expressionWithName()
		if err != nil {
			return nil, err
		}
		return &ExpressionWithNameNode{
			Value: value,
			Name:  *expr,
		}, nil
	}

	return nil, fmt.Errorf(
		"token at %d should be an expression token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// singleExpression parses a single expression.
func (p *Parser) singleExpression() (Expr, error) {
	if !p.isAtEnd() {
		return nil, fmt.Errorf(
			"'%s' is a token with no children node",
			p.previous().Value,
		)
	}
	return &SingleExpressionNode{
		Value: p.previous().Value,
	}, nil
}

// expressionWithSource parses an expression with a source.
func (p *Parser) expressionWithSource() (SourceNode, error) {
	if p.match(QueryToken) {
		return p.queryReference()
	}
	if p.match(PathToken) {
		return p.pathReference()
	}
	if p.match(HeaderToken) {
		return p.headerReference()
	}
	if p.match(BodyToken) {
		return p.bodyReference()
	}
	return nil, fmt.Errorf(
		"token at %d should be a source token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// expressionWithName parses an expression with a name.
func (p *Parser) expressionWithName() (*NameNode, error) {
	if p.match(NameToken, NameOrToken) {
		return p.name()
	}
	return nil, fmt.Errorf(
		"token at %d should be an expression with a name token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// headerReference parses a header reference.
func (p *Parser) headerReference() (SourceNode, error) {
	header := p.previous().Value
	if p.match(Token, NameOrToken) {
		token, err := p.token()
		if err != nil {
			return nil, err
		}
		return &HeaderReferenceNode{
			Value: header,
			Token: *token,
		}, nil
	}
	return nil, fmt.Errorf(
		"token at %d should be a header reference token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// queryReference parses a query reference.
func (p *Parser) queryReference() (SourceNode, error) {
	query := p.previous().Value
	if p.match(NameToken, NameOrToken) {
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		return &QueryReferenceNode{
			Value: query,
			Name:  *name,
		}, nil
	}
	return nil, fmt.Errorf(
		"token at %d should be a query reference token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// pathReference parses a path reference.
func (p *Parser) pathReference() (SourceNode, error) {
	path := p.previous().Value
	if p.match(NameToken, NameOrToken) {
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		return &PathReferenceNode{
			Value: path,
			Name:  *name,
		}, nil
	}
	return nil, fmt.Errorf(
		"token at %d should be a path reference token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// bodyReference parses a body reference.
func (p *Parser) bodyReference() (SourceNode, error) {
	body := p.previous().Value
	if p.match(JSONPointerStartToken) {
		jsonPointerStart := p.previous().Value
		if p.match(JSONPointerReferenceToken) {
			jsonPointer, err := p.jsonPointer()
			if err != nil {
				return nil, err
			}
			return &BodyReferenceNode{
				Value:            body,
				JSONPointerStart: jsonPointerStart,
				JSONPointer:      jsonPointer,
			}, nil
		}
	}
	return nil, fmt.Errorf(
		"token at %d should be a json pointer start token, instead it is a '%s' token",
		p.peek().Position,
		LexerTokenValue[p.peek().Type],
	)
}

// name parses a name token.
func (p *Parser) name() (*NameNode, error) {
	if !p.isAtEnd() {
		return nil, fmt.Errorf(
			"'%s' is a token with no children node",
			p.previous().Value,
		)
	}
	return &NameNode{
		Value: p.previous().Value,
	}, nil
}

// token parses a token.
func (p *Parser) token() (*TokenNode, error) {
	if !p.isAtEnd() {
		return nil, fmt.Errorf(
			"'%s' is a token with no children node",
			p.previous().Value,
		)
	}
	return &TokenNode{
		Value: p.previous().Value,
	}, nil
}

// jsonPointer parses a JSON Pointer.
func (p *Parser) jsonPointer() (*JSONPointerNode, error) {
	if !p.isAtEnd() {
		return nil, fmt.Errorf(
			"'%s' is a token with no children node",
			p.previous().Value,
		)
	}
	return &JSONPointerNode{
		Value: p.previous().Value,
	}, nil
}

// match checks if the current token matches any of the given types.
func (p *Parser) match(types ...LexerTokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

// check checks if the current token is of the given type.
func (p *Parser) check(t LexerTokenType) bool {
	return p.peek().Type == t
}

// advance advances to the next token.
func (p *Parser) advance() LexerToken {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// isAtEnd checks if the parser has reached the end of the tokens.
func (p *Parser) isAtEnd() bool {
	return p.current == len(p.tokens)
}

// peek returns the current token.
func (p *Parser) peek() LexerToken {
	return p.tokens[p.current]
}

// previous returns the previous token.
func (p *Parser) previous() LexerToken {
	return p.tokens[p.current-1]
}
