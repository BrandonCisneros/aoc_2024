package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.Open("_sample.txt")
	check(err)
	scanner := bufio.NewScanner(data)
	matchList := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		newLine := [][]string{}
		newLine = append(newLine, match(line))

		fmt.Println("String Match: ", newLine)

		for i := 0; i < len(newLine); i++ {
			fmt.Println("newLine i value: ", newLine[i])
			for _, j := range newLine[i] {
				fmt.Println("newLine j value: ", j)
				matchList = append(matchList, j)
			}
		}

	}

	fmt.Println("matchList", matchList)
}

func match(line string) []string {
	fmt.Println("Line: ", line)
	//mulPattern := `^mul\([0-9]+,[0-9]+\)$`
	mulRegex, err := regexp.Compile(`mul\(\d+,\d+\)`)
	check(err)

	matches := mulRegex.FindAllString(line, -1)

	fmt.Println("Matches: ", matches)
	return matches
}
