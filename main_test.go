package main

import "testing"

func TestHelloo(t *testing.T) {
	want := "Hello go"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
