package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Comma(s string) string {
	prefix := ""
	if strings.HasPrefix(s, "-") {
		prefix = "-"
		s = s[1:]
	}

	postfix := ""
	if strings.Contains(s, ".") {
		i := strings.Index(s, ".")
		postfix = s[i:]
		s = s[:i]
	}

	n := len(s)

	pre := 3
	if n%3 != 0 {
		pre = n % 3
	}

	var buf bytes.Buffer
	buf.WriteString(s[:pre])

	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return prefix + buf.String() + postfix
}

func main() {
	fmt.Println(Comma("-1234.567"))
}
