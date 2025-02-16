package necrosword

import (
	"knull/necrosword/model"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Execute() {
	log.Println("Running job with necrosword")

	jobFile, err := os.ReadFile("/Users/deepakraj/Documents/Deepak/Projects/knull/src/necrosword/test.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var job model.Job
	err = yaml.Unmarshal(jobFile, &job)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
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

		Shell(stages.Stage.Cmd)
	}
}
