package problem24

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"strings"
	"strconv"
)

type testCase struct {
	input  string
	result int
}

type portType struct {
	plug0 int
	plug1 int
	used  bool
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", result: 31},
	{input: "input1.txt", result: 1940},
}

var testCasesPartTwo = []testCase{
	{input: "input0.txt", result: 19},
	{input: "input1.txt", result: 1928},
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

var ports []portType
var maxAccum int
var maxLength int

func (tc *testCase) solvePartOne() int {
	ports = parsePortTypes(readInput(tc.input))
	maxAccum = -1
	traversePortsPartOne(0, 0)
	return maxAccum
}

func traversePortsPartOne(plug, accum int) {
	if maxAccum < accum {
		maxAccum = accum
	}
	for i, port := range ports {
		if port.used {
			continue
		}
		if port.plug0 == plug {
			ports[i].used = true
			traversePortsPartOne(port.plug1, accum+port.plug0+port.plug1)
			ports[i].used = false
		}
		if port.plug1 == plug {
			ports[i].used = true
			traversePortsPartOne(port.plug0, accum+port.plug0+port.plug1)
			ports[i].used = false
		}
	}
}

func (tc *testCase) solvePartTwo() int {
	ports = parsePortTypes(readInput(tc.input))
	maxAccum, maxLength = -1, -1
	traversePortsPartTwo(0, 0, 0)
	return maxAccum
}

func traversePortsPartTwo(plug, accum, length int) {
	if maxAccum < accum && maxLength <= length {
		maxAccum = accum
		maxLength = length
	}
	for i, port := range ports {
		if port.used {
			continue
		}
		if port.plug0 == plug {
			ports[i].used = true
			traversePortsPartTwo(port.plug1, accum+port.plug0+port.plug1, length+1)
			ports[i].used = false
		}
		if port.plug1 == plug {
			ports[i].used = true
			traversePortsPartTwo(port.plug0, accum+port.plug0+port.plug1, length+1)
			ports[i].used = false
		}
	}
}

func parsePortTypes(inputData []string) []portType {
	result := []portType{}
	for _, inputItem := range inputData {
		parts := strings.Split(inputItem, "/")
		plug0, _ := strconv.Atoi(parts[0])
		plug1, _ := strconv.Atoi(parts[1])
		result = append(result, portType{plug0: plug0, plug1: plug1, used: false})
	}
	return result
}

func readInput(fn string) []string {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem24"
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
