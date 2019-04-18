package main

import (
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
<link href="style.css">
</head>
<body>
<h1>Hello</h1>
<img src="logo.png">
<a href="/">Go to top</a>
<script>alert('hi')</script>
<script src="jquery.js">
</body>
</html>
`
	node, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		t.Error(err)
	}

	links := Visit(nil, node)

	if links[0] != "style.css" {
		t.Errorf("want: style.css, got %s", links[0])
	}
	if links[1] != "logo.png" {
		t.Errorf("want: logo.png, got %s", links[1])
	}
	if links[2] != "/" {
		t.Errorf("want: /, got %s", links[2])
	}
	if links[3] != "jquery.js" {
		t.Errorf("want: jquery.js, got %s", links[3])
	}
}
