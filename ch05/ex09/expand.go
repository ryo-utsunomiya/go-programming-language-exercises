package main

import "regexp"

var expandRegexp = regexp.MustCompile(`\$\w*`)

func expand(s string, f func(string) string) string {
	return expandRegexp.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}
