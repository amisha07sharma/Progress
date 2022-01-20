package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtraction(t *testing.T) {
	got := Subtract(9, 2)
	want := 7

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
