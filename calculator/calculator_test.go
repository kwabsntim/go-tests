package main

import (
	"testing"
)

//writing tests for the functions in main

func TestAdd(t *testing.T) {
	got := Add(5, 5)
	want := 10
	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}
func TestSub(t *testing.T) {
	got := Sub(7, 5)
	want := 2
	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}
