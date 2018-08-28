package utils

import (
	"alisp/environment"
	"fmt"
)

/*
DevDebug : used for debugging when developing
as it is really difficult to keep track where the
recursion goes and what values a function is
called with.

Current expression - Where it is called
*/
func DevDebug(where string, expression environment.Node) {
	fmt.Println(expression, " at ", where)
}
