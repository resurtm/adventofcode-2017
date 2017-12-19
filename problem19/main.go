package problem19

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
)

type testCase struct {
	input  string
	result string
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", result: "ABCDEF"},
	{input: "input1.txt", result: "-1"},
}

var testCasesPartTwo = []testCase{
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input     : %#v\n", v.input)
		fmt.Printf("expected  : %s\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartOne())
		fmt.Print("\n\n")
	}
}

func RunPartTwo() {
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input     : %#v\n", v.input)
		fmt.Printf("expected  : %s\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase) solvePartOne() string {
	tiles := readInput(tc.input)

	sx, sy := 0, 1

	px, py := 0, 0
	for i, m := range tiles[0] {
		if m == '|' {
			px = i
		}
	}

	lookAround := func() (up, right, down, left byte) {
		// up
		if py > 0 && len(tiles[py-1]) > px && px >= 0 {
			up = tiles[py-1][px]
		} else {
			up = ' '
		}
		// right
		if len(tiles[py])-1 > px {
			right = tiles[py][px+1]
		} else {
			right = ' '
		}
		// down
		if len(tiles)-1 > py && px >= 0 && px < len(tiles[py+1]) {
			down = tiles[py+1][px]
		} else {
			down = ' '
		}
		// left
		if px > 0 {
			left = tiles[py][px-1]
		} else {
			left = ' '
		}
		return
	}

	convAround := func() (up, right, down, left string) {
		u, r, d, l := lookAround()
		up, right = "<"+string(u)+">", "<"+string(r)+">"
		down, left = "<"+string(d)+">", "<"+string(l)+">"
		return
	}

	result := ""

	for {
		u, r, d, l := lookAround()
		fmt.Println(convAround())

		if u == ' ' && r == ' ' && d == ' ' && l == ' ' {
			break
		}

		tile := tiles[py][px]

		if tile != '-' && tile != '|' && tile != '+' {
			result += string(tile)
			fmt.Println(result)
		}

		if tile == '+' {
			if !(sx == 0 && sy == 1) && u != ' ' && u != '-' {
				sx, sy = 0, -1
			} else if !(sx == 0 && sy == -1) && d != ' ' && d != '-' {
				sx, sy = 0, 1
			} else if !(sx == 1 && sy == 0) && l != ' ' && l != '|' {
				sx, sy = -1, 0
			} else if !(sx == -1 && sy == 0) && r != ' ' && r != '|' {
				sx, sy = 1, 0
			}
		}

		px += sx
		py += sy
	}

	return result
}

func (tc *testCase) solvePartTwo() string {
	return ""
}

func readInput(fn string) [][]byte {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem19"
	} else {
		fp = filepath.Dir(fp)
	}
	fp += "/" + fn

	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func getCurrDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
