package main

import (
	. "go_practice/internal/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
