package main

import (
	"testing"
)

func TestMakeGreeting(t *testing.T) {
    want := "Hello, Taro"
	got := makeGreeting("Taro")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
