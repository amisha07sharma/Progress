package shapes

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want int) {
		t.Helper()
		got := shape.Area()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %q but got %q", want, got)
		}
	}

	t.Run("Area of Rectangle", func(t *testing.T) {
		rect := Rectangle{2, 3}
		want := 6

		checkArea(t, rect, want)
	})

	t.Run("Area of Square", func(t *testing.T) {
		sq := square{4}
		want := 16

		checkArea(t, sq, want)
	})
}

func Area(rect Rectangle) {
	panic("unimplemented")
}
