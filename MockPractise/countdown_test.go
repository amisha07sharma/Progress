package count

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	countdown(&buffer)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
