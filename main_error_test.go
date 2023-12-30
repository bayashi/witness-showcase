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
