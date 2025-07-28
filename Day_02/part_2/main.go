package main

import (
	"bufio"
	"fmt"
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

	data, err := os.Open("_full.txt")
	check(err)
	scanner := bufio.NewScanner(data)

	count := 0
	lineCount := 1

	for scanner.Scan() {
		line := scanner.Text()

		valList := strings.Split(line, " ")
		numList := []int{}

		for _, i := range valList {
			i, err := strconv.Atoi(i)
			check(err)
			numList = append(numList, i)
		}

		diff := numList[0] - numList[len(numList)-1]

		println("Line: ", lineCount)
		lineCount++

		fmt.Println("Num1: ", numList[0], " | Num2: ", numList[len(numList)-1], " | Diff: ", diff)

		switch {
		case diff < 0:
			verdict := asc(numList)
			if verdict {
				count++
			}
		case diff > 0:
			verdict := desc(numList)
			if verdict {
				count++
			}
		}

	}

	fmt.Println("Count: ", count)
}

func asc(line []int) bool {
	println("Entered ASC function...")

	safe := true
	valList := line
	badLevels := 0

	fmt.Println("Line: ", line, " | valList: ", valList)

	for i := 1; i < int(len(valList)-1); i++ {
		fmt.Println("ASC i: ", valList[i])

		diff := valList[i] - valList[i-1]
		println("Diff: ", diff)

		if valList[i] < valList[i+1] && diff >= 1 && diff <= 3 {
			println(valList[i], " | ", valList[i+1])
			continue
		} else {
			fmt.Println("Processing Bad Lists...")
			for i := 1; i < len(valList)-1; i++ {
				line := valList
				copyLine := make([]int, len(line))
				copy(copyLine, line)
				lineWrapper := append(copyLine[:i], copyLine[i+1:]...)
				fmt.Println("Line Wrapper: ", lineWrapper)

				diff := lineWrapper[i] - lineWrapper[i-1]
				println("Diff: ", diff)

				if lineWrapper[i] < lineWrapper[i+1] && diff >= 1 && diff <= 3 {
					println(lineWrapper[i], " | ", lineWrapper[i+1])
					continue
				} else {
					badLevels++
				}

			}
		}
	}

	fmt.Println("Bad Levels: ", badLevels)
	if badLevels > 1 {
		safe = false
	}

	fmt.Println("asc safe: ", safe)
	return safe

}

func desc(line []int) bool {
	fmt.Println("Entered DESC function...")
	safe := true

	valList := line

	fmt.Println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		diff := valList[i] - valList[i+1]
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
