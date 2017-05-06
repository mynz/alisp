package parser

import (
	"io"
	"text/scanner"

	"github.com/mynz/alisp/lexer"
	"github.com/mynz/alisp/types"
)

type Parser struct {
	lex *lexer.Lex
}

func New(lex *lexer.Lex) *Parser {
	return &Parser{lex}
}

func (p *Parser) parseList() ([]types.Expression, error) {
	var list []types.Expression

	// recursive scan until ")"
	for {
		if p.lex.Peek() == scanner.EOF {
			return nil, io.EOF
		}
		if p.lex.Peek() == ')' {
			break
		}
		ex, err := p.Parse()
		if err != nil {
			return nil, err
		}
		list = append(list, ex)
	}
	// detect by Peek(), so scanner should read next rune.
	p.lex.Scan()
	return list, nil
}
