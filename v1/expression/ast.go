package expression

import (
	"fmt"
	"strings"
)

// Visitor interface defines methods for visiting different types of
// nodes in the AST.
type Visitor interface {
	VisitSingleExpressionNode(*SingleExpressionNode) any
	VisitExpressionWithNameNode(*ExpressionWithNameNode) any
	VisitExpressionWithSourceNode(*ExpressionWithSourceNode) any
	VisitHeaderReferenceNode(*HeaderReferenceNode) any
	VisitQueryReferenceNode(*QueryReferenceNode) any
	VisitPathReferenceNode(*PathReferenceNode) any
	VisitBodyReferenceNode(*BodyReferenceNode) any
	VisitNameNode(*NameNode) any
	VisitTokenNode(*TokenNode) any
	VisitJSONPointerNode(*JSONPointerNode) any
}

// Expr interface defines a method for accepting a visitor.
type Expr interface {
	Accept(Visitor) any
}

// SingleExpressionNode represents a node with a single expression
// value.
type SingleExpressionNode struct {
	Value string
}

// Accept method for SingleExpressionNode to accept a visitor.
func (n *SingleExpressionNode) Accept(visitor Visitor) any {
	return visitor.VisitSingleExpressionNode(n)
}

// ExpressionWithNameNode represents a node with an expression value
// and a name.
type ExpressionWithNameNode struct {
	Value string
	Name  NameNode
}

// Accept method for ExpressionWithNameNode to accept a visitor.
func (n *ExpressionWithNameNode) Accept(visitor Visitor) any {
	return visitor.VisitExpressionWithNameNode(n)
}

// SourceNode interface defines methods for nodes that have a source
// and a child node.
type SourceNode interface {
	Node() Expr
	ChildNode() Expr
	Accept(Visitor) any
}

// ExpressionWithSourceNode represents a node with an expression value
// and a source node.
type ExpressionWithSourceNode struct {
	Value  string
	Source SourceNode
}

// Accept method for ExpressionWithSourceNode to accept a visitor.
func (n *ExpressionWithSourceNode) Accept(visitor Visitor) any {
	return visitor.VisitExpressionWithSourceNode(n)
}

// HeaderReferenceNode represents a node that references a header with
// a token.
type HeaderReferenceNode struct {
	Value string
	Token TokenNode
}

// Node method returns the current node as an Expr.
func (n *HeaderReferenceNode) Node() Expr {
	return n
}

// ChildNode method returns the child node as an Expr.
func (n *HeaderReferenceNode) ChildNode() Expr {
	return &n.Token
}

// Accept method for HeaderReferenceNode to accept a visitor.
func (n *HeaderReferenceNode) Accept(visitor Visitor) any {
	return visitor.VisitHeaderReferenceNode(n)
}

// QueryReferenceNode represents a node that references a query with a
// name.
type QueryReferenceNode struct {
	Value string
	Name  NameNode
}

// Node method returns the current node as an Expr.
func (n *QueryReferenceNode) Node() Expr {
	return n
}

// ChildNode method returns the child node as an Expr.
func (n *QueryReferenceNode) ChildNode() Expr {
	return &n.Name
}

// Accept method for QueryReferenceNode to accept a visitor.
func (n *QueryReferenceNode) Accept(visitor Visitor) any {
	return visitor.VisitQueryReferenceNode(n)
}

// PathReferenceNode represents a node that references a path with a
// name.
type PathReferenceNode struct {
	Value string
	Name  NameNode
}

// Node method returns the current node as an Expr.
func (n *PathReferenceNode) Node() Expr {
	return n
}

// ChildNode method returns the child node as an Expr.
func (n *PathReferenceNode) ChildNode() Expr {
	return &n.Name
}

// Accept method for PathReferenceNode to accept a visitor.
func (n *PathReferenceNode) Accept(visitor Visitor) any {
	return visitor.VisitPathReferenceNode(n)
}

// BodyReferenceNode represents a node that references a body with a
// JSON pointer.
type BodyReferenceNode struct {
	Value            string
	JSONPointerStart rune
	JSONPointer      *JSONPointerNode
}

// Node method returns the current node as an Expr.
func (n *BodyReferenceNode) Node() Expr {
	return n
}

