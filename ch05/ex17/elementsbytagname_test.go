package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	doc := `
<html>
<body>
<h1 id="title">Title</h1>
<a href="https://example.com">example</a>
<div>
	Google: <a href="https://google.com">click</a>
</div>
<!-- foo -->
</body>
</html>
`
	node, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		t.Error(err)
	}

	got := ElementsByTagName(node, "a")

	if len(got) != 2 {
		t.Errorf("got : %d, want: 2", len(got))
	}

	for _, n := range got {
		if n.Data != "a" {
			t.Errorf("got: %s, want: a", n.Data)
		}
	}
}
