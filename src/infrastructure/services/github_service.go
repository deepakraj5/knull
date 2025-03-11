package services

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func UpdateCommitStatus(token, owner, repo, commitSHA, state, targetURL, description, statusContext string) {
	// Create a context
	ctx := context.Background()

	// Create an OAuth2 client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	// Create the status
	status := &github.RepoStatus{
		State:       github.String(state),
		TargetURL:   github.String(targetURL),
		Description: github.String(description),
		Context:     github.String(statusContext),
	}

	// Update the commit status
	_, _, err := client.Repositories.CreateStatus(ctx, owner, repo, commitSHA, status)
	if err != nil {
		log.Printf("failed to update commit status: %v", err)
	}
}
