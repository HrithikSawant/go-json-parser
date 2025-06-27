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

func TestStep2(t *testing.T) {
	runParserTest(t, "SinglePair", `{"key": "value"}`, true)
	// Invalid cases
	runParserTest(t, "MissingQuotesOnKey", `{key: "value"}`, false)
	runParserTest(t, "MissingQuotesOnValue", `{"key": value}`, false)
	runParserTest(t, "MissingColon", `{"key" "value"}`, false)
	runParserTest(t, "MissingComma", `{"key1": "value1" "key2": "value2"}`, false)
	runParserTest(t, "TrailingComma", `{"key": "value",}`, false)
	runParserTest(t, "EmptyStringKey", `{"": "value"}`, true)
	runParserTest(t, "EmptyStringValue", `{"key": ""}`, true)
	runParserTest(t, "ExtraCharacters", `{"key": "value"} extra`, false)
}

func TestStep3_MixedValueTypes(t *testing.T) {
	input := `{
		"key1": true,
		"key2": false,
		"key3": null,
		"key4": "value",
		"key5": 101
	}`
	runParserTest(t, "MixedValueTypes", input, true)
}

func TestStep3_InvalidMissingValue(t *testing.T) {
	input := `{
		"key1":
	}`
	runParserTest(t, "InvalidMissingValue", input, false)
}

func TestStep3_InvalidUnexpectedComma(t *testing.T) {
	input := `{
		"key1": true,
		"key2": false,
	}`
	runParserTest(t, "InvalidUnexpectedTrailingComma", input, false)
}

func TestStep3_SingleKeyStringValue(t *testing.T) {
	runParserTest(t, "SingleKeyStringValue", `{"name": "John"}`, true)
}

func TestStep3_SingleKeyBooleanTrueAndFalse(t *testing.T) {
	runParserTest(t, "SingleKeyBooleanTrue", `{"active": true}`, true)
	runParserTest(t, "SingleKeyBooleanFalse", `{"active": false}`, true)
}

func TestStep3_SingleKeyNull(t *testing.T) {
	runParserTest(t, "SingleKeyNull", `{"deleted": null}`, true)
}

func TestStep3_SingleKeyNumber(t *testing.T) {
	runParserTest(t, "SingleKeyNumber", `{"age": 30}`, true)
}

func TestStep3_MissingColon(t *testing.T) {
	runParserTest(t, "MissingColon", `{"key" "value"}`, false)
}

func TestStep3_UnquotedKey(t *testing.T) {
	runParserTest(t, "UnquotedKey", `{key: "value"}`, false)
}

func TestStep3_UnterminatedString(t *testing.T) {
	runParserTest(t, "UnterminatedString", `{"key": "value}`, false)
}

func TestStep3_InvalidBoolean(t *testing.T) {
	runParserTest(t, "InvalidBoolean", `{"flag": tru}`, false)
}

func TestStep3_InvalidNull(t *testing.T) {
	runParserTest(t, "InvalidNull", `{"deleted": nul}`, false)
}

func TestStep3_TrailingComma(t *testing.T) {
	runParserTest(t, "TrailingComma", `{"a": 1,}`, false)
}

func TestStep3_MissingComma(t *testing.T) {
	runParserTest(t, "MissingComma", `{"a": 1 "b": 2}`, false)
}

func TestStep3_NonNumericValue(t *testing.T) {
	runParserTest(t, "NonNumericValue", `{"a": 12a}`, false)
}

func TestStep3_EmptyKey(t *testing.T) {
	runParserTest(t, "EmptyKey", `{"": 123}`, true) // technically valid in JSON
}

func TestStep3_ExtraCharactersAfterObject(t *testing.T) {
	runParserTest(t, "ExtraCharactersAfterObject", `{"a":1} trailing`, false)
}

func TestStep3_WhitespaceBetweenTokens(t *testing.T) {
	runParserTest(t, "WhitespaceBetweenTokens", `{
		"key1"   :    "value1" ,
		"key2" :    false
	}`, true)
}

func TestStep3_EmptyObjectWithWhitespace(t *testing.T) {
	runParserTest(t, "EmptyObjectWithWhitespace", ` {    } `, true)
}
