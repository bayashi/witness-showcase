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

func TestNil(t *testing.T) {
	w.Got(nil).Fail(t, "Expect not <nil>, but got <nil>")
}

func TestTrue(t *testing.T) {
	w.Got(false).Expect(true).Fail(t, "Expect <true>, but got <false>")
	w.Got("true").Expect(true).Fail(t, `Expect <true>, but got "false"`) // Not happen in Go though :-)
}

func TestStruct(t *testing.T) {
	g := struct{ ID string }{ ID: "a001" }
	e := struct{}{}
	w.Got(g).Expect(e).Fail(t, "Not same struct")
}

func TestPointer(t *testing.T) {
	g := "aiko"
	e := &g
	w.Got(&g).Expect(&e).Fail(t, "Expect pointer of pointer")
}