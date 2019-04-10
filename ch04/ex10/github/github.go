package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesUrl = "https://api.github.com/search/issues"

var aMonthAgo, aYearAgo time.Time

func init() {
	aMonthAgo = time.Now().AddDate(0, -1, 0)
	aYearAgo = time.Now().AddDate(-1, 0, 0)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      Issues
}

type Issues []*Issue

func (s Issues) Filter(predicate func(i *Issue) bool) Issues {
	result := make([]*Issue, 0)

	for _, issue := range s {
		if predicate(issue) {
			result = append(result, issue)
		}
	}

	return result
}

func (r *IssuesSearchResult) CreatedInAMonth() Issues {
	return r.Items.Filter(func(issue *Issue) bool {
		return issue.CreatedAt.After(aMonthAgo)
	})
}

func (r *IssuesSearchResult) CreatedInAYear() Issues {
	return r.Items.Filter(func(issue *Issue) bool {
		return issue.CreatedAt.After(aYearAgo) && issue.CreatedAt.Before(aMonthAgo)
	})
}

func (r *IssuesSearchResult) CreatedAYearAgo() []*Issue {
	return r.Items.Filter(func(issue *Issue) bool {
		return issue.CreatedAt.Before(aYearAgo)
	})
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
