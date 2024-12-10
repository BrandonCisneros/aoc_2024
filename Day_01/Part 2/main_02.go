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
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("There was an error!")
		panic(e)
	}
}

func main() {

	totalDist := 0

	valList1, valList2 := parse()
	valList1, valList2 = count(valList1, valList2)

	fmt.Println(totalDist)

}

func parse() ([]int, []int) {

	valList1 := []int{}
	valList2 := []int{}

	data, err := os.Open("./_full-dataset.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		listLine := strings.Split(line, "   ")

		num0, err := strconv.Atoi(listLine[0])
		check(err)
		num1, err := strconv.Atoi(listLine[1])
		check(err)

		valList1 = append(valList1, num0)
		valList2 = append(valList2, num1)

	}

	return valList1, valList2
}

func count(valList1 []int, valList2 []int) ([]int, []int) {

	product := 0

	for _, v := range valList1 {
		counter := 0

		for _, z := range valList2 {

			if v == z {
				counter++
			}
			fmt.Println("V = ", v, " | Z = ", z, " | Counter = ", counter)
		}

		product = product + (v * counter)
	}

	fmt.Println("Product: ", product)

	return valList1, valList2
}
