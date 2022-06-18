package main

import "fmt"

func main() {
	// Data Testing
	DataTesting := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{-1, 6, -4, 7, 10, 9},
		{10, 0, 20, 15, 5},
	}

	// Number of Rotating
	DataNumberOfRotating := []int{1, 2, 3, 4, 5, 6, 7,}

	for _, A := range DataTesting {
		for _, K := range DataNumberOfRotating {
			Solution(A, K)
		}
	}
}

func Solution(A []int, K int) {
	for i := 0; i < K; i++ {
		A = Rotate(A)
	}
	fmt.Println(A)
}

func Rotate(s []int) []int {
	last := s[len(s)-1]

	afterRemove := removeIndex(s, len(s)-1)
	afterAppend := appendToFirst(last, afterRemove)
	return afterAppend
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func appendToFirst(f int, s []int) []int {
	l := []int{}
	l = append(l, f)
	l = append(l, s...)
	return l
}
