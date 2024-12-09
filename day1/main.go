package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var input, _ = os.ReadFile("./input")

func main() {

	var left []int
	var right []int

	input = bytes.ReplaceAll(input, []byte("\n"), []byte(" "))
	inputMatrix := bytes.Split(input, []byte(" "))

	intList := convInt(inputMatrix)

	left, right = splitInput(intList)

	sortNums(left)
	sortNums(right)

	//fmt.Printf("Left:\n %d\n\n", left)
	//fmt.Printf("Right:\n %d\n\n", right)
	fmt.Printf("The distance between the two lists is: %d\n", solve(left, right))
}

func convInt(list [][]byte) []int {
	var nums []int

	for i := range list {
		var str string

		for _, j := range list[i] {
			if j == ' ' {
				continue
			}
			str += string(j)

		}
		num, _ := strconv.Atoi(str)
		if num == 0 {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func splitInput(nums []int) ([]int, []int) {
	var nums1, nums2 []int

	for i, j := 0, 1; j < len(nums); {
		nums1 = append(nums1, nums[i])
		nums2 = append(nums2, nums[j])
		i += 2
		j += 2

	}

	return nums1, nums2
}

func sortNums(nums []int) []int {
	slices.Sort(nums)
	return nums
}

func solve(nums1, nums2 []int) int {
	var accum int

	for i := range nums1 {

		a := nums1[i]
		b := nums2[i]
		if a == b {
			continue
		} else {
			accum += diff(a, b)
		}
	}
	return accum
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
