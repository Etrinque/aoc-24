package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers.
	For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024.
	Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted,
there are also many invalid characters that should be ignored,
even if they look like part of a mul instruction.
	Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:
	xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

Only the four highlighted sections are real mul instructions.
	Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).
*/

var input, _ = os.ReadFile("./input")

func main() {
	var result int
	//var pair []int
	//var pairMatrix [][]int

	_ = parseExp(input)

	//fmt.Printf("PairMatrix: %v", pairMatrix)

	//for i := 0; i < len(pairMatrix); i++ {
	//	pair = pairMatrix[i]
	//	result += accumulate(pair[0], pair[1])
	//}

	fmt.Printf("Result: %d\n", result)
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func accumulate(a, b int) int {
	product := a * b
	return product
}

func parseExp(input []byte) [][]int {
	var window []byte
	var matrix [][]int
	var digits string
	var num int
	var pair []int
	var mulCount = 0

	for i := 0; i < len(input)-1; i++ {

		window = input[i : i+3]

		// parse the input for "mul" + "(" + digit/s + "," + digit/s + ")"

		if string(window) == "mul" {
			mulCount++
			fmt.Printf("Found: %s\t", string(window))
			fmt.Println("MulCount\t", mulCount)

			if !isDigit(input[i+4]) {
				continue
			}

			digits += string(input[5])

			//if string(input[i]) == ")" {
			//	matrix = append(matrix, pair)
			//}
		}
	}
	num, _ = strconv.Atoi(digits)
	pair = append(pair, num)
	fmt.Printf("Pair Length: %d\t", len(pair))
	// append to matrix as pairs

	return matrix
}
