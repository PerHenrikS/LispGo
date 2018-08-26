package main

import (
	"alisp/repl"
	"bufio"
	"fmt"
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
			/*
				default:
					//Read program from file
					text, _ := ioutil.ReadFile(arg)
					e := environment.New()
					tokens := lexer.Tokenize(string(text))
					parsed, err := parser.Parse(tokens)
					if err != nil {
						fmt.Println(err)
					} else {
						for _, expr := range parsed {
							res := eval.Eval(expr, e)
							if res != "ok" {
								fmt.Println(eval.Eval(expr, e))
							}
						}
					}
			*/
		}
	} else {
		fmt.Println("Print usage")
	}
}
