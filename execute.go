package command

import (
	"bufio"
	"os/exec"
)

func Execute(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		Error("Error creating StdoutPipe for Cmd", err)
	}

	cmdErrorReader, err := cmd.StderrPipe()
	if err != nil {
		Error("Error creating StderrPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			Write(scanner.Text())
		}
	}()

	errorScanner := bufio.NewScanner(cmdErrorReader)
	go func() {
		for errorScanner.Scan() {
			Write(errorScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		Error("Error starting Cmd", err)
	}

	err = cmd.Wait()
	if err != nil {
		Warn("Error waiting for Cmd", err)
	}
}
