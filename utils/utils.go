package utils

import (
	"alisp/environment"
	"fmt"
)

//Trace : to store trace information for format print
type Trace struct {
	location string
	expr     environment.Node
	index    int
}

//Debugger : stores information to help debugging
type Debugger struct {
	Trace []Trace
	Depth int
}

//NewDebugger : constructor
func NewDebugger() *Debugger {
	return &Debugger{Trace: []Trace{}, Depth: 0}
}

/*
DevDebug : used for debugging when developing
as it is really difficult to keep track where the
recursion goes and what values a function is
called with.

Current expression - Where it is called
*/
func (d *Debugger) DevDebug(where string, expression environment.Node) {
	t := Trace{location: where, expr: expression}
	t.index = d.Depth
	d.Trace = append(d.Trace, t)
	d.Depth++
}

//PrintTrace : prints the entire stack
func (d *Debugger) PrintTrace() {
	for _, trace := range d.Trace {
		fmt.Println("Trace:", trace.index, trace.location, " at ", trace.expr)
	}
}

//Clear : clears stack for next trace
func (d *Debugger) Clear() {
	d.Trace = d.Trace[:0]
	d.Depth = 0
}
