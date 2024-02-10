package lexer 

import (
	"testing"
	"pickleprat.fari.f2/token"
	"log"
)

type expectedToken struct {
	ExpectedType token.TokenType
	ExpectedLiteral string 
} 

func TestLexer(t *testing.T){
	input := `
	fari $createFunc(a int, b int){
			return a + b 
		}

	fari $main() {
		int a = 3; 
		int b = 4; 
		let $result = $createFunc(a, b); 
	}  
	` 

	lex := New(input)
	
	var tests  = []expectedToken {
		{token.FARI, "fari"}, 
		{token.IDENT, "$createFunc"},
		{token.RBRACKET, "("}, 
		{token.IDENT, "a"}, 
		{token.INT, "int"}, 
        {token.COMMA, ","}, 
		{token.IDENT, "b"},
		{token.INT, "int"}, 
		{token.LBRACKET, ")"}, 
		{token.LPARENTH, "{"}, 
		{token.RETURN, "return"}, 
		{token.IDENT, "a"}, 
		{token.PLUS, "+"}, 
		{token.IDENT, "b"}, 
		{token.RPARENTH, "}"}, 
		{token.FARI, "fari"}, 
		{token.IDENT, "$main"}, 
		{token.RBRACKET, "("},
		{token.LBRACKET, ")"}, 
		{token.LPARENTH, "{"}, 
		{token.INT, "int"}, 
		{token.IDENT, "a"}, 
		{token.ASSIGN, "="}, 
		{token.NUMERIC, "3"},
		{token.SEMICOLON, ";"}, 
		{token.INT, "int"}, 
		{token.IDENT, "b"}, 
		{token.ASSIGN, "="}, 
		{token.NUMERIC, "4"},
		{token.SEMICOLON, ";"}, 
		{token.LET, "let"}, 
		{token.IDENT, "$result"}, 
		{token.ASSIGN, "="}, 
		{token.IDENT, "$createFunc"}, 
		{token.RBRACKET, "("}, 
		{token.IDENT, "a"},
		{token.COMMA, ","}, 
		{token.IDENT, "b"}, 
		{token.LBRACKET, ")"}, 
		{token.SEMICOLON, ";"}, 
		{token.RPARENTH, "}"}, 
	}

	for i, testSample := range tests{
		recievedToken := lex.NextToken(); 

		if testSample.ExpectedType != recievedToken.Type{
			log.Fatalf("test[%d]: Expected Token type %q != recieved token type %q, positon: %d",
				i, testSample.ExpectedType, 
				recievedToken.Type, lex.getPostion(), 
			)
		}

		if testSample.ExpectedLiteral != recievedToken.Literal{
			log.Fatalf("test[%d]: Expected Token literal %q != Recieved Token Literal %q", 
				i, testSample.ExpectedLiteral, 
				recievedToken.Literal, 	
			)
		}

	}
}