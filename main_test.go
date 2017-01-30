package prompter

import (
	"io/ioutil"
	"strings"
	"testing"
)

const (
	input   = "Nothing\n"
	noInput = "\n"
	defAns  = "Something"
)

func setup(in string) {
	In = strings.NewReader(in)
	Out = ioutil.Discard
}

func TestAsk(t *testing.T) {
	setup(input)

	ans := Ask("TestOne")

	if ans != "Nothing" {
		t.Errorf("Mismatch; Expected: %q, got %q", input, ans)
	}
}

func TestAskDef(t *testing.T) {
	var ans string

	setup(input)

	ans = AskDef("TestTwo", defAns)

	if ans != "Nothing" {
		t.Errorf("Mismatch; Expected: %q, got %q", input, ans)
	}

	setup(noInput)

	ans = AskDef("TestThree", defAns)

	if ans != defAns {
		t.Errorf("Mismatch; Expected: %q, got %q", defAns, ans)
	}
}
