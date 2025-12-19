package main

import (
	"Github-user-activity/variables"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("usage: github-acitivity <username>")
		os.Exit(1)
	}
	
	if err := getGithubApi(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	
}

func getGithubApi(username string) error {
	variables.Github_url = fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(variables.Github_url)
	if err != nil {
		return fmt.Errorf("failed to fetch github api: %w", err)
	}
	defer resp.Body.Close()

	var activity []variables.Info
	if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	for _, i := range activity {
		outputActivity(i)
	}

	return nil
}

func outputActivity(activity variables.Info) {
	switch activity.Type {
	case "PushEvent":
		fmt.Printf("- Pushed %d commits to %s.\n", len(activity.Payload.Commits), activity.Repo.Name)
	case "CreateEvent":
		fmt.Printf("- Created %s.\n", activity.Repo.Name)
	case "DeleteEvent":
		fmt.Printf("- Deleted %s.\n", activity.Repo.Name)
	case "ForkEvent":
		fmt.Printf("- Forked %s.\n", activity.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("- %s issue #%d on %s.\n", activity.Payload.Action, activity.Payload.Issue.Id, activity.Repo.Name)
	case "IssueCommentEvent":
		fmt.Printf("- Commented on issue #%d on %s.\n", activity.Payload.Issue.Id, activity.Repo.Name)
	case "WatchEvent":
		fmt.Printf("- Starred %s.\n", activity.Repo.Name)
	case "PullRequestEvent":
		fmt.Printf("- %s pull request #%d on %s.\n", activity.Payload.Action, activity.Payload.PULLRequest.Id, activity.Repo.Name)
	case "PullRequestReviewCommentEvent":
		fmt.Printf("- Commented on pull request #%d on %s.\n", activity.Payload.PULLRequest.Id, activity.Repo.Name)
	}
}

