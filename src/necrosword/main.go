package necrosword

import (
	"knull/necrosword/model"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Execute(jobFilePath string, dir string) {
	log.Println("Running job with necrosword")

	jobFile, err := os.ReadFile(jobFilePath)
	if err != nil {
		log.Printf("Failed to read YAML file: %v", err)
		return
	}

	var job model.Job
	err = yaml.Unmarshal(jobFile, &job)
	if err != nil {
		log.Printf("Failed to unmarshal YAML: %v", err)
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

		Shell(stages.Stage.Cmd, dir)
	}
}
