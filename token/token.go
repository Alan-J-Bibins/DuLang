package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INTEGER = "INTEGER"

	ASSIGN_SIGN = "="
	PLUS_SIGN = "+"
	MINUS_SIGN = "-"
	DIVISION_SIGN = "/"
	MULTIPLICATION_SIGN = "*"
	REMAINDER_SIGN = "%"

	COMMA = ","
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LCURLY = "{"
	RCURLY = "}"
	LSQUARE = "["
	RSQUARE = "]"

	FUNCTION = "FUNCTION"
	LET = "LET"
)
