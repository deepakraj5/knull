package handlers

import (
	"knull/internal/dtos"
	"knull/internal/utils"
	"knull/necrosword"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GitHubWebhook(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var headers dtos.GithubWebhookRequestHeaders
	var body dtos.GithubWebhookRequestBody

	err = utils.ParseRequest(r, &headers, &body)
	if err != nil {
		log.Printf("Error parsing request: %v\n", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	var branch string = "refs/heads/development"
	var event string = "push"

	if headers.XGitHubEvent == event && body.Ref == branch {
		log.Printf("New commit has been pushed for branch: %s in repository: %s", body.Ref, body.Repository.Name)

		// create a directory in the workspace for the repository
		repoDir := "../workspace/" + body.Repository.Name
		err := os.MkdirAll(repoDir, 0755) // 0755 -> The owner can read, write, execute. Everyone else can read and execute but not modify the file.
		if err != nil {
			log.Printf("Failed to create directory: %v\n", err)
			return
		}

		isEmptyDir, err := utils.IsDirEmpty(repoDir)
		if err != nil {
			log.Printf("Error checking directory: %v\n", err)
			return
		}

		if isEmptyDir {
			// clone the given repo
			gitRepo := strings.ReplaceAll(body.Repository.CloneUrl, "https://", "")
			gitCloneCmd := "git clone --single-branch -b development https://oauth2:" + os.Getenv("GITHUB_PAT") + "@" + gitRepo + " ."
			necrosword.Shell(gitCloneCmd, repoDir)

			log.Printf("Cloned the repo: %s successfully", body.Repository.Name)
		} else {
			// pull the latest changes
			necrosword.Shell("git pull", repoDir)

			log.Printf("Pulled the latest changes from the repo: %s successfully", body.Repository.Name)
		}

		// execute pipeline file from cloned repo
		var jobFilePath string = repoDir + "/knull.yaml"
		necrosword.Execute(jobFilePath, repoDir)
	}

	payload := dtos.ResponseDto{
		ResponseCode: 200,
	}

	utils.JsonResponse(w, payload)
}
