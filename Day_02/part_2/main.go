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
	reports := [][]int{}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		repList := []int{}

		valList := strings.Split(line, " ")

		for i := range valList {
			val, err := strconv.Atoi(valList[i])
			repList = append(repList, val)
			check(err)
		}

		reports = append(reports, repList)
	}

	fmt.Println("Reports: ", reports)

	for _, data := range reports {
		fmt.Println("Reports Data: ", data)

		//--- This works because the if statement will
		//--- only evaluate true cases; Anything that
		//--- doesn't meet base case will not increment count
		if validate(data) || dampener(data) {
			count++
		}
	}

	fmt.Println("Count: ", count)
}

func validate(data []int) bool {
	incr := allIncreasing(data)
	decr := allDecreasing(data)
	min := minDiff(data, 1)
	max := maxDiff(data, 3)

	return (incr || decr) && min && max
}

func dampener(data []int) bool {
	//--- This approach is creating a different slice each loop
	//--- by removing the index value from the orginal slice.
	//--- This a brute force method.
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if validate(data1) {
			return true
		}
	}

	return false
}

func allIncreasing(data []int) bool {
	//--- The approach of the assumption is true unless it's proven no to be.
	//--- The loop approach keeps everything within the index bounds
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}

	return true
}

func allDecreasing(data []int) bool {
	//--- The approach of the assumption is true unless it's proven no to be.
	//--- The loop approach keeps everything within the index bounds
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			return false
		}
	}

	return true
}

func minDiff(data []int, min int) bool {
	//--- The assumtion is true unless proven false.
	//--- we send the int to an absolute value function
	//--- so that we can make negatives absolute without
	//--- converting the int to a float using the bult-in absolute value function.

	for i := 1; i < len(data); i++ {
		if absVal(data[i]-data[i-1]) < min {
			return false
		}
	}
	return true
}

func maxDiff(data []int, max int) bool {
	//--- The assumtion is true unless proven false.
	//--- we send the int to an absolute value function
	//--- so that we can make negatives absolute without
	//--- converting the int to a float using the bult-in absolute value function.

	for i := 1; i < len(data); i++ {
		if absVal(data[i]-data[i-1]) > max {
			return false
		}
	}

	return true
}

func absVal(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}
