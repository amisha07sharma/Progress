package shapes

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Area of Rectangle", func(t *testing.T) {
		rect := Rectangle{2, 3}
		got := rect.Area()
		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %q but got %q", want, got)
		}
	})
}

func Area(rect Rectangle) {
	panic("unimplemented")
}
