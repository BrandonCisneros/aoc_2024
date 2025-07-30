package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	xmasSlice := [][]string{}

	origList := parse()

	for i := range origList {
		jSlice := []string{}
		for _, j := range origList[i] {
			jSlice = append(jSlice, string(j))
		}

		xmasSlice = append(xmasSlice, jSlice)
	}

	answer := processing(xmasSlice)

	fmt.Println("Answer: ", answer)

}

func parse() []string {

	dataSlice := []string{}

	data, dataErr := os.Open("_sample.txt")
	check(dataErr)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		dataSlice = append(dataSlice, line)
	}

	return dataSlice

}

// This word search allows words to be
// horizontal --- left & right
// backwards
// vertical --- up & down
// diagonal,
// overlapping other words

func processing(xmasSlice [][]string) int {
	count := 0
	mas := "MAS"
	matrixSize := len(xmasSlice)   //--- sets the row (vertical) size for the matrix
	sliceSize := len(xmasSlice[0]) //--- sets the element (horizontal) size of the matrix

	fmt.Println(mas, " | ", matrixSize, " | ", sliceSize)

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < sliceSize; j++ {

			if xmasSlice[i][j] == "A" {
				count += search(xmasSlice, i, j, mas)
			}

			//count += search(xmasSlice, i, j, mas)

		}
	}

	return count
}

func search(grid [][]string, row int, column int, mas string) int {
	mSize := len(grid)
	sSize := len(grid)
	count := 0

	if grid[row][column] != string(mas[1]) {
		return 0
	}

	masLen := len(mas)

	// Search coordinates to search all 8 directions starting from left to right
	xCoords := []int{-1, -1, 1, 1}
	yCoords := []int{-1, 1, -1, 1}

	//This loop will search in all 8 directions
	// one by one. It will return true if one of
	// the directions contain the word.
	for dir := 0; dir < 4; dir++ {

		//initialize the starting point for current direction
		charIndex := row + xCoords[dir]
		currX := row + xCoords[dir]
		currY := column + yCoords[dir]

		//First character is already checked, match remaining characters
		for charIndex = 1; charIndex < masLen; charIndex++ {

			//Break if out of bounds
			if currX >= mSize || currX < 0 || currY >= sSize || currY < 0 {
				break
			}

			if grid[currX][currY] != string(mas[charIndex]) {
				break
			}

			// Moving in particular direction
			currX += xCoords[dir]
			currY += yCoords[dir]
		}

		if charIndex == masLen {
			count++
		}
	}

	return count
}
