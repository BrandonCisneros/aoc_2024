package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 7, 16, 22, 17, 42}
	fmt.Println("We will start out with", x)

	for i := 0; i < len(x); i++ {
		fmt.Println("The current value is", x[i])
		x2 := x
		xCopy := make([]int, len(x2))
		copy(xCopy, x2)
		xWrapper := append(xCopy[:i], xCopy[i+1:]...)
		fmt.Println("And after it is removed, we get", xWrapper)
	}
}
