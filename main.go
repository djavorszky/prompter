// Package prompter is a small library which helps in setting up prompts
// to be answered by a user running an application on the command line.
// The answers provided by the user will be returned in a _ format.
package prompter

var questions = make([]simpleQ, 0)

// Prompt starts the prompting process, starting the series of questions
// with the first one it received via the Ask methods
func Prompt() {

}

// Ask stores a question that will be prompted for the user once the Prompt
// method is called. `Question` is the message that will be shown to the user
// while `defAns` is what will be saved if no answer is specified - that is,
// if the user presses enter without typing anything
func Ask(question, defAns string) {
	questions = append(questions, simpleQ{question, defAns})
}

type simpleQ struct {
	question, defAns string
}
