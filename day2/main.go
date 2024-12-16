package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var input, _ = os.ReadFile("./input")
var dampener = flag.Bool("dampener", false, "dampener active")

func main() {
	flag.Parse()
	inputMatrix := bytes.Split(input, []byte("\n"))

	var safeCount int

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

		if isAscend(integer) && isSafe(integer) || isDescend(integer) && isSafe(integer) {
			safeCount++
		} else if dampenerTolerance(i, integer) {
			safeCount++
		}

	}
	fmt.Printf("SafeCount: %v\n", safeCount)
}

/*
Check sorted order, The levels are either all increasing or all decreasing.
*/

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

/*
Any two adjacent levels differ by at least one and at most three.
*/

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
Boolean Flag activated from command line with arg  { -dampener=true }
*/

func dampenerTolerance(idx int, list []int) bool {
	if !*dampener {
		return false
	}
	var nums []int

	for i := 0; i < len(list); i++ {

		nums = slices.Concat(list[:i], list[i+1:])

		if isAscend(nums) && isSafe(nums) || isDescend(nums) && isSafe(nums) {
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

// Fifth Run : 524 Incorrect +/-?

// Sixth Run : 515 Incorrect +/-?

//  Seventh Run : 354 CORRECT!
