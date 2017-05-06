package parser

import (
	"strings"
	"testing"

	"github.com/mynz/alisp/lexer"
)

func TestParser(t *testing.T) {
	r := strings.NewReader("(print 1)")
	parser := New(lexer.New(r))

}
