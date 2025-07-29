package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.Open("_full.txt")
	check(err)
	scanner := bufio.NewScanner(data)
	matchList := []string{}
	intList := [][]int{}
	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		newLine := [][]string{}
		newLine = append(newLine, match(line))

		for i := 0; i < len(newLine); i++ {
			fmt.Println("newLine i value: ", newLine[i])
			matchList = append(matchList, newLine[i]...)
		}

	}

	for _, val := range matchList {
		fmt.Println("matchList", val)
		intPair := stringClean(val)
		intList = append(intList, intPair)

	}

	fmt.Println("Int List: ", intList)

	for _, ints := range intList {
		product := ints[0] * ints[1]

		answer = answer + product
	}

	fmt.Print("Answer: ", answer)

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

func stringClean(line string) []int {
	fmt.Println("Original Line: ", line)
	intLine := []int{}
	stringLine := strings.Split(line, "mul(")
	fmt.Println("String Line First Pass: ", stringLine)

	for _, val := range stringLine {
		val = strings.ReplaceAll(val, ")", "")
		fmt.Println("String Line Second Pass: ", val)
		stringLine = strings.Split(val, ",")
		fmt.Println("String Line Final Pass: ", stringLine)
	}

	for _, v := range stringLine {
		fmt.Println("String Line Values", v)

		intLine = append(intLine, intConv(v))
	}

	return intLine
}

func intConv(char string) int {
	num, err := strconv.Atoi(char)
	check(err)

	return num
}
