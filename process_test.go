package brc

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var op string

func BenchmarkProcess(b *testing.B) {
	file, err := os.Open("/Users/ashwini/Downloads/measurements.txt")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output, perr := Process(file)
		if perr != nil {
			panic(perr)
		}
		op = output
	}

	b.StopTimer()

	_, _ = fmt.Fprintln(io.Discard, op)

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
