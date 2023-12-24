package main

import (
	"fmt"
	"testing"

	w "github.com/bayashi/witness"
)

func TestError(t *testing.T) {
	err := fmt.Errorf("Some error") // unexpectedly

	if err != nil {
		w.Got(err).Fail(t, "Expect no error")
	}
}

func TestWrongString(t *testing.T) {
	g := "a\nb\nc"
	e := "a\nd\nc"

	if g != e {
		w.Got(g).Expect(e).Fail(t, "These should be same, but not same")
	}
}

func TestWrongStringWithDiff(t *testing.T) {
	g := "a\nb\nc"
	e := "a\nd\nc"

	if g != e {
		w.Got(g).Expect(e).ShowDiff().Fail(t, "These should be same, but not same")
	}
}

func TestWrongStringWithRawOutput(t *testing.T) {
	g := "a\nb\nc"
	e := "a\nd\nc"

	if g != e {
		w.Got(g).Expect(e).ShowRaw().Fail(t, "These should be same, but not same")
	}
}

func TestWrongStringWithAllOptionalOutput(t *testing.T) {
	g := "a\nb\nc"
	e := "a\nd\nc"

	if g != e {
		w.Got(g).Expect(e).ShowAll().Fail(t, "These should be same, but not same")
	}
}

func TestWrongInt(t *testing.T) {
	g := 123
	e := 321

	if g != e {
		w.Got(g).Expect(e).ShowAll().Fail(t, "These should be same, but not same")
	}
}