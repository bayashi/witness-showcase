package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

func TestWrongInt(t *testing.T) {
	g := 123
	e := 321

	if g != e {
		w.Got(g).Expect(e).ShowAll().Fail(t, "These should be same, but not same")
	}
}
