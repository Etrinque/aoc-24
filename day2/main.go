package main

import (
	"bytes"
	"fmt"
	"os"
)

var input, _ = os.ReadFile("./input")

func main() {

	func() {
		inputMatrix := bytes.Split(input, []byte("\n"))

		for i := 0; i < len(inputMatrix)-1; i++ {
			fmt.Printf("String: %s\n", string(inputMatrix[i]))

		}

	}()
}

func convToInt(arr [][]string) [][]int {

	return nil
}
