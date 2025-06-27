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

func TestNextToken_KeyValuePair(t *testing.T) {
	input := `{"key": "value"}`
	expectedTokens := []Token{
		{Type: TokenCurlyOpen, Literal: "{"},
		{Type: TokenString, Literal: "key"},
		{Type: TokenColon, Literal: ":"},
		{Type: TokenString, Literal: "value"},
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

func TestNextToken_MultiplePairs(t *testing.T) {
	input := `{"a": "1", "b": "2"}`
	expectedTokens := []Token{
		{Type: TokenCurlyOpen, Literal: "{"},
		{Type: TokenString, Literal: "a"},
		{Type: TokenColon, Literal: ":"},
		{Type: TokenString, Literal: "1"},
		{Type: TokenComma, Literal: ","},
		{Type: TokenString, Literal: "b"},
		{Type: TokenColon, Literal: ":"},
		{Type: TokenString, Literal: "2"},
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

func TestNextToken_UnterminatedString(t *testing.T) {
	input := `{"key": "value}`
	lex := NewLexer(input)

	// Advance to the string token
	_ = lex.NextToken() // {
	_ = lex.NextToken() // "key"
	_ = lex.NextToken() // :
	tok := lex.NextToken()

	if tok.Type != TokenInvalid {
		t.Errorf("Expected INVALID token, got %q", tok.Type)
	}
}
