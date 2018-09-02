package repl

import (
	"alisp/environment"
	"alisp/eval"
	"alisp/lexer"
	"alisp/parser"
	"alisp/utils"
	"bufio"
	"fmt"
)

type env = environment.Environment
type node = environment.Node

//if true - prints entire stack in eval
var debug = false

func read(reader *bufio.Reader, en *env, d *utils.Debugger) {
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')
	l := lexer.New(text)
	p := parser.New(l)
	exprs, err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
	for _, expr := range exprs {
		evaluated := eval.Eval(expr, en, d)
		if len(en.Errors) > 0 {
			for _, err := range en.Errors {
				fmt.Println(err)
			}
			d.PrintTrace()
			d.Clear()
			break
		}
		if debug {
			d.PrintTrace()
		}
		fmt.Println(evaluated)
		d.Clear()
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
	debugger := utils.NewDebugger()
	for {
		read(reader, env, debugger)
	}
}
