package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

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
