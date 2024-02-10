package lexer

import (
	"fmt"
	"pickleprat.fari.f2/token"
)


type Lexer struct {
	input       string
	readPostion int
	position    int
	ch          byte
}

func (lex *Lexer) readChar() {
	if lex.readPostion >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPostion]
	}

	lex.position = lex.readPostion
	lex.readPostion++
}


func New(input string) *Lexer {
	lex := &Lexer{ input: input}
	lex.readChar()
	return lex; 
}

func createNewToken(t token.TokenType, c byte) token.Token {
	tok := token.Token{
		Type: t, 
		Literal: string(c), 
	} 
	return tok  
}


func (lex *Lexer) eatWhiteSpaces(){
	for ((lex.ch == ' ') || (lex.ch == '\n') || (lex.ch == '\t') || (lex.ch == '\r')){
		fmt.Println(string(lex.ch))
		lex.readChar(); 
	}
}

func (lex *Lexer ) NextToken() token.Token {
	var tok token.Token; 

	lex.eatWhiteSpaces(); 

	switch lex.ch {

		case '-':
			tok = createNewToken(token.MINUS, lex.ch) 

		case ',':
			tok = createNewToken(token.COMMA, lex.ch) 

		case '=': 
			tok = createNewToken(token.ASSIGN, lex.ch); 
		
		case '+': 
			tok = createNewToken(token.PLUS, lex.ch)
		
		case ';':
			tok = createNewToken(token.SEMICOLON, lex.ch)

		case ':':
			tok = createNewToken(token.ASSIGN, lex.ch)
		
		case '{':
			tok = createNewToken(token.LPARENTH, lex.ch)

		case '}':
			tok = createNewToken(token.RPARENTH, lex.ch)

		case '(':
			tok = createNewToken(token.RBRACKET, lex.ch)

		case ')': 
			tok = createNewToken(token.LBRACKET, lex.ch)

		case 0: 
			tok.Type = ""
			tok.Literal = token.EOF
			
		default: 
			if isLetter(lex.ch) {
				tok.Literal = lex.readIdentifier();
				identifier, ok := token.LookUpIdentifier[tok.Literal]
				
				if ok {
					tok.Type = identifier; 

				}else {
					tok.Type = token.IDENT; 
				}

				return tok;
			
			}else if (isDigit(lex.ch)){

				tok = createNewToken(token.NUMERIC, lex.ch);  
			
			}else{
				tok = createNewToken(token.ILLEGAL, lex.ch); 
			}
	}


	lex.readChar(); 
	return tok 
}


func (lex *Lexer) readIdentifier() string {
	start := lex.position;
	for isLetter(lex.ch){
		lex.readChar(); 
	}
	return lex.input[start:lex.position]; 
}


func (lex * Lexer) getPostion() int {
	return lex.position;  
}

func isLetter(ch byte) bool {
	return ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || 
							(ch == '$') || (ch == '_') || (ch == '#')) 
}

func isDigit(ch byte) bool {
	return ((ch >= 0) || (ch <= 9))
}