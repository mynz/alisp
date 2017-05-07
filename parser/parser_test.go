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

	expected := []types.Expression{
		types.Symbol("print"),
		types.Number(1),
	}
	exps, err := parser.Parse()
	if err != nil {
		t.Fatalf("paser failed: %s", err)
	}
	if !reflect.DeepEqual(exps, expected) {
		t.Fatalf("expressions are not expected. %v", exps)
	}
}

func TestParserRecursive(t *testing.T) {
	r := strings.NewReader("(define (square x) (* x x))")
	parser := New(lexer.New(r))
	expected := []types.Expression{
		types.Symbol("define"),
		[]types.Expression{
			types.Symbol("square"),
			types.Symbol("x"),
		},
		[]types.Expression{
			types.Symbol("*"),
			types.Symbol("x"),
			types.Symbol("x"),
		},
	}
	actual, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser failed: %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expressions is not expected. %#v", actual)
	}
}

func TestParserQuote(t *testing.T) {
	r := strings.NewReader("'(1 2)")
	parser := New(lexer.New(r))
	expected := types.NewList(types.Number(1), types.Number(2))
	actual, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser failed: %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expression is not expected. %v", actual)
	}
}
