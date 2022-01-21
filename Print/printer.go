package Printer

import (
	"bytes"
	"fmt"
)

func printer(writer *bytes.Buffer, str string) {
	fmt.Fprintf(writer, "Hello, %s", str)
}
