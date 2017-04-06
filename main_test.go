package prompter

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/djavorszky/sutils"
)

const (
	input        = "Nothing\n"
	inputWindows = "Nothing\r\n"

	noInput        = "\n"
	noInputWindows = "\r\n"

	numInput        = "1\n"
	numInputWindows = "1\r\n"

	boolInput        = "y\n"
	boolInputWindows = "y\r\n"

	defAns     = "Something"
	boolDefAns = "n"
)

var buf bytes.Buffer

func setup(in string) {
	SetPrompt(">")

	In = strings.NewReader(in)

	buf.Reset()
	Out = &buf
}

func TestAsk(t *testing.T) {
	for i, ins := range []string{input, inputWindows} {
		setup(ins)

		q := "TestOne_" + strconv.Itoa(i)
		expectedOut := fmt.Sprintf("%s\n> ", q)

		ans := Ask(q)

		if msg, ok := expect(expectedOut, buf.String()); !ok {
			t.Error(msg)
		}

		if msg, ok := expect(sutils.TrimNL(ins), ans); !ok {
			t.Error(msg)
		}
	}

}

func TestAskDef(t *testing.T) {
	for i, ins := range []string{input, inputWindows} {
		setup(ins)

		q := "TestTwo_" + strconv.Itoa(i)
		expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

		ans := AskDef(q, defAns)

		if msg, ok := expect(expectedOut, buf.String()); !ok {
			t.Error(msg)
		}

		if msg, ok := expect(sutils.TrimNL(ins), ans); !ok {
			t.Error(msg)
		}
	}
}

func TestAskDefNoInput(t *testing.T) {
	for i, ins := range []string{noInput, noInputWindows} {

		setup(ins)

		q := "TestThree_" + strconv.Itoa(i)
		expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

		ans := AskDef(q, defAns)

		if msg, ok := expect(expectedOut, buf.String()); !ok {
			t.Error(msg)
		}

		if msg, ok := expect(defAns, ans); !ok {
			t.Error(msg)
		}
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

func TestSelectionValidInput(t *testing.T) {
	for i, ins := range []string{numInput, numInputWindows} {
		setup(ins)

		q := "TestFive_" + strconv.Itoa(i)
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

		if msg, ok := expect(sutils.TrimNL(ins), strconv.Itoa(ans)); !ok {
			t.Error(msg)
		}
	}
}

func TestSelectionInValidInput(t *testing.T) {
	setup(input)

	q := "TestSix_Invalid"

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
	for i, ins := range []string{numInput, numInputWindows} {
		setup(ins)

		q := "TestSeven_" + strconv.Itoa(i)

		const defAns = 1
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

		if msg, ok := expect(sutils.TrimNL(ins), strconv.Itoa(ans)); !ok {
			t.Error(msg)
		}
	}
}

func TestSelectionDefInValidInput(t *testing.T) {
	setup(input)

	q := "TestSeven_Invalid"

	const defAns = 1
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

func TestAskBoolDef(t *testing.T) {
	for i, ins := range []string{boolInput, boolInputWindows} {
		setup(ins)

		expectedAns := true

		q := "TestTen_" + strconv.Itoa(i)
		expectedOut := fmt.Sprintf("%s (y/n) (%v):\n> ", q, boolDefAns)

		ans := AskBoolDef(q, false)

		if msg, ok := expect(expectedOut, buf.String()); !ok {
			t.Error(msg)
		}

		if msg, ok := expect(expectedAns, ans); !ok {
			t.Error(msg)
		}
	}
}

func TestAskBoolDefNoInput(t *testing.T) {
	for i, ins := range []string{noInput, noInputWindows} {
		setup(ins)

		expectedAns := false

		q := "TestEleven_" + strconv.Itoa(i)
		expectedOut := fmt.Sprintf("%s (y/n) (%v):\n> ", q, boolDefAns)

		ans := AskBoolDef(q, false)

		if msg, ok := expect(expectedOut, buf.String()); !ok {
			t.Error(msg)
		}

		if msg, ok := expect(expectedAns, ans); !ok {
			t.Error(msg)
		}
	}
}

func expect(expected, actual interface{}) (string, bool) {
	if actual != expected {
		return fmt.Sprintf("Mismatch: Expected %q, got %q", expected, actual), false
	}

	return "", true
}
