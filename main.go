package main

// Reference: https://roadmap.sh/projects/github-user-activity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Activity struct {
	Type      string `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		RefType string `json:"ref_type"`
		Commit  []struct {
			Message string `json:"message"`
		} `json:"commit"`
	} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		println("Usage: go run main.go <name>")
		return
	}

	endpoint := "https://api.github.com/users/" + args[0] + "/events"
	response, err := http.Get(endpoint)
	if err != nil {
		println("Error fetching GitHub activity:", err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			println("User not found:", args[0])
		} else {
			println("Error fetching GitHub activity: HTTP status", response.StatusCode)
		}
		return
	}

	var acrvities []Activity
	err = json.NewDecoder(response.Body).Decode(&acrvities)
	if err != nil {
		println("Error decoding GitHub activity:", err.Error())
		return
	}

	for _, activity := range acrvities {
		switch activity.Type {
		case "PushEvent":
			fmt.Printf("Pushed to %s at %s\n", activity.Repo.Name, activity.CreatedAt)
			for _, commit := range activity.Payload.Commit {
				fmt.Printf("  - %s\n", commit.Message)
			}
		case "PullRequestEvent":
			fmt.Printf("Pull request %s in %s at %s\n", activity.Payload.Action, activity.Repo.Name, activity.CreatedAt)
		case "CreateEvent":
			fmt.Printf("Created %s in %s at %s\n", activity.Payload.RefType, activity.Repo.Name, activity.CreatedAt)
		case "WatchEvent":
			fmt.Printf("Started to watch %s at %s\n", activity.Repo.Name, activity.CreatedAt)
		case "IssuesEvent":
			fmt.Printf("%s issue in %s at %s\n", activity.Payload.Action, activity.Repo.Name, activity.CreatedAt)
		case "IssueCommentEvent":
			fmt.Printf("%s comment in %s at %s\n", activity.Payload.Action, activity.Repo.Name, activity.CreatedAt)
		default:
			fmt.Printf("Other event: %s in %s at %s\n", activity.Type, activity.Repo.Name, activity.CreatedAt)
		}
	}

}
