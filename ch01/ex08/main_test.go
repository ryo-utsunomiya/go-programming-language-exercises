package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type MockHttpClient struct{}

func (c *MockHttpClient) Get(url string) (resp *http.Response, err error) {
	if !strings.HasPrefix(url, "http://") {
		return nil, errors.New("url must have http prefix")
	}

	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<html><html>")),
	}, nil
}

func TestFetch(t *testing.T) {
	out := new(bytes.Buffer)
	Fetch(&MockHttpClient{}, "example.com", out)

	got := out.String()
	want := "<html><html>"
	if got != want {
		t.Errorf("got\n %v\n want\n %v", got, want)
	}
}
