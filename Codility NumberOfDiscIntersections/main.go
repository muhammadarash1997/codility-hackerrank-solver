package main

import "fmt"

func main() {
	// Data Testing
	DataTesting := [][]int{
		{1, 5, 2, 1, 4, 0},
		{1, 5, 2, 1, 4, 0, 1},
		{1, 1, 1},
	}

	for _, A := range DataTesting {
		Solution(A)
	}
}

func Solution(A []int) {
	// write your code in Go 1.4

	var Count int

	for i := 0; i < len(A); i++ {

		CenterA := i
		LeftA := i - A[i]
		RightA := i + A[i]

		for j := i + 1; j < len(A); j++ {

			CenterB := j
			LeftB := j - A[j]
			RightB := j + A[j]

			// (A && B) || (C && D) || (E && C) || (B && F)
			if LeftA <= CenterB && CenterB <= CenterA || CenterA <= CenterB && CenterB <= RightA {
				Count++
				continue
			}
			if LeftB <= CenterA && CenterA <= CenterB || CenterB <= CenterA && CenterA <= RightB {
				Count++
				continue
			}
			if LeftA == RightB || RightA == LeftB {
				Count++
				continue
			}
		}
	}

	fmt.Println(Count)
}

// func Solution(A []int) {
// 	// write your code in Go 1.4

// 	var Count int

// 	for i := 0; i < len(A); i++ {

// 		CenterA := i
// 		LeftA := i - A[i]
// 		RightA := i + A[i]

// 		for j := i + 1; j < len(A); j++ {

// 			CenterB := j
// 			LeftB := j - A[j]
// 			RightB := j + A[j]

// 			if LeftA <= CenterB && CenterB <= CenterA || CenterA <= CenterB && CenterB <= RightA {
// 				Count++
// 				continue
// 			}
// 			if LeftB <= CenterA && CenterA <= CenterB || CenterB <= CenterA && CenterA <= RightB {
// 				Count++
// 				continue
// 			}
// 			if LeftA == RightB || RightA == LeftB {
// 				Count++
// 				continue
// 			}
// 		}
// 	}

// 	fmt.Println(Count)
// }
