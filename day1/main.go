package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var input, _ = os.ReadFile("./input2")

func main() {

	var left []int
	var right []int

	input = bytes.ReplaceAll(input, []byte("\n"), []byte(" "))

	fmt.Println("Input", input)

	inputMatrix := bytes.Split(input, []byte(" "))

	fmt.Println("InputMatrix", inputMatrix)

	intList := convInt(inputMatrix)

	fmt.Println("IntList", intList)

	left, right = splitInput(intList)
	//sortNums(left)
	//sortNums(right)
	fmt.Printf("Left:\n %d\n\n", left)
	fmt.Printf("Right:\n %d\n\n", right)

	//fmt.Printf("The distance between the two lists is: %d\n", solve(left, right))
	//fmt.Printf("The similarity score is: %d", calcSimScore(left, right))

}

func convInt(list [][]byte) []int {
	var nums []int
	//var numsList [][]int
	fmt.Printf("List Length: %d\n", len(list))

	for _, i := range list {
		var str string

		if i == nil {
			continue
		}

		for _, j := range i {
			if j == byte(' ') {
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

	for i, j := 0, 1; j < len(nums)-1; {
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

/*
PART 2:

This time, you'll need to figure out exactly how often each number from the left list appears in the right list.
Calculate a total similarity score by adding up each number in the left list
after multiplying it by the number of times that number appears in the right list.

*/

func calcSimScore(nums1, nums2 []int) int {
	// naive solution. could be done with map
	var accum int

	for i := range nums1 {
		var count int
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				count += 1
			} else {
				continue
			}
		}
		if count > 0 {
			accum += count * nums1[i]
		} else {
			accum += nums1[i]
		}
	}
	return accum
}
