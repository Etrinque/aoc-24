package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

var input, _ = os.ReadFile("./input")

func main() {

	var nums []int

	func() {
		input = bytes.ReplaceAll(input, []byte("\n"), []byte(" "))
		inputMatrix := bytes.Split(input, []byte(" "))
		for i := range inputMatrix {
			num, _ := strconv.Atoi(string(inputMatrix[i]))
			if num == 0 {
				continue
			} else {
				nums = append(nums, num)
			}
		}
	}()

	nums1, nums2 := sortNums(splitInput(nums))

	//fmt.Printf("Nums1:\n %d\n\n", nums1)
	//fmt.Printf("Nums2:\n %d\n\n", nums2)
	fmt.Println(solve(nums1, nums2))

}

func splitInput(nums []int) ([]int, []int) {
	var nums1, nums2 []int

	for i := 0; i < len(nums)-2; i += 2 {
		nums1 = append(nums1, nums[i])
		nums2 = append(nums2, nums[i+1])
	}

	return nums1, nums2
}

func sortNums(nums1, nums2 []int) ([]int, []int) {
	slices.Sort(nums1)
	slices.Sort(nums2)
	return nums1, nums2
}

func solve(nums1, nums2 []int) float64 {
	var accum float64

	for i := range nums1 {

		a := nums1[i]
		b := nums2[i]
		if a == b {
			continue
		} else {
			accum += math.Abs(float64(diff(a, b)))
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
