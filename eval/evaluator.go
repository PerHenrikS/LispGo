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
type vars = map[symbol]node

//Eval : checks type of expression and evaluates it
func Eval(expr node, en *env) node {
	var val node
	//Type switch to determine type of passed expression
	switch e := expr.(type) {
	case number:
		val = e
	case symbol:
		val = en.Find(e).Vars[e]
	case []node:
		if len(e) == 0 {
			return "ok"
		}
		switch ex, _ := e[0].(symbol); ex {
		case "print":
			fmt.Println(Eval(e[1], en))
		case "defn":
			/*
				val = fun{Params: params, Body: e[i+1], En: en}
			*/
			en.Vars[e[1].(symbol)] = Eval(e[2], en)
			val = "ok"
		case "defun":
			//defun is followed by a symbol, it is therefore associated with a name
			if _, ok := e[1].(symbol); ok {
				/*
					needs to evaluate to fun{p, b, e} at the end.
					so the symbol needs to be added to the environment with a
					new call to Eval() with "(defn (e[2]) [3])"
					Super hacky
				*/
				val = Eval(namedFuncSugar(e), en)
			} else {
				val = fun{Params: e[1], Body: e[2], En: en}
			}
		case "if":
			if Eval(e[1], en).(bool) {
				val = Eval(e[2], en)
			} else {
				val = Eval(e[3], en)
			}
		default:
			arguments := e[1:] //Operands of the function
			values := make([]node, len(arguments))
			for i, val := range arguments {
				values[i] = Eval(val, en)
			}
			val = apply(Eval(e[0], en), values) //Applies function to values (operands)
		}
	default:
		fmt.Println("EVAL ERROR - unknown expression type", e)
	}
	return val
}

func apply(function node, args []node) node {
	var value node
	switch f := function.(type) {
	case func(...node) node:
		value = f(args...)
	case fun:
		en := &env{Vars: make(vars), Parent: f.En}
		switch params := f.Params.(type) {
		case []node:
			//If many parameters. add to function scope environment
			for i, param := range params {
				en.Vars[param.(symbol)] = args[i]
			}
		default:
			//Add parameter value to function scope environment
			en.Vars[params.(symbol)] = args
		}
		value = Eval(f.Body, en)
	default:
		fmt.Println("EVAL ERROR (apply) - Undefined function call ", f)
	}
	return value
}

func namedFuncSugar(e []node) []node {
	return []node{
		symbol("defn"),
		e[1],
		[]node{
			symbol("defun"),
			e[2],
			e[3],
		},
	}
}
