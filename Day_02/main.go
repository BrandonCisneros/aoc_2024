package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		println("There was an error!")
		panic(e)
	}
}

func main() {

	data, err := os.Open("_sample.txt")
	check(err)
	scanner := bufio.NewScanner(data)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		diff := int(line[0]) - int(line[len(line)-1])

		fmt.Println(diff)

		switch diff < 0 {
		case true:
			verdict := asc(line)
			fmt.Println(verdict)
			if verdict == true {
				count++
			}
		case false:
			verdict := desc(line)
			fmt.Println(verdict)
			if verdict == true {
				count++
			}
		}

	}

	fmt.Println("Count: ", count)
}

func asc(line string) bool {
	fmt.Println("Entered ASC function...")
	safe := false

	for i := 0; i < int(len(line)-1); i++ {
		fmt.Println("ASC i: ", i)

	}

	fmt.Println("asc safe: ", safe)
	return safe

}

func desc(line string) bool {
	fmt.Println("Entered DESC function...")
	safe := false

	for i := 0; i < int(len(line)-1); i++ {
		fmt.Println("DESC i: ", i)

	}

	return safe
}
