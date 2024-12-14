package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var input, _ = os.ReadFile("./input")

func main() {

	inputMatrix := bytes.Split(input, []byte("\n"))

	var safeCount int
	var safe bool

	for i := 0; i < len(inputMatrix); i++ {
		var integer []int

		//fmt.Printf("InputMatrix: %v\n", inputMatrix[i])
		str := bytes.Split(inputMatrix[i], []byte(" "))

		for j := 0; j < len(str); j++ {

			num, _ := strconv.Atoi(string(str[j]))
			integer = append(integer, num)

		}

		if len(integer) < 1 {
			continue
		}

		if isAscend(integer) || isDescend(integer) {
			if isSafe(integer) {
				safeCount++
				safe = true
			}
		}

		//fmt.Printf("String: %v\n", str)
		fmt.Printf("Integer: %v\n", integer)
		fmt.Printf("SafeCount: %v\n", safeCount)
		fmt.Printf("Safe: %v\n", safe)
	}
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

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
