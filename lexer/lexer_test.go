package lexer

import (
	"strings"
	"testing"
)

func TestLex(t *testing.T) {
	lex := New(strings.NewReader("(print 1)"))
	lex.Scan()
	if lex.TokenText() != "(" {
		t.Fatalf("first string is expected to be '(' but '%s'.", lex.TokenText())
	}
	lex.Scan()
	if lex.TokenText() != "print" {
		t.Fatalf("first string is expected to be 'print' but '%s'.", lex.TokenText())
	}
}
