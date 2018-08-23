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
			repl.Repl(reader)
		default:
			//Read program from file
			text, _ := ioutil.ReadFile(arg)
			e := environment.New()
			tokens := lexer.Tokenize(string(text))
			parsed := parser.Parse(tokens)
			for _, expr := range parsed {
				res := eval.Eval(expr, e)
				if res != "ok" {
					fmt.Println(eval.Eval(expr, e))
				}
			}
		}
	} else {
		fmt.Println("Print usage")
	}
}
