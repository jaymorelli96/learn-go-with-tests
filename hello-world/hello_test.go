package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Jay")
	want := "Hello, Jay"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
