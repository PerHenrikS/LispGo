package parser

import (
	"alisp/lexer"
	"testing"
)

func TestSimpleParse(t *testing.T) {
	program := `
	(+ 2 3)
	`

	expected := []Node{
		symbol('+'),
		number(2),
		number(3),
	}

	l := lexer.New(program)
	p := New(l)
	results, _ := p.Parse()

	for i, expr := range expected {
		slice := results[0].([]Node)
		if expr != slice[i] {
			t.Fatalf("Test [%d] - parser returned: %s, expected: %s", i, slice[i], expr)
		}
	}
}
