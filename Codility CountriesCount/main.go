package main

import "fmt"

func main() {
	// JAWABAN MAP INI ADALAH 8
	// A := [][]int{
	// 	{5, 4, 4, 5},
	// 	{4, 4, 4, 5},
	// 	{4, 5, 4, 5},
	// 	{4, 4, 2, 5},
	// 	{4, 4, 2, 5},
	// 	{4, 4, 3, 5},
	// 	{4, 1, 1, 0},
	// }

	// JAWABAN MAP INI ADALAH 12
	A := [][]int{
		{5, 4, 4},
		{4, 3, 4},
		{3, 2, 4},
		{2, 2, 2},
		{3, 3, 4},
		{1, 4, 4},
		{4, 1, 5},
	}

	// JAWABAN MAP INI ADALAH 15
	// A := [][]int{
	// 	{5, 4, 4, 3, 3},
	// 	{4, 3, 4, 3, 6},
	// 	{3, 2, 4, 6, 6},
	// 	{2, 2, 2, 7, 7},
	// 	{3, 3, 4, 2, 7},
	// 	{1, 4, 4, 2, 2},
	// 	{4, 1, 1, 1, 2},
	// }

	// JAWABAN MAP INI ADALAH 1
	// A := [][]int{
	// 	{5},
	// }

	// JAWABAN MAP INI ADALAH 4
	// A := [][]int{
	// 	{5, 6},
	// 	{7, 8},
	// }

	// JAWABAN MAP INI ADALAH 9
	// A := [][]int{
	// 	{5, 6, 7},
	// 	{7, 8, 9},
	// 	{5, 6, 7},
	// }

	count(A)
}

func count(A [][]int) {
	k := make(map[string]int)
	var count = 0
	for i := 0; i < len(A); i++ {

		for j := 0; j < len(A[i]); j++ {

			left := A[i][j]
			strLeft := fmt.Sprintf("A[%v][%v]", i, j)

			if j < len(A[i])-1 {
				right := A[i][j+1]
				fmt.Println(left, right)
				strRight := fmt.Sprintf("A[%v][%v]", i, j+1)

				// compare left and right]
				if left == right {
					if k[strLeft] == 0 && k[strRight] == 0 {
						count++
						k[strLeft] = count
						k[strRight] = count
					} else if k[strLeft] == 0 && k[strRight] != 0 {
						k[strLeft] = k[strRight]
					} else if k[strLeft] != 0 && k[strRight] == 0 {
						k[strRight] = k[strLeft]
					}
				}
				if left != right {
					if k[strLeft] == 0 && k[strRight] == 0 {
						count++
						k[strLeft] = count
					} else if k[strLeft] == 0 && k[strRight] != 0 {
						count++
						k[strLeft] = count
					} else if i == len(A)-2 && j == len(A[i])-2 && k[strRight] == 0 {	// Masalah tdk kehitung ketika pojok kanan terakhir dan berbeda sendiri, sedang diterapkan pada baris ini yaitu jika pengulangan sudah mencapai akhir dan jika strRight masih 0
						count++
						k[strRight] = count
					}
				}
			}

			if i < len(A)-1 {
				below := A[i+1][j]
				fmt.Println(left, below)
				strBelow := fmt.Sprintf("A[%v][%v]", i+1, j)

				// compare left and below
				if left == below {
					if k[strLeft] == 0 && k[strBelow] == 0 {
						count++
						k[strLeft] = count
						k[strBelow] = count
					} else if k[strLeft] != 0 && k[strBelow] == 0 {
						k[strBelow] = k[strLeft]
					}
				}
				if left != below {
					if k[strLeft] == 0 {
						count++
						k[strLeft] = count
					}
				}
			}

			if len(A) == 1 && len(A[i]) == 1 {
				fmt.Println(1)
				return
			}
		}
	}

	amount := 0
	for i, j := range k {
		fmt.Println(i, j)
		if j > amount {
			amount = j
		}
	}

	fmt.Println(amount)
}

// CHECK WHETHER SAME
// A[0][0] A[0][1]
// A[0][0] A[1][0]

// A[0][1] A[0][2]
// A[0][1] A[1][1]

// A[0][2]
// A[0][2] A[1][2]

// CHECK WHETHER SAME
// A[1][0] A[1][1]
// A[1][0] A[2][0]

// A[1][1] A[1][2]
// A[1][1] A[2][1]

// A[1][2]
// A[1][2] A[2][2]

// CHECK WHETHER SAME
// A[2][0] A[2][1]
// A[2][0] A[3][0]

// A[2][1] A[2][2]
// A[2][1] A[3][1]

// A[2][2]
// A[2][2] A[3][2]