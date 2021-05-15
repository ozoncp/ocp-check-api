package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("friend")
	want := "Hello, friend!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
