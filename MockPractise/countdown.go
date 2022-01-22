package count

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

func countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
