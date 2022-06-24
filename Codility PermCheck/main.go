package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{4, 1, 3, 2}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	sort.Ints(A)

	if A[0] != 1 {
		return 0
	}

	if len(A) == 1 {
		return 1
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i+1]-A[i] > 1 || A[i+1]-A[i] == 0 {
			return 0
		}
	}

	return 1
}
