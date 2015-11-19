package command

import (
	"bufio"
	"os/exec"
)

type Command struct {
	Writer Writer
}

func Default() *Command {
    return New(DefaultWriter{})
}

func New(writer Writer) *Command {
    return &Command{ Writer:writer }
}

func (self Command) Execute(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		self.Writer.Error("Error creating StdoutPipe for Cmd", err)
	}

	cmdErrorReader, err := cmd.StderrPipe()
	if err != nil {
		self.Writer.Error("Error creating StderrPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			self.Writer.Write(scanner.Text())
		}
	}()

	errorScanner := bufio.NewScanner(cmdErrorReader)
	go func() {
		for errorScanner.Scan() {
			self.Writer.Write(errorScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		self.Writer.Error("Error starting Cmd", err)
	}

	err = cmd.Wait()
	if err != nil {
		self.Writer.Warn("Error waiting for Cmd", err)
	}
}

func (self Command) Error(errorMessage string, error error) {
	self.Writer.Error(errorMessage, error)
}

func (self Command) Warn(errorMessage string, error error) {
	self.Writer.Warn(errorMessage, error)
}

func (self Command) Write(text string) {
	self.Writer.Write(text)
}
