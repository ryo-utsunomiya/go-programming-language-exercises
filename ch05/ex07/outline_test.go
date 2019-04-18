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

	want := `<html>
  <head/>
  <body>
    <a href="https://example.com">
      example
    </a>
    <div>
      Google:
      <a href="https://google.com">
        click
      </a>
    </div>
    <!--  foo  -->
  </body>
</html>
`

	got := ForEachNode(node, startElement, endElement)
	if got != want {
		t.Error()
	}
}
