package parser

import (
	"github.com/mynz/alisp/lexer"
)

type Parser struct {
	lex *lexer.Lex
}

func New(lex *lexer.Lex) *Parser {
	return &Parser{lex}
}
