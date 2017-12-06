package problem6

import (
	"fmt"
	"strconv"
)

type testCase struct {
	input  []int
	result int
}

var testCasesPartOne = []testCase{
	{input: []int{0, 2, 7, 0}, result: 5},
	{input: []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}, result: -1},
}

var testCasesPartTwo = []testCase{
	{input: []int{0, 2, 7, 0}, result: 4},
	{input: []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}, result: -1},
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 30 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Printf("result    : %d\n\n", v.solvePartOne())
	}
}

func RunPartTwo() {
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 30 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %#v\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Printf("result    : %d\n\n", v.solvePartTwo())
	}
}

func (tc *testCase) solvePartOne() int {
	arr := make([]int, len(tc.input))
	copy(arr, tc.input)
	arrLen := len(arr)

	cache := make([]string, 0)

	steps, result := 1, 0

Loop:
	for /*ii := 0; ii < 5; ii++*/ {
		//fmt.Printf("%#v\n", arr)

		maxVal, maxPos := findMax(arr)

		arr[maxPos] = 0
		chunks := divideSlice(maxVal, arrLen)
		//fmt.Printf("%#v\n", chunks)
		//fmt.Printf("-------------------\n")
		for chK, chV := range chunks {
			effectivePos := (maxPos + 1 + chK) % arrLen
			arr[effectivePos] += chV
		}

		cacheKey := buildKey(arr)

		for _, tmpV := range cache {
			if tmpV == cacheKey {
				result = steps
				break Loop
			}
		}

		cache = append(cache, cacheKey)
		steps++
	}

	//fmt.Println(len(cache))
	return result
}

func (tc *testCase) solvePartTwo() int {
	arr := make([]int, len(tc.input))
	copy(arr, tc.input)
	arrLen := len(arr)

	cache := make([]string, 0)

	steps, result := 1, 0

Loop:
	for /*ii := 0; ii < 5; ii++*/ {
		//fmt.Printf("%#v\n", arr)

		maxVal, maxPos := findMax(arr)

		arr[maxPos] = 0
		chunks := divideSlice(maxVal, arrLen)
		//fmt.Printf("%#v\n", chunks)
		//fmt.Printf("-------------------\n")
		for chK, chV := range chunks {
			effectivePos := (maxPos + 1 + chK) % arrLen
			arr[effectivePos] += chV
		}

		cacheKey := buildKey(arr)

		for tmpK, tmpV := range cache {
			if tmpV == cacheKey {
				result = steps - tmpK - 1
				break Loop
			}
		}

		cache = append(cache, cacheKey)
		steps++
	}

	//fmt.Println(len(cache))
	return result
}

func findMax(arr []int) (val, pos int) {
	val, pos = -1, -1
	for i, k := 0, len(arr)-1; i < len(arr); i, k = i+1, k-1 {
		v := arr[k]
		if val <= v {
			val, pos = v, k
		}
	}
	return
}

func divideSlice(val, parts int) []int {
	chunks := make([]int, parts)
	for i := 0; val > 0; i++ {
		chunks[i % parts] += 1
		val--
	}
	return chunks

	/*chunks := []int{}
	chunkSize := (val + parts - 1) / parts
	for i := 0; i < val; i += chunkSize {
		end := i + chunkSize
		if end > val {
			end = val
		}
		chunks = append(chunks, end-i)
	}
	return chunks*/
}

func buildKey(arr []int) (cacheKey string) {
	for _, arrV := range arr {
		cacheKey += strconv.Itoa(arrV) + ","
	}
	return
}
