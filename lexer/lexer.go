package lexer

import (
	"strings"
	"unicode"
)

// Token types
const (
	TokenCurlyOpen   = "{"       // Represents `{`
	TokenCurlyClose  = "}"       // Represents `}`
	TokenSquareOpen  = "["       // Represents `[`
	TokenSquareClose = "]"       // Represents `]`
	TokenColon       = ":"       // Represents `:`
	TokenComma       = ","       // Represents `,`
	TokenString      = "STRING"  // Represents strings
	TokenEOF         = "EOF"     // End of file/input
	TokenInvalid     = "INVALID" // Invalid token
	TokenNumber      = "NUMBER"  // Represents digit 0-9 including floats and exponents
	TokenBool        = "BOOL"    // Represents Bool true/false
	TokenNull        = "NULL"    // Represents Null
)

type Token struct {
	Type    string
	Literal string
}

// Lexer for tokenizing the input
type Lexer struct {
	input string
	pos   int
}

// NewLexer creates a new Lexer
func NewLexer(input string) *Lexer {
	return &Lexer{input: strings.TrimSpace(input)}
}

func (l *Lexer) NextToken() Token {
	// Skip whitespace
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}

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
	case '[':
		l.pos++
		return Token{Type: TokenSquareOpen, Literal: "["}
	case ']':
		l.pos++
		return Token{Type: TokenSquareClose, Literal: "]"}
	case ':':
		l.pos++
		return Token{Type: TokenColon, Literal: ":"}
	case ',':
		l.pos++
		return Token{Type: TokenComma, Literal: ","}
	case '"':
		l.pos++
		start := l.pos
		escaped := false
		for l.pos < len(l.input) {
			ch := l.input[l.pos]
			if escaped {
				escaped = false
				l.pos++
				continue
			}
			if ch == '\\' {
				escaped = true
				l.pos++
				continue
			}
			if ch == '"' {
				break
			}
			l.pos++
		}
		if l.pos >= len(l.input) {
			return Token{Type: TokenInvalid, Literal: "Unterminated string"}
		}
		literal := l.input[start:l.pos]
		l.pos++ // skip closing quote
		return Token{Type: TokenString, Literal: literal}

	default:
		if isAlpha(ch) {
			start := l.pos
			for l.pos < len(l.input) && isAlpha(l.input[l.pos]) {
				l.pos++
			}
			word := l.input[start:l.pos]
			switch word {
			case "true", "false":
				return Token{Type: TokenBool, Literal: word}
			case "null":
				return Token{Type: TokenNull, Literal: word}
			default:
				return Token{Type: TokenInvalid, Literal: word}
			}
		} else if isDigit(ch) || ch == '-' {
			start := l.pos

			// minus
			if l.input[l.pos] == '-' {
				l.pos++
			}

			// Integer part
			for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
				l.pos++
			}

			// Fractional part
			if l.pos < len(l.input) && l.input[l.pos] == '.' {
				l.pos++
				for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
					l.pos++
				}
			}

			// Exponent part
			if l.pos < len(l.input) && (l.input[l.pos] == 'e' || l.input[l.pos] == 'E') {
				l.pos++ // skip 'e' or 'E'

				if l.pos < len(l.input) && (l.input[l.pos] == '+' || l.input[l.pos] == '-') {
					l.pos++ // optional '+' or '-'
				}

				// Require at least one digit after e/E
				if l.pos >= len(l.input) || !isDigit(l.input[l.pos]) {
					return Token{Type: TokenInvalid, Literal: l.input[start:l.pos]}
				}
				for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
					l.pos++
				}
			}

			return Token{Type: TokenNumber, Literal: l.input[start:l.pos]}
		}

		// Unknown/invalid character
		l.pos++
		return Token{Type: TokenInvalid, Literal: string(ch)}
	}
}

// isAlpha returns true if ch is a letter (A-Z or a-z)
func isAlpha(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
