package parser

import (
	"fmt"

	"github.com/HrithikSawant/go-json-parser/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

type parserState int

const (
	stateStart parserState = iota
	stateExpectKeyOrEnd
	stateExpectColon
	stateExpectValue
	stateExpectCommaOrEnd
	stateDone
)

func (s parserState) String() string {
	switch s {
	case stateStart:
		return "Start"
	case stateExpectKeyOrEnd:
		return "ExpectKeyOrEnd"
	case stateExpectColon:
		return "ExpectColon"
	case stateExpectValue:
		return "ExpectValue"
	case stateExpectCommaOrEnd:
		return "ExpectCommaOrEnd"
	case stateDone:
		return "Done"
	default:
		return "UnknownState"
	}
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{lexer: l}
}

// Parse starts parsing from the top-level JSON object.
func (p *Parser) Parse() bool {
	state := stateStart

	tok := p.lexer.NextToken()
	fmt.Printf("DEBUG: State = %-20s | Token = %-10s | Literal = %s\n", state, tok.Type, tok.Literal)

	if tok.Type != lexer.TokenCurlyOpen {
		fmt.Println("DEBUG: JSON must start with '{'")
		return false
	}

	if !p.parseObject() {
		return false
	}

	state = stateDone
	tok = p.lexer.NextToken()
	fmt.Printf("DEBUG: State = %-20s | Token = %-10s | Literal = %s\n", state, tok.Type, tok.Literal)

	if tok.Type != lexer.TokenEOF {
		fmt.Println("DEBUG: Extra tokens after end of object")
		return false
	}

	return true
}

// parseObject parses a JSON object and any nested objects recursively.
func (p *Parser) parseObject() bool {
	state := stateExpectKeyOrEnd
	justSawComma := false

	for {
		tok := p.lexer.NextToken()
		fmt.Printf("DEBUG: State = %-20s | Token = %-10s | Literal = %s\n", state, tok.Type, tok.Literal)

		switch tok.Type {
		case lexer.TokenCurlyClose:
			if state == stateExpectKeyOrEnd && justSawComma {
				fmt.Println("DEBUG: Trailing comma before '}' is not allowed")
				return false
			}
			if state != stateExpectKeyOrEnd && state != stateExpectCommaOrEnd {
				fmt.Println("DEBUG: Unexpected '}' in state", state)
				return false
			}
			return true

		case lexer.TokenColon:
			if state != stateExpectColon {
				fmt.Println("DEBUG: Unexpected ':' — expected key first")
				return false
			}
			state = stateExpectValue

		case lexer.TokenString, lexer.TokenNumber, lexer.TokenBool, lexer.TokenNull:
			switch state {
			case stateExpectKeyOrEnd:
				if tok.Type != lexer.TokenString {
					fmt.Printf("DEBUG: Object key must be STRING but got %s\n", tok.Type)
					return false
				}
				state = stateExpectColon

			case stateExpectValue:
				state = stateExpectCommaOrEnd

			default:
				fmt.Printf("DEBUG: Unexpected value in state %s\n", state)
				return false
			}

		case lexer.TokenCurlyOpen:
			if state != stateExpectValue {
				fmt.Printf("DEBUG: Unexpected '{' in state %s\n", state)
				return false
			}
			if !p.parseObject() {
				return false
			}
			state = stateExpectCommaOrEnd

		case lexer.TokenComma:
			if state != stateExpectCommaOrEnd {
				fmt.Printf("DEBUG: Unexpected comma in state %s\n", state)
				return false
			}
			state = stateExpectKeyOrEnd
			justSawComma = true

		case lexer.TokenInvalid:
			fmt.Println("DEBUG: Invalid token encountered")
			return false

		case lexer.TokenEOF:
			fmt.Println("DEBUG: Unexpected end of input")
			return false

		default:
			fmt.Printf("DEBUG: Unknown token type: %s\n", tok.Type)
			return false
		}
	}
}
