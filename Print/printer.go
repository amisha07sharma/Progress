package Printer

import (
	"bytes"
	"fmt"
)

func printer(writer *bytes.Buffer, str string) {
	fmt.Printf("Hello, %s", str)
}
