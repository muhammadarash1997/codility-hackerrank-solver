package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{10, 0, 8, 2, -1, 12, 11, 3}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)

	distanceList := []int{}

	for i := 0; i < len(A)-1; i++ {
		distanceList = append(distanceList, A[i+1]-A[i])
	}

	sort.Ints(distanceList)

	distanceList[len(distanceList)-1] = distanceList[len(distanceList)-1]/2

	sort.Ints(distanceList)

	return distanceList[len(distanceList)-1]
}