package lexer

import (
	"testing"

	"github.com/tekihei2317/go-interpreter/token"
)

type TestCase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runTests(l *Lexer, t *testing.T, tests []TestCase) {
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []TestCase{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	runTests(l, t, tests)
}

func TestNextToken2(t *testing.T) {
	input := `let five = 13;`

	tests := []TestCase{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "13"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	runTests(l, t, tests)
}
