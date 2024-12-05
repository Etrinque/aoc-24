package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var input, _ = os.ReadFile("./input")

func main() {

	//mod to work with delimited ints, use \n as slice delimiter
	var arr []int

	func() {
		inputMatrix := bytes.Split(input, []byte("\n"))
		for i := 0; i < len(inputMatrix); i++ {
			for j := 0; j < len(inputMatrix[i]); j++ {
				if inputMatrix[i][j] == ' ' {
					continue
				}
				num, _ := strconv.Atoi(string(inputMatrix[i][j]))
				arr = append(arr, num)
			}
		}
		fmt.Println(arr)
	}()
}
