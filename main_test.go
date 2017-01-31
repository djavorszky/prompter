package prompter

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const (
	input       = "Nothing\n"
	noInput     = "\n"
	numberInput = "1"
	defAns      = "Something"
)

var buf bytes.Buffer

func setup(in string) {
	SetPrompt(">")

	In = strings.NewReader(in)

	buf.Reset()
	Out = &buf
}

func TestAsk(t *testing.T) {
	setup(input)

	const q = "TestOne"
	expectedOut := fmt.Sprintf("%s\n> ", q)

	ans := Ask(q)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(strings.TrimSuffix(input, "\n"), ans); !ok {
		t.Error(msg)
	}
}

func TestAskDef(t *testing.T) {
	setup(input)

	const q = "TestTwo"
	expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

	ans := AskDef(q, defAns)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(strings.TrimSuffix(input, "\n"), ans); !ok {
		t.Error(msg)
	}
}

func TestAskDefNoInput(t *testing.T) {
	setup(noInput)

	const q = "TestThree"
	expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

	ans := AskDef(q, defAns)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(defAns, ans); !ok {
		t.Error(msg)
	}
}

func TestSetPrompt(t *testing.T) {
	setup(input)

	const q, p = "TestFour", "?"
	expectedOut := fmt.Sprintf("%s\n%s ", q, p)

	SetPrompt(p)

	_ = Ask(q)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Error(msg)
	}
}

func expect(expected, actual string) (string, bool) {
	if actual != expected {
		return fmt.Sprintf("Mismatch: Expected %q, got %q", expected, actual), false
	}

	return "", true
}

func TestSelectionValidInput(t *testing.T) {
	setup(numberInput)

	const q = "TestFive"
	options := []string{"one", "two"}

	var expectedOut bytes.Buffer

	expectedOut.WriteString(fmt.Sprintf("%s\n", q))
	for i, o := range options {
		expectedOut.WriteString(fmt.Sprintf("  [%d] %s\n", i, o))
	}
	expectedOut.WriteString("> ")

	ans, ok := AskSelection(q, options)
	if !ok {
		t.Error("AskSelection failed due to invalid input where valid input was provided.")
	}

	if msg, ok := expect(expectedOut.String(), buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(numberInput, strconv.Itoa(ans)); !ok {
		t.Error(msg)
	}
}

func TestSelectionInValidInput(t *testing.T) {
	setup(input)

	const q = "TestSix"
	options := []string{"one", "two"}

	var expectedOut bytes.Buffer

	expectedOut.WriteString(fmt.Sprintf("%s\n", q))
	for i, o := range options {
		expectedOut.WriteString(fmt.Sprintf("  [%d] %s\n", i, o))
	}
	expectedOut.WriteString("> ")
	expectedOut.WriteString("Invalid input. Can only be between 0-1\n")

	ans, ok := AskSelection(q, options)
	if ok {
		t.Error("AskSelection provided positive response even though it should have failed with invalid input.")
	}

	if msg, ok := expect(expectedOut.String(), buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect("0", strconv.Itoa(ans)); !ok {
		t.Error(msg)
	}
}

func TestSelectionDefValidInput(t *testing.T) {
	setup(numberInput)

	const q, defAns = "TestSeven", 1
	options := []string{"one", "two"}

	var expectedOut bytes.Buffer

	expectedOut.WriteString(fmt.Sprintf("%s (default: %d)\n", q, defAns))
	for i, o := range options {
		expectedOut.WriteString(fmt.Sprintf("  [%d] %s\n", i, o))
	}
	expectedOut.WriteString("> ")

	ans, ok := AskSelectionDef(q, defAns, options)
	if !ok {
		t.Error("AskSelectionDef failed due to invalid input where valid input was provided.")
	}

	if msg, ok := expect(expectedOut.String(), buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(numberInput, strconv.Itoa(ans)); !ok {
		t.Error(msg)
	}
}

func TestSelectionDefInValidInput(t *testing.T) {
	setup(input)

	const q, defAns = "TestEight", 1
	options := []string{"one", "two"}

	var expectedOut bytes.Buffer

	expectedOut.WriteString(fmt.Sprintf("%s (default: %d)\n", q, defAns))
	for i, o := range options {
		expectedOut.WriteString(fmt.Sprintf("  [%d] %s\n", i, o))
	}
	expectedOut.WriteString("> ")
	expectedOut.WriteString("Invalid input. Can only be between 0-1\n")

	ans, ok := AskSelectionDef(q, defAns, options)
	if ok {
		t.Error("AskSelectionDef provided positive response even though it should have failed with invalid input.")
	}

	if msg, ok := expect(expectedOut.String(), buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect("0", strconv.Itoa(ans)); !ok {
		t.Error(msg)
	}
}

func TestSelectionDefNoInput(t *testing.T) {
	setup(noInput)

	const q, defAns = "TestNine", 0
	options := []string{"one", "two"}

	var expectedOut bytes.Buffer

	expectedOut.WriteString(fmt.Sprintf("%s (default: %d)\n", q, defAns))
	for i, o := range options {
		expectedOut.WriteString(fmt.Sprintf("  [%d] %s\n", i, o))
	}
	expectedOut.WriteString("> ")

	ans, ok := AskSelectionDef(q, defAns, options)
	if !ok {
		t.Error("AskSelectionDef failed due to invalid input where valid input was provided.")
	}

	if msg, ok := expect(expectedOut.String(), buf.String()); !ok {
		t.Error(msg)
	}

	if msg, ok := expect(strconv.Itoa(defAns), strconv.Itoa(ans)); !ok {
		t.Error(msg)
	}
}
