package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func BenchmarkEcho3(b *testing.B) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{"echo1", "blargle", "flargle"}

	for i := 0; i < b.N; i++ {
		echo3()
	}
}
