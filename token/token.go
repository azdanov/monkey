package token

type Type string

// Token represents a token in the Monkey programming language.
// It contains the type of the token and its literal value.
type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Represents illegal/unknown tokens
	EOF     = "EOF"     // End of file marker

	IDENT = "IDENT" // Identifiers like variable names
	INT   = "INT"   // Integer literals

	ASSIGN   = "=" // Assignment operator
	PLUS     = "+" // Addition operator
	MINUS    = "-" // Subtraction operator
	BANG     = "!" // Bang operator
	ASTERISK = "*" // Multiplication operator
	SLASH    = "/" // Division operator

	LT  = "<"  // Less than operator
	GT  = ">"  // Greater than operator
	EQ  = "==" // Equality operator
	NEQ = "!=" // Not equal operator

	COMMA     = "," // Comma delimiter
	SEMICOLON = ";" // Semicolon delimiter

	LPAREN = "(" // Left parenthesis
	RPAREN = ")" // Right parenthesis
	LBRACE = "{" // Left brace
	RBRACE = "}" // Right brace

	FUNCTION = "FUNCTION" // Function keyword
	LET      = "LET"      // Let keyword for variable declarations
	TRUE     = "TRUE"     // Boolean true
	FALSE    = "FALSE"    // Boolean false
	IF       = "IF"       // If keyword
	ELSE     = "ELSE"     // Else keyword
	RETURN   = "RETURN"   // Return keyword
)

func LookupIdent(ident string) Type {
	switch ident {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	default:
		return IDENT
	}
}
