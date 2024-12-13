package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var input, _ = os.ReadFile("./input")

func main() {

	func() {
		inputMatrix := bytes.Split(input, []byte("\n"))
		//fmt.Printf("InputMatrix: %v", inputMatrix)

		matrix, _ := convertBytesToInts(inputMatrix)
		fmt.Printf("Int Matrix: %v", matrix)

	}()
}

/*func convToInt(arr [][]string) [][]int {

	return nil
}*/

/*func convInt(list [][]byte) [][]int {
	var nums []int
	var matrix [][]int

	for _, i := range list {
		var str string
		if i == nil {
			continue
		}

		for _, j := range i {
			if j == byte(' ') || j == 32 {
				continue
			}
			str += string(j)

			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}

		matrix = append(matrix, nums)
	}
	return matrix
}*/

func convertBytesToInts(byteSlices [][]byte) ([][]int, error) {
	var result [][]int

	for i, slice := range byteSlices {
		for _, b := range slice {
			// Convert byte to string, then to int
			strVal := string(b)
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				return nil, err // Handle error if conversion fails
			}
			result[i] = append(result[i], intVal)
		}
	}

	return result, nil
}
