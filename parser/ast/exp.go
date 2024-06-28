package ast

import (
	"bytes"
	"interpreter/token"
	"strings"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() string { return "Identifier" }
func (i *Identifier) TokenLiteral() string   { return i.Token.Literal }
func (i *Identifier) String() string         { return i.Value }

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() string { return "ExpressionStatement" }
func (es *ExpressionStatement) TokenLiteral() string  { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() string { return "IntegerLiteral" }
func (il *IntegerLiteral) TokenLiteral() string   { return il.Token.Literal }
func (il *IntegerLiteral) String() string         { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() string { return "PrefixExpression" }
func (pe *PrefixExpression) TokenLiteral() string   { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	if pe.Right != nil {
		out.WriteString(pe.Right.String())
	}
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() string { return "InfixExpression" }
func (ie *InfixExpression) TokenLiteral() string   { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	if ie.Right != nil {
		out.WriteString(ie.Right.String())
	}
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() string { return "Boolean" }
func (b *Boolean) TokenLiteral() string   { return b.Token.Literal }
func (b *Boolean) String() string         { return b.Token.Literal }

type Else struct {
	token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (e *Else) String() string {
	var out bytes.Buffer
	out.WriteString("else ")
	if e.Condition != nil {
		out.WriteString(e.Condition.String())
	} else if e.Consequence != nil {
		out.WriteString(e.Consequence.String())
	}
	return out.String()
}
func (e *Else) expressionNode() string { return "Else" }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *Else
}

func (ie *IfExpression) expressionNode() string { return "IfExpression" }
func (ie *IfExpression) TokenLiteral() string   { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString(" " + ie.Alternative.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() string { return "FunctionLiteral" }
func (fl *FunctionLiteral) TokenLiteral() string   { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() string { return "CallExpression" }
func (ce *CallExpression) TokenLiteral() string   { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() string { return "IndexExpression" }
func (ie *IndexExpression) TokenLiteral() string   { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token // '['
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() string { return "ArrayLiteral" }
func (al *ArrayLiteral) TokenLiteral() string   { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type StringExpression struct {
	Token token.Token
	Value string
}

func (s *StringExpression) expressionNode() string { return "StringExpression" }
func (s *StringExpression) TokenLiteral() string   { return s.Token.Literal }
func (s *StringExpression) String() string {
	var out bytes.Buffer
	out.WriteString("\"")
	out.WriteString(s.Value)
	out.WriteString("\"")
	return out.String()
}
