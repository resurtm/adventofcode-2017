package problem21

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"strings"
	"math"
)

type testCase struct {
	input  string
	iters  int
	result int
}

type convRule struct {
	src [][]byte
	dst [][]byte
}

var testCasesPartOne = []testCase{
	{input: "input0.txt", iters: 2, result: 12},
	{input: "input1.txt", iters: 5, result: -1},
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
	rules := parseRules(readInput(tc.input))

	replace := func(subMat [][]byte) [][]byte {
		for _, rule := range rules {
			if compareMats(rule.src, subMat) {
				return rule.dst
			}
			if compareMats(flipMatrix1(rule.src), subMat) {
				return rule.dst
			}
			if compareMats(flipMatrix2(rule.src), subMat) {
				return rule.dst
			}
		}
		log.Fatal("rule not found, this should not happen")
		return [][]byte{}
	}

	pattern := [][]byte{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}

	for iters := 0; iters < tc.iters; iters++ {
		toGlue := [][][]byte{}
		size := len(pattern[0])
		elemCount := 0

		if size%2 == 0 {
			elemCount = int(size / 2)
			for partsJ := 0; partsJ < elemCount; partsJ++ {
				for partsI := 0; partsI < elemCount; partsI++ {
					subMat := [][]byte{{0, 0}, {0, 0}}
					for subMatJ := 0; subMatJ < 2; subMatJ++ {
						for subMatI := 0; subMatI < 2; subMatI++ {
							subMat[subMatJ][subMatI] = pattern[subMatJ+2*partsJ][subMatI+2*partsI]
						}
					}
					toGlue = append(toGlue, replace(subMat))
					//printMat(subMat)
				}
			}
		} else {
			elemCount = int(size / 3)
			for partsJ := 0; partsJ < elemCount; partsJ++ {
				for partsI := 0; partsI < elemCount; partsI++ {
					subMat := [][]byte{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
					for subMatJ := 0; subMatJ < 3; subMatJ++ {
						for subMatI := 0; subMatI < 3; subMatI++ {
							subMat[subMatJ][subMatI] = pattern[subMatJ+3*partsJ][subMatI+3*partsI]
						}
					}
					toGlue = append(toGlue, replace(subMat))
					//printMat(replace(subMat))
				}
			}
		}

		newDim := int(math.Sqrt(float64(len(toGlue))))
		glueSize := len(toGlue[0][0])
		newSize := newDim * glueSize
		pattern = make([][]byte, newSize)
		for xx := 0; xx < newSize; xx++ {
			pattern[xx] = make([]byte, newSize)
		}

		for toGlueC := 0; toGlueC < len(toGlue); toGlueC++ {
			toGlueI := toGlue[toGlueC]
			posI := toGlueC % newDim
			posJ := int(toGlueC / newDim)
			for j := 0; j < glueSize; j++ {
				for i := 0; i < glueSize; i++ {
					pattern[posJ*(glueSize+0)+j][posI*(glueSize+0)+i] = toGlueI[j][i]
				}
			}
		}
	}

	stars := 0
	for j := 0; j < len(pattern); j++ {
		for i := 0; i < len(pattern[j]); i++ {
			if pattern[j][i] == '#' {
				stars++
			}
		}
	}
	return stars
}

func (tc *testCase) solvePartTwo() int {
	return 0
}

func printMat(inp [][]byte) {
	fmt.Println("==========")
	for j := 0; j < len(inp); j++ {
		for i := 0; i < len(inp[j]); i++ {
			if inp[j][i] == 0 {
				fmt.Print(string('_'))
			} else {
				fmt.Print(string(inp[j][i]))
			}
		}
		fmt.Println()
	}
	fmt.Println("==========")
}

func compareMats(a, b [][]byte) bool {
	sizeA := len(a[0])
	sizeB := len(b[0])
	if sizeA != sizeB {
		return false
	}
	for j := 0; j < sizeA; j++ {
		for i := 0; i < sizeA; i++ {
			if a[j][i] != b[j][i] {
				return false
			}
		}
	}
	return true
}

func parseRules(inputData []string) []convRule {
	rules := []convRule{}
	for _, inputLine := range inputData {
		var srcS, dstS string
		fmt.Sscanf(inputLine, "%s => %s", &srcS, &dstS)

		srcSM := strings.Split(srcS, "/")
		srcSize := len(srcSM[0])
		srcM := make([][]byte, srcSize)
		for j := 0; j < srcSize; j++ {
			srcM[j] = make([]byte, srcSize)
			for i := 0; i < srcSize; i++ {
				srcM[j][i] = srcSM[j][i]
			}
		}

		dstSM := strings.Split(dstS, "/")
		dstSize := len(dstSM[0])
		dstM := make([][]byte, dstSize)
		for j := 0; j < dstSize; j++ {
			dstM[j] = make([]byte, dstSize)
			for i := 0; i < dstSize; i++ {
				dstM[j][i] = dstSM[j][i]
			}
		}

		for rots := 0; rots < 4; rots++ {
			//rules = append(rules, convRule{flipMatrix1(srcM), dstM})
			//rules = append(rules, convRule{flipMatrix2(srcM), dstM})
			rules = append(rules, convRule{srcM, dstM})
			srcM = rotateMatrix(srcM)
		}
	}
	return rules
}

func rotateMatrix(input [][]byte) [][]byte {
	size := len(input[0])
	result := make([][]byte, size)
	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
	}
	for j, row := range input {
		for i, elem := range row {
			result[i][size-j-1] = elem
		}
	}
	return result
}

func flipMatrix1(input [][]byte) [][]byte {
	size := len(input[0])
	result := make([][]byte, size)
	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
	}
	for j, row := range input {
		for i, elem := range row {
			result[j][size-i-1] = elem
		}
	}
	return result
}

func flipMatrix2(input [][]byte) [][]byte {
	size := len(input[0])
	result := make([][]byte, size)
	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
	}
	for j, row := range input {
		for i, elem := range row {
			result[size-j-1][i] = elem
		}
	}
	return result
}

func readInput(fn string) []string {
	fp := getCurrDir()
	if fp[:5] == "/tmp/" {
		fp = "/home/resurtm/dev/go/src/github.com/resurtm/adventofcode-2017/problem21"
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
