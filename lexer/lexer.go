package lexer

import (
	"strings"
)

//Tokenize : input -> tokens
func Tokenize(input string) []string {
	var output string
	output = strings.Replace(input, "(", " ( ", -1)
	output = strings.Replace(output, ")", " ) ", -1)

	return strings.Fields(output)
}
