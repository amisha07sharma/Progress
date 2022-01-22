package count

import (
	"bytes"
	"testing"
)

type MockSleeper struct {
	call int
}

func (s *MockSleeper) Sleep() {
	s.call++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleep := MockSleeper{}
	countdown(&buffer, &sleep)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
	if sleep.call != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", sleep.call)
	}
}
