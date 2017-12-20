package problem20

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"math"
)

type testCase struct {
	input  string
	result int64
}

type vector struct {
	x, y, z int64
}

type particle struct {
	p, v, a vector
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", result: 0},
	{input: "input1.txt", result: -1},
}

var testCasesPartTwo = []testCase{
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
	particles := []particle{}
	for _, inputLine := range inputData {
		part := particle{}
		fmt.Sscanf(
			inputLine,
			"p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&part.p.x, &part.p.y, &part.p.z,
			&part.v.x, &part.v.y, &part.v.z,
			&part.a.x, &part.a.y, &part.a.z,
		)
		particles = append(particles, part)
	}

	for i := 0; i < 100000; i++ {
		for i, part := range particles {
			part.v.x += part.a.x
			part.v.y += part.a.y
			part.v.z += part.a.z
			part.p.x += part.v.x
			part.p.y += part.v.y
			part.p.z += part.v.z
			particles[i] = part
		}
	}

	nearest := -1
	minDistance := 0.0
	firstItem := true
	for i, part := range particles {
		distance := math.Sqrt(float64(part.p.x*part.p.x + part.p.y*part.p.y + part.p.z*part.p.z))
		if firstItem || minDistance > distance {
			firstItem = false
			nearest = i
			minDistance = distance
		}
	}

	return nearest
}

func (tc *testCase) solvePartTwo() int {
	return 0
}

func readInput(fn string) []string {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem20"
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
