package parser

import (
	"alisp/environment"
	"alisp/lexer"
	"errors"
	"strconv"
)

type Node = environment.Node
type number = environment.Number
type symbol = environment.Symbol

//Parser : parses the program
type Parser struct {
	l *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token
}

//New : creates and initializes new parser given a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

//Parse : parses program and returns list of expressions
func (p *Parser) Parse() ([]Node, error) {
	expressions := make([]Node, 0)
	for p.curToken.Type != lexer.EOF {
		expression, err := p.read()
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, expression)
	}
	return expressions, nil
}

//advances parsers tokens by 1
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) read() (Node, error) {
	if p.curToken.Type == lexer.EOF {
		return nil, errors.New("Parse error - Unexpected EOF")
	}

	token := p.curToken
	p.nextToken()

	switch token.Type {
	case lexer.LPAREN:
		L := make([]Node, 0)
		for p.curToken.Type != lexer.RPAREN {
			if tok, err := p.read(); tok != symbol("") {
				if err != nil {
					return nil, err
				}
				L = append(L, tok)
			}
		}
		p.nextToken()
		return L, nil
	case lexer.RPAREN:
		return nil, errors.New("Parse error - unexpected ')'")
	default:
		if f, err := strconv.ParseFloat(token.Literal, 64); err == nil {
			return number(f), nil
		}
		return symbol(token.Literal), nil
	}
}
