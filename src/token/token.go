package token

// TokenType is the value of the token
type TokenType string

// Token contains a token type and its value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL is an invalid syntax exception
	ILLEGAL = "ILLEGAL"
	// EOF says we're at the end of the file and can stop
	EOF = "EOF"

	// IDENTIFIER is any variable or function name
	IDENTIFIER = "IDENTIFIER"
	// INT is any integer
	INT = "INT"

	// ASSIGN is the basic assignment operator
	ASSIGN = "="
	// INFASSIGN is Go's type inference operator
	INFASSIGN = ":="
	// PLUS does what you expect it to, silly billy
	PLUS = "+"

	// COMMA is a comma
	COMMA = ","
	// SEMICOLON is a semicolon
	SEMICOLON = ";"

	// LPAREN opens arguments
	LPAREN = "("
	// RPAREN closes them
	RPAREN = ")"
	// LBRACE opens functions
	LBRACE = "{"
	// RBRACE closes them
	RBRACE = "}"

	// FUNCTION is the keyword to tell our interpreter that we are defining a function
	FUNCTION = "FUNCTION"
	// LET tells the interpreter we're assigning something to a variable
	LET = "LET"
)
