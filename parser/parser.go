package parser

import (
	"alisp/environment"
	"fmt"
	"strconv"
)

type Node = environment.Node
type number = environment.Number
type symbol = environment.Symbol

func Parse(tokens []string) []Node {
	expressions := make([]Node, 0)
	for len(tokens) > 0 {
		expressions = append(expressions, read(&tokens))
	}
	return expressions
}

func read(tokens *[]string) Node {
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]
	switch token {
	case "(":
		L := make([]Node, 0)
		for (*tokens)[0] != ")" {
			if tok := read(tokens); tok != symbol("") {
				L = append(L, tok)
			}
		}
		*tokens = (*tokens)[1:]
		return L
	case ")":
		fmt.Println("unexpected )")
		return nil
	default:
		if f, err := strconv.ParseFloat(token, 64); err == nil {
			return number(f)
		}
		return symbol(token)
	}
}
