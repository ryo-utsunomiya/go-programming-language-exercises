package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
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
	result := Visit(nil, node)

	if len(result) != 2 {
		t.Errorf("expected length: 2")
	}
}
