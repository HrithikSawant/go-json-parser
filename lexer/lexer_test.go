package lexer

import (
	"testing"
)

func TestNextToken_EmptyObject(t *testing.T) {
	input := `{}`
	expectedTokens := []Token{
		{Type: TokenCurlyOpen, Literal: "{"},
		{Type: TokenCurlyClose, Literal: "}"},
		{Type: TokenEOF, Literal: ""},
	}

	lex := NewLexer(input)
	for i, expected := range expectedTokens {
		tok := lex.NextToken()
		if tok.Type != expected.Type || tok.Literal != expected.Literal {
			t.Errorf("Token %d - got (%q, %q), expected (%q, %q)", i, tok.Type, tok.Literal, expected.Type, expected.Literal)
		}
	}
}
