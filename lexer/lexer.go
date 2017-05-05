package lexer

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

type Lex struct {
	*scanner.Scanner
	Token rune
}

func New(r io.Reader) *Lex {
	var s scanner.Scanner
	s.Init(r)

	// Only scan characters that implements lexer itself.
	s.Mode &^= scanner.ScanChars | scanner.ScanRawStrings

	return &Lex{
		Scanner: &s,
	}
}

func (l *Lex) Next() (string, error) {

	// skip comments until EOF or newline.
	for {
		token := l.Peek()
		if token != ';' {
			break
		}
		l.Scan()
		for {
			next := l.Scanner.Next()
			if next == scanner.EOF {
				l.Token = scanner.EOF
				return "", nil
			} else if next == '\n' || next == '\r' {
				break
			}
		}
	}

	l.Scan()
	switch l.Token {
	case '#':
		switch peek := l.Peek(); peek {
		case 't', 'f':
			l.Scan()
			return fmt.Sprintf("#%s", l.TokenText()), nil
		default:
			return "", errors.New("unknown hash symbol")
		}

	case '(', ')', '\'', scanner.EOF:
		return l.TokenText(), nil
	case '\n', '\r', '\t', ' ':
		l.Scan()
		return l.Next()
	default:
		token := l.TokenText()
		for {
			r := l.Peek()
			if unicode.IsSpace(r) || strings.ContainsRune("()'.,;", r) || r == scanner.EOF {
				break
			}
			token = fmt.Sprintf("%s%c", token, r)
			l.Scanner.Next()
		}
		return token, nil
	}
}

func (l *Lex) IsTokenString() bool {
	return l.Token == scanner.String
}

func (lex *Lex) Scan() {
	// fmt.Println("Scan")
	lex.Token = lex.Scanner.Scan()
}
