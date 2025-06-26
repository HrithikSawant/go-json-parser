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
	stateDone
)

func (s parserState) String() string {
	switch s {
	case stateStart:
		return "Start"
	case stateExpectKeyOrEnd:
		return "ExpectKeyOrEnd"
	case stateDone:
		return "Done"
	default:
		return "UnknownState"
	}
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{lexer: l}
}

func (p *Parser) Parse() bool {
	state := stateStart

	tok := p.lexer.NextToken()
	fmt.Printf("DEBUG: State = %-20s | Token = %-10s | Literal = %s\n", state, tok.Type, tok.Literal)

	if tok.Type != lexer.TokenCurlyOpen {
		fmt.Println("DEBUG: JSON must start with '{'")
		return false
	}

	// Entering object
	state = stateExpectKeyOrEnd

	for {
		tok := p.lexer.NextToken()
		fmt.Printf("DEBUG: State = %-20s | Token = %-10s | Literal = %s\n", state, tok.Type, tok.Literal)

		switch tok.Type {
		case lexer.TokenCurlyClose:
			if state != stateExpectKeyOrEnd {
				fmt.Println("DEBUG: Unexpected '}'")
				return false
			}
			// End of object
			state = stateDone

		case lexer.TokenEOF:
			if state == stateDone {
				return true
			}
			fmt.Println("DEBUG: Unexpected EOF")
			return false

		default:
			fmt.Println("DEBUG: Invalid content inside empty object")
			return false
		}
	}
}
