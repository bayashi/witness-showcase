package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

func TestStruct(t *testing.T) {
	g := struct{ ID string }{ ID: "a001" }
	e := struct{}{}
	w.Got(g).Expect(e).Fail(t, "Not same struct")
}
