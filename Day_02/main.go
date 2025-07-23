package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		println("There was an error!")
		panic(e)
	}
}

func main() {

	data, err := os.Open("_sample2.txt")
	check(err)
	scanner := bufio.NewScanner(data)

	count := 0
	lineCount := 1

	for scanner.Scan() {
		line := scanner.Text()

		valList := strings.Split(line, " ")

		num1, err := strconv.Atoi(valList[0])
		check(err)
		num2, err := strconv.Atoi(valList[len(valList)-1])
		check(err)

		diff := num1 - num2

		println("Line: ", lineCount)
		lineCount++

		//fmt.Println("Num1: ", num1, " | Num2: ", num2, " | Diff: ", diff)

		switch diff < 0 {
		case true:
			verdict := asc(valList)
			if verdict {
				count++
			}
		case false:
			verdict := desc(valList)
			if verdict {
				count++
			}
		}

	}

	fmt.Println("Count: ", count)
}

func asc(line []string) bool {
	//println("Entered ASC function...")
	//println("List: ", line)

	safe := true
	valList := line

	//println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		num1, err := strconv.Atoi(valList[i])
		check(err)
		num2, err := strconv.Atoi(valList[i+1])
		check(err)

		diff := math.Abs(float64(num2 - num1))
		//println("Diff: ", diff)

		if valList[i] < valList[i+1] && (diff >= 1 && diff <= 3) {
			continue
		} else {
			safe = false
		}

	}

	fmt.Println("asc safe: ", safe)
	return safe

}

func desc(line []string) bool {
	//fmt.Println("Entered DESC function...")
	safe := true

	valList := line

	//println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		num1, err := strconv.Atoi(valList[i])
		check(err)
		num2, err := strconv.Atoi(valList[i+1])
		check(err)

		diff := math.Abs(float64(num1 - num2))
		//println("Diff: ", diff)

		if valList[i] > valList[i+1] && diff <= 3 {
			continue
		} else {
			safe = false
		}

	}
	println("desc safe:", safe)
	return safe
}
