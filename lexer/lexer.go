package lexer

import (
	"strings"
	"unicode"
)

const (
	TokenCurlyOpen  = "{"
	TokenCurlyClose = "}"
	TokenColon      = ":"       // Represents `:`
	TokenComma      = ","       // Represents `,`
	TokenString     = "STRING"  // Represents strings
	TokenEOF        = "EOF"     // End of file/input
	TokenInvalid    = "INVALID" // Invalid token
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
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}

	// Check if its a ending token
	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	ch := l.input[l.pos]

	switch ch {
	case '{':
		l.pos++
		return Token{Type: TokenCurlyOpen, Literal: "{"}
	case '}':
		l.pos++
		return Token{Type: TokenCurlyClose, Literal: "}"}
	case ':':
		l.pos++
		return Token{Type: TokenColon, Literal: ":"}
	case ',':
		l.pos++
		return Token{Type: TokenComma, Literal: ","}
	case '"':
		l.pos++
		start := l.pos
		for l.pos < len(l.input) && l.input[l.pos] != '"' {
			l.pos++
		}
		if l.pos >= len(l.input) {
			return Token{Type: TokenInvalid, Literal: "Unterminated string"}
		}
		literal := l.input[start:l.pos]
		l.pos++
		return Token{Type: TokenString, Literal: literal}
	default:
		return Token{Type: TokenInvalid, Literal: string(ch)}
	}
}
