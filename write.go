package command

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Writer interface {
    Error(errorMessage string, error error)
    Warn(errorMessage string, error error)
    Write(text string)
}

type DefaultWriter struct {}

var Cyan = color.New(color.FgCyan).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Bold = color.New(color.FgWhite, color.Bold).SprintFunc()

func (d DefaultWriter) Error(errorMessage string, error error) {
	fmt.Println()
	fmt.Println("---------------------------------------------")
	fmt.Fprintf(os.Stderr, "ERROR\n")
	fmt.Fprintf(os.Stderr, "%s: %s\n", Cyan(errorMessage), Red(error))
	fmt.Println("---------------------------------------------")
	fmt.Println()
	os.Exit(1)
}

func (d DefaultWriter) Warn(errorMessage string, error error) {
	fmt.Fprintf(os.Stderr, "WARNING\n")
	fmt.Fprintf(os.Stderr, "%s: %s\n", Cyan(errorMessage), Yellow(error))
	fmt.Println()
}

func (d DefaultWriter) Write(text string) {
	fmt.Printf("%s\n", text)
}

// Probably a nicer way to do this
func Error(errorMessage string, error error) {
	DefaultWriter{}.Error(errorMessage, error)
}

func Warn(errorMessage string, error error) {
	DefaultWriter{}.Warn(errorMessage, error)
}

func Write(text string) {
	DefaultWriter{}.Write(text)
}
