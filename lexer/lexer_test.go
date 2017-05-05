package lexer

import (
	"strings"
	"testing"
)

func TestLex(t *testing.T) {
	lex := New(strings.NewReader("(print 1)"))
	lex.Scan()
	if lex.TokenText() != "(" {
		t.Fatalf("first string is expected to be ( but %s", lex.TokenText())
	}
}
