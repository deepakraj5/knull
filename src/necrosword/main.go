package necrosword

import (
	"knull/infrastructure/services"
	"knull/internal/dtos"
	"knull/necrosword/model"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Execute(jobFilePath string, dir string, body dtos.GithubWebhookRequestBody, githubPat string) {
	log.Println("Running job with necrosword")

	jobFile, err := os.ReadFile(jobFilePath)
	if err != nil {
		log.Printf("Failed to read YAML file: %v", err)
		services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
		return
	}

	var job model.Job
	err = yaml.Unmarshal(jobFile, &job)
	if err != nil {
		log.Printf("Failed to unmarshal YAML: %v", err)
		services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
		return
	}

	log.Printf("Running job %d with id: %s", job.Id, job.Name)

	log.Println("Env's")
	for _, env := range job.Environment {
		for key, value := range env {
			log.Printf("%s - %s", key, value)
		}
	}

	log.Println("Stages:")
	for _, stages := range job.Stages {
		log.Printf("Executing stage: %s, with command: %s", stages.Stage.Name, stages.Stage.Cmd)

		err := Shell(stages.Stage.Cmd, dir)

		if err != nil {
			services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "failure", "https://knull.com", "error", "ci/knull")
			return
		}
	}
	services.UpdateCommitStatus(githubPat, body.Repository.Owner.Name, body.Repository.Name, body.HeadCommit.Id, "success", "https://knull.com", "building", "ci/knull")
}
