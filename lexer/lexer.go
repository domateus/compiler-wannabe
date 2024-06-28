package lexer

import (
	"fmt"
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

type LexError struct {
	Message string
}

func (l *Lexer) noError() LexError {
	return LexError{}
}

func (l *Lexer) NextToken() (token.Token, LexError) {
	var tkn token.Token
	err := l.noError()

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tkn = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tkn = newToken(token.BANG, l.ch)
		}
	case '/':
		tkn = newToken(token.SLASH, l.ch)
	case '*':
		tkn = newToken(token.ASTERISK, l.ch)
	case '<':
		tkn = newToken(token.LT, l.ch)
	case '>':
		tkn = newToken(token.GT, l.ch)
	case '-':
		tkn = newToken(token.MINUS, l.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '{':
		tkn = newToken(token.LBRACE, l.ch)
	case '}':
		tkn = newToken(token.RBRACE, l.ch)
	case '[':
		tkn = newToken(token.LBRACKET, l.ch)
	case ']':
		tkn = newToken(token.RBRACKET, l.ch)
	case '"':
		tkn.Type = token.STRING
		tkn.Literal = l.readString()
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn, l.noError()
		} else if isDigit(l.ch) {
			tkn.Literal = l.readNumber()
			tkn.Type = token.INT
			return tkn, l.noError()
		} else {
			tkn = newToken(token.ILLEGAL, l.ch)
			err.Message = fmt.Sprintf("Erro léxico.\n\ttoken \" %s \" não é aceito na linguagem", tkn.Literal)
		}
	}
	l.readChar()
	return tkn, err
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
