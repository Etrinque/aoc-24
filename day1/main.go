package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const url = "https://adventofcode.com/2024/day/1/input"

func main() {
	fmt.Println("Hello, AOC!")
	var client = http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failure to make request: %v", err)
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}
	fmt.Printf("Buffer length: %d", len(buf))

	var accum int

	for i, j := 0, 1; j < len(buf); {

		var a, _ = strconv.Atoi(string(buf[i]))
		var b, _ = strconv.Atoi(string(buf[j]))

		fmt.Printf("A: %d", a)
		fmt.Printf("B: %d", b)

		accum += diff(a, b)
		i++
		j++
	}

	fmt.Println(accum)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
