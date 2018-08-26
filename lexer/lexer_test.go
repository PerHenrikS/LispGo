package lexer

import "testing"

func TestNextToken(t *testing.T) {
	input := `+-*/()'()(defn x 255)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
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
