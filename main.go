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
	var cr int // characters removed
	np := make([]byte, len(p))
	for i, c := range p {
		switch c {
		case '\n', ' ', '\t', '\r':
			cr++
			continue
		}
		np[i-cr] = c
	}
	bw, err := nw.Writer.Write(np[0 : len(np)-cr])
	if err != nil {
		return bw, err
	}
	return bw + cr, err
}
