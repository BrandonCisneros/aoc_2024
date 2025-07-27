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

	data, err := os.Open("_sample2.txt")
	check(err)
	scanner := bufio.NewScanner(data)

	count := 0
	lineCount := 1
	badList := [][]int{}
	badCount := 0

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
			} else {
				badList = append(badList, numList)
			}
		case diff > 0:
			verdict := desc(numList)
			if verdict {
				count++
			} else {
				badList = append(badList, numList)
			}
		case diff == 0:
			if numList[1] < numList[2] {
				verdict := asc(numList)
				if verdict {
					count++
				} else {
					badList = append(badList, numList)
				}
			} else if numList[1] > numList[2] {
				verdict := desc(numList)
				if verdict {
					count++
				} else {
					badList = append(badList, numList)
				}
			}
		}

	}

	fmt.Println("Processing Bad List...")
	for i := 0; i < len(badList); i++ {
		fmt.Println("Bad List Line: ", badList[i])
	}

	if badFunc(badList) {
		badCount++
	}

	fmt.Println("Normal Safe Count: ", count)
	fmt.Println("Bad List Safe Count: ", badCount)

	fmt.Println("Final Count: ", count+badCount)
}

func asc(line []int) bool {
	fmt.Println("Entered ASC function...")

	safe := true
	valList := line

	//println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		fmt.Println("List: ", valList)

		diff := valList[i+1] - valList[i]

		if valList[i] < valList[i+1] && diff >= 1 && diff <= 3 {
			//println(valList[i], " | ", valList[i+1])
			continue
		} else {
			safe = false
		}

	}

	fmt.Println("asc safe: ", safe)
	return safe

}

func desc(line []int) bool {
	fmt.Println("Entered DESC function...")
	safe := true
	valList := line

	//println("Line: ", line, " | valList: ", valList)

	for i := 0; i < int(len(valList)-1); i++ {
		//fmt.Println("DESC Index Value: ", valList[i])

		diff := valList[i] - valList[i+1]
		//println("Diff: ", diff)

		if valList[i] > valList[i+1] && diff >= 1 && diff <= 3 {
			continue
		} else {
			safe = false
		}

	}

	println("desc safe:", safe)
	return safe
}

func badFunc(list [][]int) (safe bool) {
	safe = false
	//count := 0

	for i := 0; i < len(list); i++ {
		fmt.Println("The current value is", list[i])
		line := list[i]
		lineCopy := make([]int, len(line))
		copy(lineCopy, line)
		lineWrapper := append(lineCopy[:i], lineCopy[i+1:]...)
		fmt.Println("And after it is removed, we get", lineWrapper)
	}

	return safe
}
