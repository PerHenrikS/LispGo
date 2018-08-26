package eval

import (
	"alisp/environment"
	"fmt"
)

type node = environment.Node
type env = environment.Environment
type number = environment.Number
type symbol = environment.Symbol
type fun = environment.Func

//Eval : checks type of expression and evaluates it
func Eval(expr node, en *env) node {
	var val node
	//Type switch to determine type of passed expression
	switch e := expr.(type) {
	case number:
		val = e
	case symbol:
		val = en.Vars[e]
	case []node:
		if len(e) == 0 {
			return "ok"
		}
		switch ex, _ := e[0].(symbol); ex {
		case "defn":
			//TODO: Function definition here somehow?
			en.Vars[e[1].(symbol)] = Eval(e[2], en)
			val = "ok"
		default:
			arguments := e[1:] //Operands of the function
			values := make([]node, len(arguments))
			for i, val := range arguments {
				values[i] = Eval(val, en)
			}
			val = apply(Eval(e[0], en), values) //Applies function to values (operands)
		}
	default:
		fmt.Println("Unknown expression type - ERROR")
	}
	return val
}

func apply(function node, args []node) node {
	var value node

	switch f := function.(type) {
	case func(...node) node:
		value = f(args...)
	//TODO: define function definition
	default:
		fmt.Println("Undefined function call")
	}
	return value
}
