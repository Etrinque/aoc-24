package main

import (
	"bytes"
	"fmt"
	"os"
)

var input, _ = os.ReadFile("./input")

func main() {

	//mod to work with delimited int, use \n as slice delimiter
	var arr []int

	//parse slice into string, if == 32 continue else conv to str and concat then append back to array

	func() {
		inputMatrix := bytes.Split(input, []byte("\n"))
		for i := 0; i < len(inputMatrix); i++ {
			cut := bytes.TrimSpace(inputMatrix[i])
			fmt.Println(cut)
			//for j := 0; j < len(inputMatrix[i]); j++ {
			//	if inputMatrix[i][j] == ' ' {
			//		continue
			//	}
			//	num, _ := strconv.Atoi(string(inputMatrix[i][j]))
			//	arr = append(arr, num)
			//}
		}
		fmt.Println(arr)
	}()
}
