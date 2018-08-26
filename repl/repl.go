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
	exprs := p.Parse()
	for _, expr := range exprs {
		fmt.Println(eval.Eval(expr, en))
	}

	/*
		tokens := lexer.Tokenize(text)
		parsed, err := parser.Parse(tokens)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, expr := range parsed {
				fmt.Println(eval.Eval(expr, en))
			}
		}
	*/
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
