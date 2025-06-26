package lexer

import "strings"

const (
	TokenCurlyOpen  = "{"
	TokenCurlyClose = "}"
	TokenEOF        = "EOF"
	TokenInvalid    = "INVALID"
)

type Token struct {
	Type    string
	Literal string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: strings.TrimSpace(input)}
}

func (l *Lexer) NextToken() Token {
	// Skip whitespace
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		switch ch {
		case ' ', '\t', '\n', '\r':
			l.pos++
			continue
		case '{':
			l.pos++
			return Token{Type: TokenCurlyOpen, Literal: "{"}
		case '}':
			l.pos++
			return Token{Type: TokenCurlyClose, Literal: "}"}
		default:
			return Token{Type: TokenInvalid, Literal: string(ch)}
		}
	}
	return Token{Type: TokenEOF}
}
