package main

import (
	"fmt"
	"sort"
)

func FindSmallest(slice []int) int {
	if slice[len(slice)-1] < 0 {
		return 1
	}

	for i, d := range slice {
		if 1 <= d && d <= slice[len(slice)-1] {
			if i < len(slice)-1 {
				if slice[i+1] - d > 1 {
					return d + 1
				}
			} else if i == len(slice)-1 {
				return d + 1
			}
		}
	}

	return -1
}

func main() {
	// Data Testing
	DataTesting := [][]int{
		{-3, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5},
		{1, 2, 3},
		{-1, -2, -3},
	}

	for _, slice := range DataTesting {
		Solution(slice)
	}
}

func Solution(slice []int) {
	sort.Ints(slice)

	smallest := FindSmallest(slice)

	fmt.Println(smallest)
}