package main

import "fmt"

func main() {
	A := []int{
		10, 2, 5,
		10, 8, 5,
		1, 8, 6,
	}

	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// write your code in Go 1.4

	if len(A) < 5 {
		if len(A) == 0 || len(A) == 1 || len(A) == 2 || len(A) == 3 {

			return 1
		}
		return 0
	}

	loop := (len(A)-5)/3 + 1
	count := 0

	j := 0
	for i := 0; i < loop; i++ {
		p := j
		q := j + 2
		r := j + 4

		if A[p]+A[q] > A[r] && A[q]+A[r] > A[p] && A[r]+A[p] > A[q] {
			count++
		}
		j += 3
	}

	return count
}
