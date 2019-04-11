package xkcd

import (
	"fmt"
	"strings"
)

type Comic struct {
	Num        int
	Transcript string
}

func (c *Comic) Url() string {
	return fmt.Sprintf("https://xkcd.com/%d/", c.Num)
}

func (c *Comic) Contains(keyword string) bool {
	return strings.Index(c.Transcript, keyword) != -1
}
