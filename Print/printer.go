package Printer

import (
	"bytes"
	"fmt"
	"io"
)

func printer(writer *bytes.Buffer, str string) {
	fmt.Fprintf(writer, "Hello, %s", str)
}

func printer2(writer io.Writer, str string) {
	fmt.Fprintf(writer, "Hello, %s", str)
}
