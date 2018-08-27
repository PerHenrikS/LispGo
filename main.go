package main

import (
	"alisp/environment"
	"alisp/eval"
	"alisp/lexer"
	"alisp/parser"
	"alisp/repl"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		switch arg {
		case "help":
			fmt.Println("Show help")
		case "repl":
			reader := bufio.NewReader(os.Stdin)
			repl.Start(reader)
		default:
			arg := os.Args[1]
			text, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Println(err)
			}
			en := environment.New()
			l := lexer.New(string(text))
			p := parser.New(l)
			exprs, err := p.Parse()
			if err != nil {
				fmt.Println("Parse error", err)
			}
			for _, expr := range exprs {
				eval.Eval(expr, en)
			}
		}
	} else {
		fmt.Println("Print usage")
	}
}
