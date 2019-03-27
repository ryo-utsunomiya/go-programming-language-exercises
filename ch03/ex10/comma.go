package main

import (
	"bytes"
	"fmt"
)

func Comma(s string) string {
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

	return buf.String()
}

func main() {
	fmt.Println(Comma("12345"))
}
