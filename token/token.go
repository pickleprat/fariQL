package token; 

type TokenType string; 

type Token struct {
	Type TokenType;  
	Literal string;  
}

var LookUpIdentifier = map[string] TokenType {
	"fari": FARI, 
	"var": VAR, 
	"const": CONST, 
	"let": LET, 
	"f2": F2, 
	"return": RETURN, 
	"int": INT, 
}

const (
	SEMICOLON = ";"
	COLON = ":"  
	FARI = "FARI"
	ASSIGN = "=" 
	VAR = "VAR"
	PLUS = "+"
	NUMERIC = "NUMERIC" 
	MINUS = "-" 
	LPARENTH = "{" 
	RPARENTH = "}" 
	CONST = "CONST"  
	LBRACKET = "("	
	RBRACKET = ")" 
	F2 = "F2"
	IDENT = "IDENT" 
	INT = "INT"
	LET = "LET" 
	RETURN = "RETURN"
	COMMA = "," 
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
) 