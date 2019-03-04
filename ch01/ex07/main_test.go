package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type FakeClient struct{}

func (c *FakeClient) Get(url string) (resp *http.Response, err error) {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<html><html>")),
	}, nil
}

func TestFetch(t *testing.T) {
	out := new(bytes.Buffer)
	Fetch(&FakeClient{}, "https://example.com", out)

	got := out.String()
	want := "<html><html>"
	if got != want {
		t.Errorf("got\n %v\n want\n %v", got, want)
	}
}
