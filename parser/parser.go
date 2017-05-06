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

func (p *Parser) Parse() (exps types.Expression, err error) {
	token, err := p.lex.Next()
	if err != nil {
		return nil, err
	}

	switch token {
	case "'":
		if p.lex.Peek() == '(' {
			if _, err := p.lex.Next(); err != nil {
				return nil, err
			}
			tokens, err := p.parseList()
			if err != nil {
				return nil, err
			}
			return types.NewList(tokens...), nil
		}
	case "(":
		// start s-expression. Parse as list.
		return p.parseList()
	case ")":
		return nil, errors.New("unexpected ')'")
	case "#t":
		return types.Boolean(true), nil
	case "#f":
		return types.Boolean(false), nil
	default:
		if p.lex.IsTokenString() {
			return strconv.Unquote(p.lex.TokenText())
		}
		if n, err := strconv.ParseFloat(token, 64); err == nil {
			return types.Number(n), nil
		}
		return types.Symbol(token), nil
	}
}
