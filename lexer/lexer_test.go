package lexer

import (
	"testing"
	"github.com/metonimie/monkeyInterpreter/token"
)

func TestLexer_TestNextToken(t *testing.T) {
	input := `=+-(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("TestNextToken test[%d]: Wrong token type, expected %q got %q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken test[%d]: Wrong token literal, expected %q got %q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestLexer_NextToken2(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("NextToken2 test[%d]: Wrong token type, expected %q got %q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("NextToken2 test[%d]: Wrong token literal, expected %q got %q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestLexer_TestNextToken3(t *testing.T) {
	input := `<>!/*`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LT, "<"},
		{token.GT, ">"},
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.TIMES, "*"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("TestNextToken test[%d]: Wrong token type, expected %q got %q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken test[%d]: Wrong token literal, expected %q got %q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestLexer_TestNextToken4(t *testing.T) {
	input := `
	if (true > false) {
		return 0;
	} else return 1 < 0;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.GT, ">"},
		{token.FALSE, "false"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("TestNextToken test[%d]: Wrong token type, expected %q got %q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken test[%d]: Wrong token literal, expected %q got %q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}
func TestLexer_TestNextToken5(t *testing.T) {
	input := `
	1 != 1
	1 == 1
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "1"},
		{token.NOT_EQ, "!="},
		{token.INT, "1"},

		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("TestNextToken test[%d]: Wrong token type, expected %q got %q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken test[%d]: Wrong token literal, expected %q got %q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}
