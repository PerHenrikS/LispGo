package lexer

import (
	"testing"
)

type tokenTest struct {
	expectedType    TokenType
	expectedLiteral string
}

func TestValidProg(t *testing.T) {
	program := `
	;This should be a comment	
	(defn x 5)
	;-.-
	(defun hello (x) (x + 1)) ;Another comment 
	(hello x) 
	;COMMENTS EVERYWHERE 
	`

	tests := []tokenTest{
		{LPAREN, "("},
		{DEFINE, "defn"},
		{IDENT, "x"},
		{NUMBER, "5"},
		{RPAREN, ")"},
		{LPAREN, "("},
		{DEFUN, "defun"},
		{IDENT, "hello"},
		{LPAREN, "("},
		{IDENT, "x"},
		{RPAREN, ")"},
		{LPAREN, "("},
		{IDENT, "x"},
		{PLUS, "+"},
		{NUMBER, "1"},
		{RPAREN, ")"},
		{RPAREN, ")"},
		{LPAREN, "("},
		{IDENT, "hello"},
		{IDENT, "x"},
		{RPAREN, ")"},
	}

	l := New(program)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - TokenLiteral wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `+-*/()'()(defn x 255 cons)`

	tests := []tokenTest{
		{PLUS, "+"},
		{MINUS, "-"},
		{MUL, "*"},
		{DIV, "/"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{QUOTE, "'"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LPAREN, "("},
		{DEFINE, "defn"},
		{IDENT, "x"},
		{NUMBER, "255"},
		{CONS, "cons"},
		{RPAREN, ")"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - TokenLiteral wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
