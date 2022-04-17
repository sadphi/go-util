package io

import (
	"bufio"
	"io"
	"os"
)

func ReadFromStdIn() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return io.ReadAll(reader)
}
