package main

import (
	"fmt"
	"os"
	"testing"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func BenchmarkEcho2(b *testing.B) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{"echo1", "blargle", "flargle"}

	for i := 0; i < b.N; i++ {
		echo2()
	}
}
