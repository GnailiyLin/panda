package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal = "illegal"

	Ident = "ident"
	Int   = "int"

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Asterisk = "*"
	Slash    = "/"
	Bang     = "!"

	LT    = "<"
	GT    = ">"
	EQ    = "=="
	NotEQ = "!="

	Comma     = ","
	Semicolon = ";"

	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	Function = "function"
	Let      = "let"
	True     = "true"
	False    = "false"
	If       = "if"
	Else     = "else"
	Return   = "return"

	EOF = "EOF"
)

var keyword = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdent(ident string) Type {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return Ident
}
