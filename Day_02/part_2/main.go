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
	zeroCount := 0
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
		case diff == 0:
			zeroCount++
		}

	}

	fmt.Println("Count: ", count)
	println("zeroCount", zeroCount)
}

func asc(line []int) bool {
	println("Entered ASC function...")

	safe := true
	valList := line

	println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		fmt.Println("ASC i: ", valList[i])

		diff := valList[i+1] - valList[i]
		println("Diff: ", diff)

		if valList[i] < valList[i+1] && diff >= 1 && diff <= 3 {
			println(valList[i], " | ", valList[i+1])
			continue
		} else {
			switch correctionAsc(valList) {
			case true:
				return safe
			case false:
				safe = false
			}
		}

	}

	fmt.Println("asc safe: ", safe)
	return safe

}

func desc(line []int) bool {
	//fmt.Println("Entered DESC function...")
	safe := true

	valList := line

	//println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		diff := valList[i] - valList[i+1]
		//println("Diff: ", diff)

		if valList[i] > valList[i+1] && diff >= 1 && diff <= 3 {
			continue
		} else {
			switch correctionDesc(valList) {
			case true:
				return safe
			case false:
				safe = false
			}
		}

	}
	println("desc safe:", safe)
	return safe
}

func correctionAsc(line []int) bool {
	println("Enterning correction function...")
	safe := true
	errCount := 0

	for i := 0; i < int(len(line)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		diff := line[i+1] - line[i]
		//println("Diff: ", diff)

		if line[i] < line[i+1] && diff >= 1 && diff <= 3 {
			continue
		} else {
			errCount++
			continue
		}
	}

	if errCount > 1 {
		safe = false
	}

	println("Error Count: ", errCount)

	return safe
}

func correctionDesc(line []int) bool {
	println("Enterning correction function...")
	safe := true
	errCount := 0

	for i := 0; i < int(len(line)-1); i++ {
		//fmt.Println("ASC i: ", valList[i])

		diff := line[i] - line[i+1]
		//println("Diff: ", diff)

		if line[i] > line[i+1] && diff >= 1 && diff <= 3 {
			continue
		} else {
			errCount++
			continue
		}
	}

	if errCount > 1 {
		safe = false
	}

	println("Error Count: ", errCount)

	return safe
}
