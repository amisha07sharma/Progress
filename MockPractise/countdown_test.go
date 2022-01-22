package count

import (
	"bytes"
	"testing"
)

type Sleeper struct {
	call int
}

func (s *Sleeper) Sleep() {
	s.call++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleep := Sleeper{}
	countdown(&buffer, &sleep)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
