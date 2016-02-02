package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// does this work on windows?
func InputAreYouSure() bool {
	return InputAffirmative("Are you sure?")
}

func InputAffirmative(message string) bool {
	fmt.Printf("%s (yes)\n", Cyan(message))
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text) == "yes"
}
