package Printer

import (
	"bytes"
	"testing"
)

func TestPrinter(t *testing.T) {

	buffer := bytes.Buffer{}
	printer(&buffer, "World")
	got := buffer.String()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPrinter2(t *testing.T) {
	buffer := bytes.Buffer{}
	printer2(&buffer, "Amisha")
	got := buffer.String()
	want := "Hello, Amisha"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
