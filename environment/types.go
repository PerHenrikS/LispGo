package environment

//Node : base for Ast node that every node will implement
type Node interface{}

type Symbol string
type Number float64

type Func struct {
	Params Node
	Body   Node
	En     *Environment
}
