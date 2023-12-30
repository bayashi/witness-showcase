package main

import (
	"testing"

	w "github.com/bayashi/witness"
)

func TestBoolean(t *testing.T) {
	w.Got(false).Expect(true).Fail(t, "Expect <true>, but got <false>")
	w.Got("true").Expect(true).Fail(t, `Expect <true>, but got "false"`) // Not happen in Go though :-)
}
