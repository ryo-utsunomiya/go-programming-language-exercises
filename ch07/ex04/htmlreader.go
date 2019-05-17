package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type stringReader struct {
	s string
	n int
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(r.s)[r.n:])
	r.n += n
	if r.n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func NewStringReader(s string) io.Reader {
	return &stringReader{s: s}
}

func main() {
	doc := `
<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 50px;
        background-color: #fff;
        border-radius: 1em;
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        body {
            background-color: #fff;
        }
        div {
            width: auto;
            margin: 0 auto;
            border-radius: 0;
            padding: 1em;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is established to be used for illustrative examples in documents. You may use this
    domain in examples without prior coordination or asking for permission.</p>
    <p><a href="http://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>

`
	node, err := html.Parse(NewStringReader(doc))
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range Visit(nil, node) {
		fmt.Println(link)
	}
}

func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = Visit(links, n.FirstChild)
	links = Visit(links, n.NextSibling)

	return links
}
