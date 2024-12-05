package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// URL for input, input is 1_000 lines of INT pairs.
// Find the difference from col1 - col2, subtract larger from smaller
const url = "https://adventofcode.com/2024/day/1/input"

var input, _ = os.ReadFile("./input")

func main() {

	//var client = http.DefaultClient
	//resp, err := client.Get(url)
	//if err != nil {
	//	log.Fatalf("failure in GET: %v", err)
	//}
	//defer resp.Body.Close()
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalf("failure in reading body: %v", err)
	//}

	var buf []string

	func() {
		input = bytes.ReplaceAll(input, []byte("\n"), []byte(""))
		for i := range input {
			buf = append(buf, string(input[i]))
		}
	}()

	fmt.Println(solve(buf))

}

func solve(data []string) int {
	fmt.Println(data)
	var accum int

	for i := 0; i < len(data)-1; {

		a, _ := strconv.Atoi(data[i])
		b, _ := strconv.Atoi(data[i+1])
		//fmt.Printf("A: %d, B: %d\n", a, b)
		accum += diff(a, b)
		i += 2
	}

	return accum
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
