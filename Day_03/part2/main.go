package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, _ := os.Open("_full.txt")
	input, err := io.ReadAll(data)
	check(err)
	//scanner := bufio.NewScanner(data)
	answer := 0

	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	/*for scanner.Scan() {
		line := scanner.Text()

		matches := pattern.FindAllStringSubmatch(line, -1)

		fmt.Println(matches)

		do := true

		for _, match := range matches {

			if match[0] == "do()" {
				do = true
				continue
			} else if match[0] == "don't()" {
				do = false
				continue
			} else if !do {
				continue
			}

			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			answer += num1 * num2
		}

	}*/

	matches := pattern.FindAllStringSubmatch(string(input), -1)

	fmt.Println(matches)

	do := true

	for _, match := range matches {

		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		} else if !do {
			continue
		}

		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		answer += num1 * num2
	}

	fmt.Print("Answer: ", answer)

}
