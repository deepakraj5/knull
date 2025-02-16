package necrosword

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

func Shell(command string) {

	executableCommand := strings.Fields(command)

	cmd := exec.Command(executableCommand[0], executableCommand[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error creating stdout pipe: %v\n", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Error creating stderr pipe: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command %v\n", err)
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command: %v\n", err)
		return
	}

	log.Printf("Command Executed successfully %s", command)
}
