package eval

import (
	"alisp/environment"
	"alisp/utils"
	"fmt"
)

type node = environment.Node
type env = environment.Environment
type number = environment.Number
type symbol = environment.Symbol
type fun = environment.Func
type vars = map[symbol]node

//Eval : checks type of expression and evaluates it
func Eval(expr node, en *env, d *utils.Debugger) node {
	var val node
	if len(en.Errors) > 0 {
		return nil
	}
	//Type switch to determine type of passed expression
	switch e := expr.(type) {
	case number:
		d.DevDebug("EVAL - NUMBER", e)
		val = e
	case symbol:
		d.DevDebug("EVAL - SYMBOL", e)
		val = en.Find(e).Vars[e]
	case []node:
		d.DevDebug("EVAL - LIST OF NODES", e)
		if len(e) == 0 {
			return "ok"
		}
		switch ex, _ := e[0].(symbol); ex {
		case "quote":
			val = e[1]
		case "print":
			d.DevDebug("EVAL - PRINT", e)
			fmt.Println(Eval(e[1], en, d))
		case "defn":
			d.DevDebug("EVAL - DEFN", e)
			/*
				val = fun{Params: params, Body: e[i+1], En: en}
			*/
			en.Vars[e[1].(symbol)] = Eval(e[2], en, d)
			val = "ok"
		case "defun":
			d.DevDebug("EVAL - DEFUN", e)
			//defun is followed by a symbol, it is therefore associated with a name
			if _, ok := e[1].(symbol); ok {
				/*
					needs to evaluate to fun{p, b, e} at the end.
					so the symbol needs to be added to the environment with a
					new call to Eval() with "(defn (e[2]) [3])"
					Super hacky
				*/
				val = Eval(namedFuncSugar(e), en, d)
			} else {
				val = fun{Params: e[1], Body: e[2], En: en}
			}
		case "if":
			d.DevDebug("EVAL - IF", e)
			if Eval(e[1], en, d).(bool) {
				val = Eval(e[2], en, d)
			} else {
				val = Eval(e[3], en, d)
			}
		default:
			d.DevDebug("EVAL - FUNCTION APPLICATION", e)
			if len(e) < 2 {
				en.Errors = append(en.Errors, "evaluating function with too little args")
				return nil
			}
			arguments := e[1:] //Operands of the function
			values := make([]node, len(arguments))
			for i, val := range arguments {
				values[i] = Eval(val, en, d)
			}
			val = apply(Eval(e[0], en, d), values, d) //Applies function to values (operands)
		}
	default:
		fmt.Println("EVAL ERROR - unknown expression type", e)
	}
	return val
}

func apply(function node, args []node, d *utils.Debugger) node {
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
		value = Eval(f.Body, en, d)
	default:
		fmt.Println("EVAL ERROR (apply) - Undefined function call ", f)
	}
	return value
}

func namedFuncSugar(e []node) []node {
	if len(e) < 4 {
		return nil
	}
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
