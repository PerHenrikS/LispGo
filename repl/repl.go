package repl

import (
	"alisp/environment"
	"alisp/eval"
	"alisp/lexer"
	"alisp/parser"
	"bufio"
	"fmt"
)

type env = environment.Environment
type node = environment.Node

func read(reader *bufio.Reader, en *env) {
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')
	l := lexer.New(text)
	p := parser.New(l)
	exprs, err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
	for _, expr := range exprs {
		evaluated := eval.Eval(expr, en)
		if len(en.Errors) > 0 {
			for _, err := range en.Errors {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(evaluated)
	}
}

/*
a simple lisp interpreter only needs to implement eval and apply
eval(expression, environment)
apply() - function application
*/
func Start(reader *bufio.Reader) {
	fmt.Println("Welcome to alisp repl")
	env := environment.New()
	for {
		read(reader, env)
	}
}
