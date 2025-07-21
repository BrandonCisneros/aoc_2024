package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		println("There was an error!")
		panic(e)
	}
}

func main() {
	//count := 0

	data, err := os.Open("_sample.txt")
	check(err)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		println(line)
	}
}
