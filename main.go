package main

import (
	"os"
	"github.com/resurtm/adventofcode-2017/problem1"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	case "--problem1":
		problem1.Run()
	default:
		panic("nothing to do")
	}
}
