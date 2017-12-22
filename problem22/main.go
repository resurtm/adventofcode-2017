package problem22

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
)

type testCase struct {
	input  string
	result int
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", result: 5587},
	{input: "input1.txt", result: 5462},
}

var testCasesPartTwo = []testCase{
	{input: "input0.txt", result: 2511944},
	{input: "input1.txt", result: 2512135},
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		fmt.Printf("input     : %#v\n", v.input)
		fmt.Printf("expected  : %d\n", v.result)
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
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Print("result    : ")
		fmt.Print(v.solvePartTwo())
		fmt.Print("\n\n")
	}
}

func (tc *testCase) solvePartOne() int {
	inputData := readInput(tc.input)
	gridSize := len(inputData[0])
	constSize := 50000
	halfSize := constSize / 2
	grid := make([][]bool, constSize)
	for i := 0; i < constSize; i++ {
		grid[i] = make([]bool, constSize)
	}
	for j := 0; j < gridSize; j++ {
		for i := 0; i < gridSize; i++ {
			grid[j+halfSize][i+halfSize] = inputData[j][i] == '#'
		}
	}

	result := 0

	cx, cy, dir := int(gridSize/2)+halfSize, int(gridSize/2)+halfSize, 0
	for iter := 0; iter < 10000; iter++ {
		if grid[cy][cx] {
			dir++
			grid[cy][cx] = false
		} else {
			dir--
			grid[cy][cx] = true
			result++
		}

		if dir < 0 {
			dir = 3
		}
		if dir > 3 {
			dir = 0
		}

		switch dir {
		case 0:
			cy--
		case 1:
			cx++
		case 2:
			cy++
		case 3:
			cx--
		}
	}

	return result
}

func (tc *testCase) solvePartTwo() int {
	inputData := readInput(tc.input)
	gridSize := len(inputData[0])
	constSize := 50000
	halfSize := constSize / 2
	grid := make([][]byte, constSize)
	for i := 0; i < constSize; i++ {
		grid[i] = make([]byte, constSize)
	}
	for j := 0; j < gridSize; j++ {
		for i := 0; i < gridSize; i++ {
			if inputData[j][i] == '#' {
				grid[j+halfSize][i+halfSize] = 1
			}
		}
	}

	result := 0

	cx, cy, dir := int(gridSize/2)+halfSize, int(gridSize/2)+halfSize, 0
	for iter := 0; iter < 10000000; iter++ {
		switch grid[cy][cx] {
		case 0:
			dir--
			grid[cy][cx] = 3
		case 1:
			dir++
			grid[cy][cx] = 2
		case 2:
			if dir == 0 {
				dir = 2
			} else if dir == 1 {
				dir = 3
			} else if dir == 2 {
				dir = 0
			} else if dir == 3 {
				dir = 1
			}
			grid[cy][cx] = 0
		case 3:
			grid[cy][cx] = 1
			result++
		}

		if dir < 0 {
			dir = 3
		}
		if dir > 3 {
			dir = 0
		}

		switch dir {
		case 0:
			cy--
		case 1:
			cx++
		case 2:
			cy++
		case 3:
			cx--
		}
	}

	return result
}

func printBoolGrid(inp [][]bool) {
	fmt.Println("==========")
	for j := 0; j < len(inp); j++ {
		for i := 0; i < len(inp[j]); i++ {
			if inp[j][i] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("==========")
}

func printByteGrid(inp [][]byte) {
	fmt.Println("==========")
	for j := 0; j < len(inp); j++ {
		for i := 0; i < len(inp[j]); i++ {
			switch inp[j][i] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("F")
			case 3:
				fmt.Print("W")
			}
		}
		fmt.Println()
	}
	fmt.Println("==========")
}

func readInput(fn string) []string {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem22"
	} else {
		fp = filepath.Dir(fp)
	}
	fp += "/" + fn

	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
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
