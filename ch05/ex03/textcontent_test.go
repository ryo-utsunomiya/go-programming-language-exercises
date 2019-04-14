package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestTextContent(t *testing.T) {
	doc := `
<html>
<head>
<style>
h1{color:red}
</style>
</head>
<body>
<h1>Hello</h1>
<script>alert('hi')</script>
</body>
</html>
`
	node, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		t.Error(err)
	}

	out := new(bytes.Buffer)
	TextContent(out, node)

	got := out.String()
	want := "Hello\n"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
