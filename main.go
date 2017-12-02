package main

import (
	"os"
	"github.com/resurtm/adventofcode-2017/problem1"
	"github.com/resurtm/adventofcode-2017/problem2"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	case "--problem1":
		problem1.RunPartOne()
		problem1.RunPartTwo()
	case "--problem2":
		problem2.Run()
	default:
		panic("nothing to do")
	}
}
