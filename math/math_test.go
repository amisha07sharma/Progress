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

// Table Drivem tests

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{1, 2, 3},
	addTest{2, 3, 5},
	addTest{8, 4, 12},
	addTest{3, 10, 13},
}

func AddTestTable(t *testing.T) {
	for _, test := range addTests {
		output := Add(test.arg1, test.arg2)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
