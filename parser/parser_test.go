package parser

import (
	"testing"

	"github.com/HrithikSawant/go-json-parser/lexer"
)

func runParserTest(t *testing.T, name string, input string, expectValid bool) {
	t.Run(name, func(t *testing.T) {
		lex := lexer.NewLexer(input)
		p := NewParser(lex)
		valid := p.Parse()

		if valid != expectValid {
			status := "valid"
			if !expectValid {
				status = "invalid"
			}
			t.Errorf("Test %s failed. Expected input to be %s, but got %v", name, status, valid)
		}
	})
}

func TestStep1(t *testing.T) {
	runParserTest(t, "EmptyObject", `{}`, true)
	runParserTest(t, "OnlyOpeningBrace", `{`, false)
	runParserTest(t, "OnlyClosingBrace", `}`, false)
	runParserTest(t, "ExtraComma", `{,}`, false)
	runParserTest(t, "TrailingCharacters", `{} extra`, false)
	runParserTest(t, "JustString", `"key"`, false)
	runParserTest(t, "EmptyInput", ``, false)
}
