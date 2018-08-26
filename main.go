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
			//TODO: default: read from file and evaluate !!
		}
	} else {
		fmt.Println("Print usage")
	}
}
