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
	tokens := lexer.Tokenize(text)
	parsed := parser.Parse(tokens)
	for _, expr := range parsed {
		fmt.Println(eval.Eval(expr, en))
	}
}

/*
a simple lisp interpreter only needs to implement eval and apply
eval(expression, environment)
apply() - function application
*/
func Repl(reader *bufio.Reader) {
	fmt.Println("Welcome to alisp repl")
	env := environment.New()
	for {
		read(reader, env)
	}
}
