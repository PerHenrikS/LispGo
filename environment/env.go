package environment

//Environment : TODO: comment explain why we need a parent environment (was confusing to me atleast)
type Environment struct {
	Errors []string
	Vars   map[Symbol]Node
	Parent *Environment
}

//New : Environment constructor
func New() *Environment {
	return &Environment{Vars: initializeFuncs(), Parent: nil}
}

//Find : returns correct environment in context. ie. global or function scope variables?
func (e *Environment) Find(s Symbol) *Environment {
	if _, ok := e.Vars[s]; ok {
		return e
	} else {
		return e.Parent.Find(s)
	}
}

//TODO: add error to these for runtime error reporting instead of panics
func initializeFuncs() map[Symbol]Node {
	return map[Symbol]Node{
		"+": func(a ...Node) Node {
			res := a[0].(Number)
			for _, i := range a[1:] {
				res += i.(Number)
			}
			return res
		},
		"-": func(a ...Node) Node {
			res := a[0].(Number)
			for _, i := range a[1:] {
				res -= i.(Number)
			}
			return res
		},
		"/": func(a ...Node) Node {
			res := a[0].(Number)
			for _, i := range a[1:] {
				res /= i.(Number)
			}
			return res
		},
		"*": func(a ...Node) Node {
			res := a[0].(Number)
			for _, i := range a[1:] {
				res *= i.(Number)
			}
			return res
		},
		"<": func(a ...Node) Node {
			return a[0].(Number) < a[1].(Number)
		},
		"<=": func(a ...Node) Node {
			return a[0].(Number) <= a[1].(Number)
		},
		">": func(a ...Node) Node {
			return a[0].(Number) > a[1].(Number)
		},
		">=": func(a ...Node) Node {
			return a[0].(Number) >= a[1].(Number)
		},
		"=": func(a ...Node) Node {
			return a[0] == a[1]
		},
		"cons": func(a ...Node) Node {
			//Calling them head and tail because it makes sense to me from elixir and LYAH
			switch head := a[0]; tail := a[1].(type) {
			case []Node:
				return append([]Node{head}, tail...)
			default:
				return []Node{head, tail}
			}
		},
		"head": func(a ...Node) Node {
			//return first element of list
			return a[0].([]Node)[0]
		},
		"tail": func(a ...Node) Node {
			//return all but first of list
			return a[0].([]Node)[1:]
		},
		"list": func(a ...Node) Node {
			return a
		},
	}
}
