package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch04/ex14/github"
)

// 練習問題4.14 GitHubへの一度の問い合わせで、バグレポート、マイルストーン、
// ユーザの一覧を閲覧可能にするWebサーバ
func main() {
	issues, err := github.FetchIssues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	milestones, err := github.FetchMilestones()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		renderHTMLStart(w)
		renderIssues(w, issues)
		renderMilestones(w, milestones)
		renderHTMLEnd(w)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func renderHTMLStart(w io.Writer) {
	_, err := w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
</head>
<body>
<h1>GitHub Dashboard</h1>
`))
	if err != nil {
		fmt.Fprint(w)
	}
}

func renderHTMLEnd(w io.Writer) {
	_, err := w.Write([]byte(`
</body>
</html>
`))
	if err != nil {
		fmt.Fprint(w)
	}
}

const issuesHTML = `
<h2>Issues</h2>
{{range .}}
    <div>
        <a href="{{.HTMLURL}}">#{{.Number}}</a>
        {{.Title}}
    </div>
{{end}}
`

func renderIssues(w io.Writer, issues *github.Issues) {
	t := template.Must(template.New("issues").Parse(issuesHTML))
	if err := t.Execute(w, issues); err != nil {
		fmt.Fprint(w, err)
	}
}

const milestonesHTML = `
<h2>Milestones</h2>
{{range .}}
    <div>
        {{.Title}}:{{.State}}
    </div>
{{end}}
`

func renderMilestones(w io.Writer, milestones *github.Milestones) {
	t := template.Must(template.New("milestones").Parse(milestonesHTML))
	if err := t.Execute(w, milestones); err != nil {
		fmt.Fprint(w, err)
	}
}
