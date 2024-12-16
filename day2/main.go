package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var input, _ = os.ReadFile("./input2")
var dampener = flag.Bool("dampener", false, "dampener active")

func main() {
	flag.Parse()
	inputMatrix := bytes.Split(input, []byte("\n"))

	var safeCount int
	var safe bool

	for i := 0; i < len(inputMatrix); i++ {
		var integer []int

		str := bytes.Split(inputMatrix[i], []byte(" "))

		for j := 0; j < len(str); j++ {

			num, _ := strconv.Atoi(string(str[j]))
			integer = append(integer, num)

		}

		if len(integer) < 1 {
			continue
		}

		if isAscend(integer) || isDescend(integer) && isSafe(integer) {
			safeCount++
			safe = true
		} else if dampenerTolerance(integer) && isSafe(integer) || isDescend(integer) {
			safeCount++
			safe = true
		} else {
			safe = false
		}

		//fmt.Printf("Integer: %v\n", integer)
		fmt.Printf("List %d: | isSafe: %v\n", i, safe)
	}
	fmt.Printf("SafeCount: %v\n", safeCount)
}

/*
The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
*/

// Helper Functions
func isAscend(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func isDescend(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] < list[i+1] {
			return false
		}
	}
	return true
}

func isSafe(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if diff(list[i], list[i+1]) > 3 || diff(list[i], list[i+1]) < 1 {
			return false
		}
	}
	return true
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

/*
part 2
if removing a single level from an unsafe report would make it safe,
the report instead counts as safe.
*/

func dampenerTolerance(list []int) bool {
	if !*dampener {
		return false
	}
	var nums []int

	// need to flag or count if index was dropped and list is now safe
	// only 1 allowed index drop to validate safety else returns unsafe

	//var count int
	//fmt.Println("List:", list)
	for i := 0; i < len(list); i++ {

		nums = append(nums, list[:i]...)
		//fmt.Printf("Nums pre: %d", nums)

		nums = append(nums, list[i+1:]...)
		//fmt.Printf("Nums post: %d", nums)

		//if len(nums) == len(list)-1 {
		//	count += 1
		//}

		if isSafe(nums) {
			return true
		}
	}
	return false
}

// First Run : 649 Too High

// Second Run : 304 Too Low

// Third Run: 326 Too Low

// Fourth Run : 381 incorrect +/-?

// Firth Run : 329 incorrect +/-?
