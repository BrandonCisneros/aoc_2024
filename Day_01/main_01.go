/*
1. Sort each list from smallest to largest
2. get the absolute-value difference between the respective line items
3. Sum the differences
4.



*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	valList1, valList2 := parse()
	valList1, valList2 = sorter(valList1, valList2)

	fmt.Println("List 1: ", valList1)
	fmt.Println("List 2: ", valList2)

}

func parse() ([]int, []int) {

	a := 0

	valList1 := []int{}
	valList2 := []int{}

	data, err := os.Open("./_sample.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		//fmt.Println(line)

		listLine := strings.Split(line, "   ")

		fmt.Println("A Value: ", a)
		fmt.Println("0 Value: ", listLine[0])
		fmt.Println("1 Value: ", listLine[1])

		fmt.Println("valList Position: ", valList1[1])

		valList1[a], err = strconv.Atoi(listLine[0])
		valList2[a], err = strconv.Atoi(listLine[1])
		check(err)

		a++

	}

	return valList1, valList2
}

func sorter(valList1 []int, valList2 []int) ([]int, []int) {
	list := [4]int{1, 2, 45, 5}

	fmt.Println(slices.Sort(list))

	return valList1, valList2
}
