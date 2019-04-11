package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const MilestonesURL = "https://api.github.com/repos/golang/go/milestones"

type Milestone struct {
	Title       string
	State       string
	Description string
}

type Milestones []*Milestone

func FetchMilestones() (*Milestones, error) {
	resp, err := http.Get(MilestonesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result Milestones
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
