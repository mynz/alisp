package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/mynz/alisp/lexer"
	"github.com/mynz/alisp/types"
)

func TestParser(t *testing.T) {
	r := strings.NewReader("(print 1)")
	parser := New(lexer.New(r))

	actuals := []types.Expression{
		types.Symbol("print"),
		types.Number(1),
	}
	exps, err := parser.Parse()
	if err != nil {
		t.Fatalf("paser failed: %s", err)
	}
	if !reflect.DeepEqual(exps, actuals) {
		t.Fatalf("expressions are not expected. %v", exps)
	}

}