// ChildNode method returns the child node as an Expr.
func (n *BodyReferenceNode) ChildNode() Expr {
	return n.JSONPointer
}

// Accept method for BodyReferenceNode to accept a visitor.
func (n *BodyReferenceNode) Accept(visitor Visitor) any {
	return visitor.VisitBodyReferenceNode(n)
}

// NameNode represents a node with a name value.
type NameNode struct {
	Value string
}

// Accept method for NameNode to accept a visitor.
func (n *NameNode) Accept(visitor Visitor) any {
	return visitor.VisitNameNode(n)
}

// TokenNode represents a node with a token value.
type TokenNode struct {
	Value string
}

// Accept method for TokenNode to accept a visitor.
func (n *TokenNode) Accept(visitor Visitor) any {
	return visitor.VisitTokenNode(n)
}

// JSONPointerNode represents a node with a JSON pointer value.
type JSONPointerNode struct {
	Value string
}

// Accept method for JSONPointerNode to accept a visitor.
func (n *JSONPointerNode) Accept(visitor Visitor) any {
	return visitor.VisitJSONPointerNode(n)
}

// ASTPrinter is a visitor that prints the AST nodes.
type ASTPrinter struct {
}

// parenthesize method formats the node and its children as a string.
func (ast *ASTPrinter) parenthesize(
	name string,
	exprs ...Expr,
) string {
	builder := strings.Builder{}

	builder.WriteRune('(')
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteRune(' ')
		str := fmt.Sprintf("%v", expr.Accept(ast))
		builder.WriteString(str)
	}
	builder.WriteRune(')')

	return builder.String()
}

// Stringify method returns the string representation of the AST.
func (ast *ASTPrinter) Stringify(expr Expr) string {
	return fmt.Sprintf("%v", expr.Accept(ast))
}

// VisitSingleExpressionNode method for visiting SingleExpressionNode.
func (ast *ASTPrinter) VisitSingleExpressionNode(
	n *SingleExpressionNode,
) any {
	return n.Value
}

// VisitExpressionWithNameNode method for visiting
// ExpressionWithNameNode.
func (ast *ASTPrinter) VisitExpressionWithNameNode(
	n *ExpressionWithNameNode,
) any {
	return ast.parenthesize(n.Value, &n.Name)
}

// VisitExpressionWithSourceNode method for visiting
// ExpressionWithSourceNode.
func (ast *ASTPrinter) VisitExpressionWithSourceNode(
	n *ExpressionWithSourceNode,
) any {
	return ast.parenthesize(n.Value, n.Source)
}

// VisitHeaderReferenceNode method for visiting HeaderReferenceNode.
func (ast *ASTPrinter) VisitHeaderReferenceNode(
	n *HeaderReferenceNode,
) any {
	return ast.parenthesize(n.Value, &n.Token)
}

// VisitQueryReferenceNode method for visiting QueryReferenceNode.
func (ast *ASTPrinter) VisitQueryReferenceNode(
	n *QueryReferenceNode,
) any {
	return ast.parenthesize(n.Value, &n.Name)
}

// VisitPathReferenceNode method for visiting PathReferenceNode.
func (ast *ASTPrinter) VisitPathReferenceNode(
	n *PathReferenceNode,
) any {
	return ast.parenthesize(n.Value, &n.Name)
}

// VisitBodyReferenceNode method for visiting BodyReferenceNode.
func (ast *ASTPrinter) VisitBodyReferenceNode(
	n *BodyReferenceNode,
) any {
	if n.JSONPointer != nil {
		return ast.parenthesize(
			n.Value+" "+string(n.JSONPointerStart),
			n.JSONPointer,
		)
	}
	return n.Value
}

// VisitNameNode method for visiting NameNode.
func (ast *ASTPrinter) VisitNameNode(n *NameNode) any {
	return n.Value
}

// VisitTokenNode method for visiting TokenNode.
func (ast *ASTPrinter) VisitTokenNode(n *TokenNode) any {
	return n.Value
}

// VisitJSONPointerNode method for visiting JSONPointerNode.
func (ast *ASTPrinter) VisitJSONPointerNode(n *JSONPointerNode) any {
	return n.Value
}
