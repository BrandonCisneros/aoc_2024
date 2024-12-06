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
	"strings"
)

type List [2][]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	valList := parse()

	fmt.Println("List: ", valList)

}

func parse() List {

	var valList List

	data, err := os.Open("./_sample.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		//fmt.Println(line)

		listLine := strings.Split(line, "   ")

		fmt.Println(listLine[0])
		//fmt.Println(listLine1[1])

		for _, v := range listLine[0] {
			fmt.Println("Value: ", string(v))
			a := 1

			valList[1][a] = int(v)

			a++
		}

		for _, z := range listLine[1] {
			fmt.Println("Z Value: ", string(z))
		}

	}

	return valList
}
