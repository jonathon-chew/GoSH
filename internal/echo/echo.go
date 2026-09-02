package echo

import (
	"fmt"
	"io"
	"os"
	"slices"
)

func Echo() {
	var buff = make([]byte, 32*1024)

	for {
		n, ErrReadingSTDIN := os.Stdin.Read(buff)

		if n > 0 {
			os.Stdout.Write(buff[:n])
		}

		if ErrReadingSTDIN == io.EOF {
			if !slices.Contains(os.Args[1:], "-n") {
				os.Stdout.Write([]byte("\n"))
			}
			return
		}

		if ErrReadingSTDIN != nil && ErrReadingSTDIN != io.EOF {
			panic(fmt.Errorf("[error]: reading from standard in: %s", ErrReadingSTDIN))
		}
	}
}
