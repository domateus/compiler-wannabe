package ast

import (
	"bytes"
	"interpreter/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() string { return "LetStatement" }
func (ls *LetStatement) TokenLiteral() string  { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() string { return "ReturnStatement" }
func (rs *ReturnStatement) TokenLiteral() string  { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() string { return "BlockStatement" }
func (bs *BlockStatement) TokenLiteral() string  { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type ForStatement struct {
	Token     token.Token
	Condition Node
	Before    Node
	After     Node
	Body      BlockStatement
}

func (f *ForStatement) statementNode() string { return "ForStatement" }
func (f *ForStatement) TokenLiteral() string  { return f.Token.Literal }
func (f *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	if nil != f.Before {
		out.WriteString(f.Before.String() + "; ")
	}
	out.WriteString(f.Condition.String())
	if nil != f.After {
		out.WriteString("; " + f.After.String())
	}
	out.WriteString(f.Body.String())
	return out.String()
}
