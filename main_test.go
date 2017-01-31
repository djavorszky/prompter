package prompter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	input   = "Nothing\n"
	noInput = "\n"
	defAns  = "Something"
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
		t.Errorf(msg)
	}

	if msg, ok := expect(strings.TrimSuffix(input, "\n"), ans); !ok {
		t.Errorf(msg)
	}
}

func TestAskDef(t *testing.T) {
	setup(input)

	const q = "TestTwo"
	expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

	ans := AskDef(q, defAns)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Errorf(msg)
	}

	if msg, ok := expect(strings.TrimSuffix(input, "\n"), ans); !ok {
		t.Errorf(msg)
	}
}

func TestAskDefNoInput(t *testing.T) {
	setup(noInput)

	const q = "TestThree"
	expectedOut := fmt.Sprintf("%s (%s):\n> ", q, defAns)

	ans := AskDef(q, defAns)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Errorf(msg)
	}

	if msg, ok := expect(defAns, ans); !ok {
		t.Errorf(msg)
	}
}

func TestSetPrompt(t *testing.T) {
	setup(input)

	const q, p = "TestFour", "?"
	expectedOut := fmt.Sprintf("%s\n%s ", q, p)

	SetPrompt(p)

	_ = Ask(q)

	if msg, ok := expect(expectedOut, buf.String()); !ok {
		t.Errorf(msg)
	}
}

func expect(expected, actual string) (string, bool) {
	if actual != expected {
		return fmt.Sprintf("Mismatch: Expected %q, got %q", expected, actual), false
	}

	return "", true
}
