package environment

type Environment struct {
	Vars map[Symbol]Node
}

func New() *Environment {
	return &Environment{Vars: initializeFuncs()}
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
