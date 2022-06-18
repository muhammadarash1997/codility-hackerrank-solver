package main

import (
	"fmt"
)

func main() {
	// Data Testing
	DataTesting := [][]int{
		{1, 3, 3, 3, 5, 7, 9, 7, 9, 1, 5},
		{1, 3, 1, 5, 5, 3, 7},
	}

	for _, A := range DataTesting {
		Solution(A)
	}
}

func Solution(A []int) {
	keys := make(map[int]int)
	for i := 0; i < len(A); i++ {
		amount := keys[A[i]]
		amount++
		keys[A[i]] = amount
	}

	for i, d := range keys {
		if d%2 == 1 {
			fmt.Println(i)
		}
	}
}