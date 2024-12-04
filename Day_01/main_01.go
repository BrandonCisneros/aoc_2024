/*
1. Sort each list from smallest to largest
2. get the absolute-value difference between the respective line items
3. Sum the differences
4.



*/

package main

import (
	"bufio"
	"os"
)

type List []int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

}

func parse() (List, List){

	list_1 := List{}
	list_2 := List{}

	data, err := os.Open("./_sample.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	for scanner.Scan(){
		line := scanner.Text()

		fmt.printf()



	}

	return list_1, list_2
}
