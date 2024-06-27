package ast

import (
	"bytes"
)

// Abstract Syntax Tree
// Every node in our AST has to implement the Node interface
type Node interface {
	TokenLiteral() string
	String() string
}
type Expression interface {
	Node
	expressionNode() string
}
type Statement interface {
	Node
	statementNode() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) root() string {
	return p.Statements[0].TokenLiteral()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.root()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
