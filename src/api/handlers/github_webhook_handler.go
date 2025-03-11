package handlers

import (
	"knull/infrastructure/services"
	"knull/internal/dtos"
	"knull/internal/utils"
	"knull/necrosword"
	"log"
	"net/http"
	"os"
	"strings"
)

func GitHubWebhook(w http.ResponseWriter, r *http.Request) {

	var headers dtos.GithubWebhookRequestHeaders
	var body dtos.GithubWebhookRequestBody

	err := utils.ParseRequest(r, &headers, &body)
	if err != nil {
		log.Printf("Error parsing request: %v\n", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	var branch string = "refs/heads/development"
	var event string = "push"

	if headers.XGitHubEvent == event && body.Ref == branch {
		log.Printf("New commit has been pushed for branch: %s in repository: %s", body.Ref, body.Repository.Name)

		githubPat := os.Getenv("GITHUB_PAT")

		services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "pending", "https://knull.com", "building", "ci/knull")

		// create a directory in the workspace for the repository
		repoDir := "../workspace/" + body.Repository.Name
		err := os.MkdirAll(repoDir, 0755) // 0755 -> The owner can read, write, execute. Everyone else can read and execute but not modify the file.
		if err != nil {
			log.Printf("Failed to create directory: %v\n", err)
			services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
			return
		}

		isEmptyDir, err := utils.IsDirEmpty(repoDir)
		if err != nil {
			log.Printf("Error checking directory: %v\n", err)
			services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
			return
		}

		if isEmptyDir {
			// clone the given repo
			gitRepo := strings.ReplaceAll(body.Repository.CloneUrl, "https://", "")
			gitCloneCmd := "git clone --single-branch -b development https://oauth2:" + githubPat + "@" + gitRepo + " ."

			err := necrosword.Shell(gitCloneCmd, repoDir)

			if err != nil {
				services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
				return
			}

			log.Printf("Cloned the repo: %s successfully", body.Repository.Name)
		} else {
			// pull the latest changes
			err := necrosword.Shell("git pull", repoDir)

			if err != nil {
				services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
				return
			}

			log.Printf("Pulled the latest changes from the repo: %s successfully", body.Repository.Name)
		}

		// execute pipeline file from cloned repo
		var jobFilePath string = repoDir + "/knull.yaml"
		necrosword.Execute(jobFilePath, repoDir, body, githubPat)
	}

	payload := dtos.ResponseDto{
		ResponseCode: 200,
	}

	utils.JsonResponse(w, payload)
}
