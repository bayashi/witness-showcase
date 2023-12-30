package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

func TestNil(t *testing.T) {
	w.Got(nil).Fail(t, "Expect not <nil>, but got <nil>")
}
