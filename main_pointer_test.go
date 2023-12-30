package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

func TestPointer(t *testing.T) {
	g := "aiko"
	e := &g
	w.Got(&g).Expect(&e).Fail(t, "Expect pointer of pointer")
}
