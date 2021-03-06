package lexer

//TokenType : type for the token types
type TokenType string

//Token : represents the individual token ()
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENTIFIER"
	NUMBER    = "NUMBER"
	PLUS      = "+"
	MINUS     = "-"
	DIV       = "/"
	MUL       = "*"
	LPAREN    = "("
	RPAREN    = ")"
	LSQBRACK  = "["
	RSQBRACK  = "]"
	GTHEN     = ">"
	GEQUAL    = ">="
	LTHEN     = "<"
	LEQUAL    = "<="
	EQUAL     = "="
	SLCOMMENT = ";"
)

//Lexer : holds information needed for tokenizing
type Lexer struct {
	input    string //input program
	position int    //lexers current position in program
	readPos  int    //lexers "peek" position in program (position + 1)
	ch       byte   //current character
}

//New : given input program returns lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() //initialize lexer to state pos = 0 readPos = 1
	return l
}

func (l *Lexer) readChar() {
	//if end of file is reached
	if l.readPos >= len(l.input) {
		l.ch = 0 //0 to represent eof
	} else {
		l.ch = l.input[l.readPos] //set current char to next in input
	}
	l.position = l.readPos
	l.readPos++
}

//returns next character in stream without advancing lexer state
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

//NextToken : returns a token based on the current character
func (l *Lexer) NextToken() Token {
	var tok Token

	//Skips comments and whitespace
	l.skip()

	switch l.ch {
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '/':
		tok = newToken(DIV, l.ch)
	case '*':
		tok = newToken(MUL, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '<':
		if l.peekChar() == '=' {
			tok.Literal = "<="
			tok.Type = LEQUAL
			l.readChar() //Otherwise it will get an extra =
		} else {
			tok = newToken(LTHEN, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok.Literal = ">="
			tok.Type = GEQUAL
			l.readChar() //Otherwise it will get an extra =
		} else {
			tok = newToken(GTHEN, l.ch)
		}
	case '=':
		tok = newToken(EQUAL, l.ch)
	case 0: //EOF as defined in readChar()
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = IDENT
			return tok //return instantly because readChar has been executed in readIdentifier()
		} else if isDigit(l.ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			return tok ////return instantly because readChar has been executed in readNumber()
		}
		tok = newToken(ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

//as long as current char is a number increment Lexer.Position -> return slice of input from start - end
func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

//as long as current char is a letter increment Lexer.Position -> return slice of input from start - end
func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skip() {
	l.skipWhitespace()
	l.skipSingleLineComment()
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//Skips all after ; and then whitespace
func (l *Lexer) skipSingleLineComment() {
	if l.ch == ';' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
		l.skipWhitespace()
		if l.ch == ';' {
			l.skip()
		}
	}
}

//util function to create a new token
func newToken(tokenType TokenType, literal byte) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isNextLine(ch byte) bool {
	return ch == '\n'
}
