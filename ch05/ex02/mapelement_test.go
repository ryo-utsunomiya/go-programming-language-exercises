package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestMapElementNum(t *testing.T) {
	doc := `
<html>
<body>
<a href="https://example.com">example</a>
<div>
	Google: <a href="https://google.com">click</a>
</div>
</body>
</html>
`
	node, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		t.Error(err)
	}
	result := MapElementNum(nil, node)

	if result["a"] != 2 {
		t.Errorf("a should be 2")
	}
	if result["div"] != 1 {
		t.Errorf("div should bex 1")
	}
}
