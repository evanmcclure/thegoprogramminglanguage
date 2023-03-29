package main

import (
	"fmt"
	"os"
	"testing"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func BenchmarkEcho1(b *testing.B) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{"echo1", "blargle", "flargle"}

	for i := 0; i < b.N; i++ {
		echo1()
	}
}
