package environment

type Environment struct {
	Vars map[Symbol]Node
}

func New() *Environment {
	return &Environment{Vars: initializeFuncs()}
}

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
	}
}
