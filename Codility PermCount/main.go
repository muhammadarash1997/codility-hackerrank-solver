package main

import "sort"

func Solution(A []int) int {
    sort.Ints(A)

	if A[0] != 1 {
		return 0
	}

    if len(A) == 1 {
		return 1
    }

    for i := 0; i < len(A)-1; i++ {
        if A[i+1] - A[i] > 1 || A[i+1] - A[i] == 0 {
            return 0
        }
    }

    return 1
}