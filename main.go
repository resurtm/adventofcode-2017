package main

import (
	"os"
	"github.com/resurtm/adventofcode-2017/problem10"
	"github.com/resurtm/adventofcode-2017/problem01"
	"github.com/resurtm/adventofcode-2017/problem02"
	"github.com/resurtm/adventofcode-2017/problem03"
	"github.com/resurtm/adventofcode-2017/problem04"
	"github.com/resurtm/adventofcode-2017/problem05"
	"github.com/resurtm/adventofcode-2017/problem06"
	"github.com/resurtm/adventofcode-2017/problem07"
	"github.com/resurtm/adventofcode-2017/problem08"
	"github.com/resurtm/adventofcode-2017/problem09"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	case "--problem1":
	case "--problem01":
		problem01.RunPartOne()
		problem01.RunPartTwo()
	case "--problem2":
	case "--problem02":
		problem02.RunPartOne()
		problem02.RunPartTwo()
	case "--problem3":
	case "--problem03":
		problem03.RunPartOne()
		problem03.RunPartTwo()
	case "--problem4":
	case "--problem04":
		problem04.RunPartOne()
		problem04.RunPartTwo()
	case "--problem5":
	case "--problem05":
		problem05.RunPartOne()
		problem05.RunPartTwo()
	case "--problem6":
	case "--problem06":
		problem06.RunPartOne()
		problem06.RunPartTwo()
	case "--problem7":
	case "--problem07":
		problem07.RunPartOne()
		problem07.RunPartTwo()
	case "--problem8":
	case "--problem08":
		problem08.RunPartOne()
		problem08.RunPartTwo()
	case "--problem9":
	case "--problem09":
		problem09.RunPartOne()
		problem09.RunPartTwo()
	case "--problem10":
		problem10.RunPartOne()
		problem10.RunPartTwo()
	default:
		panic("nothing to do")
	}
}
