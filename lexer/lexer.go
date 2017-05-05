package lexer

import (
	// "fmt"
	"io"
	"text/scanner"
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

func (lex *Lex) Scan() {
	// fmt.Println("Scan")
	lex.Token = lex.Scanner.Scan()
}
