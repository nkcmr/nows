package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := io.Copy(noWhitespaceWriter{Writer: os.Stdout}, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "nows: %s\n", err.Error())
		os.Exit(1)
		return
	}
}

type noWhitespaceWriter struct {
	io.Writer
}

func (nw noWhitespaceWriter) Write(p []byte) (n int, err error) {
	whitespaceRemoved := p[:0]
	for _, c := range p {
		switch c {
		case '\n', ' ', '\t', '\r':
			continue
		default:
			whitespaceRemoved = append(whitespaceRemoved, c)
		}
	}
	diff := len(p) - len(whitespaceRemoved)
	n, err = nw.Writer.Write(whitespaceRemoved)
	return n + diff, err
}
