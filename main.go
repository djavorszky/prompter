// Package prompter is a small library which helps in setting up prompts
// to be answered by a user running an application on the command line.
// The answers provided by the user will be returned as strings.
//
// Question marks are not added for the questions, that should be done
// by the caller.
package prompter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cliPrompt = "> "

// SetPrompt sets the command line prompt character. Also adds a space at the end.
// The default is "> "
func SetPrompt(prompt string) {
	cliPrompt = fmt.Sprintf("%s ", prompt)
}

// Ask prompts the user for input. The value of `question` will be shown to the user.
// There is no default answer, so if no answer is provided, then an empty string
// will be returned.
func Ask(question string) string {
	fmt.Println(question)

	return prompt()
}

// AskDef prompts the user for input.he value of `question` will be shown to the
// user, while `defAns` is what will be saved if no answer is provided.
func AskDef(question, defAns string) string {
	fmt.Printf("%s (%s):\n", question, defAns)
	ans := prompt()

	if ans == "\n" {
		ans = defAns
	}

	return ans
}

// AskSecret prompts the user for an input that should not show on the terminal
// either by replacing the characters with asterisks, or not showing anything at all.
//
// Currently, this does not happen, however a warning message is shown after the
// question.
func AskSecret(question string) string {
	fmt.Printf("%s - %s\n", question, "WARNING: What you type will be shown!")

	return prompt()
}

// AskSelection takes a slice of strings to display as a selection box in
// the form of `[index] question`, from which the user can choose easily.
// Returns the chosen index and a true if correctly chosen, or empty string
// and false if a non-number was specified or if number was out of range
// of the selections.
func AskSelection(question string, options []string) (string, bool) {
	fmt.Println(question)
	for i, v := range options {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	intAns, err := strconv.Atoi(prompt())
	if err != nil || intAns < 0 || intAns > len(options)-1 {
		fmt.Printf("Invalid input. Can only be between 0-%d\n", len(options)-1)
		return "", false
	}

	return strconv.Itoa(intAns), true
}

// AskSelectionDef takes a slice of strings to display as a selection box in
// the form of `[index] question`, from which the user can choose easily.
// Returns the chosen index and a true if correctly chosen, or empty string
// and false if a non-number was specified or if number was out of range
// of the selections.
//
// Also shows a default selection which will be chosen if no input is
// specified.
func AskSelectionDef(question string, defAns int, options []string) (string, bool) {
	if defAns < 0 || defAns > len(options)-1 {
		fmt.Print("Default answer was out of bounds of number of options.")
		return "", false
	}

	fmt.Printf("%s (default: %d)\n", question, defAns)
	for i, v := range options {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	ans := prompt()

	if ans == "" {
		return strconv.Itoa(defAns), true
	}

	intAns, err := strconv.Atoi(ans)
	if err != nil || intAns < 0 || intAns > len(options)-1 {
		fmt.Printf("Invalid input. Can only be between 0-%d\n", len(options)-1)
		return "", false
	}

	return strconv.Itoa(intAns), true
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(cliPrompt)
	ans, _ := reader.ReadString('\n')
	return strings.TrimSuffix(ans, "\n")
}
