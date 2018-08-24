package parser

import (
	"alisp/environment"
	"errors"
	"strconv"
)

type Node = environment.Node
type number = environment.Number
type symbol = environment.Symbol

func Parse(tokens []string) ([]Node, error) {
	expressions := make([]Node, 0)
	for len(tokens) > 0 {
		nodes, err := read(&tokens)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, nodes)
	}
	return expressions, nil
}

func read(tokens *[]string) (Node, error) {
	if len(*tokens) <= 0 {
		return nil, errors.New("Len too short")
	}
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]
	switch token {
	case "(":
		L := make([]Node, 0)
		for (*tokens)[0] != ")" {
			if tok, err := read(tokens); tok != symbol("") {
				if err != nil {
					return nil, err
				}
				L = append(L, tok)
			}
		}
		*tokens = (*tokens)[1:]
		return L, nil
	case ")":
		return nil, errors.New("Unexpected )")
	default:
		if f, err := strconv.ParseFloat(token, 64); err == nil {
			return number(f), nil
		}
		return symbol(token), nil
	}
}
