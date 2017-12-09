package main

import (
	"os"
	"github.com/resurtm/adventofcode-2017/problem1"
	"github.com/resurtm/adventofcode-2017/problem2"
	"github.com/resurtm/adventofcode-2017/problem3"
	"github.com/resurtm/adventofcode-2017/problem4"
	"github.com/resurtm/adventofcode-2017/problem5"
	"github.com/resurtm/adventofcode-2017/problem6"
	"github.com/resurtm/adventofcode-2017/problem7"
	"github.com/resurtm/adventofcode-2017/problem8"
	"github.com/resurtm/adventofcode-2017/problem9"
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
		problem2.RunPartOne()
		problem2.RunPartTwo()
	case "--problem3":
		problem3.RunPartOne()
		problem3.RunPartTwo()
	case "--problem4":
		problem4.RunPartOne()
		problem4.RunPartTwo()
	case "--problem5":
		problem5.RunPartOne()
		problem5.RunPartTwo()
	case "--problem6":
		problem6.RunPartOne()
		problem6.RunPartTwo()
	case "--problem7":
		problem7.RunPartOne()
		problem7.RunPartTwo()
	case "--problem8":
		problem8.RunPartOne()
		problem8.RunPartTwo()
	case "--problem9":
		problem9.RunPartOne()
		problem9.RunPartTwo()
	default:
		panic("nothing to do")
	}
}
