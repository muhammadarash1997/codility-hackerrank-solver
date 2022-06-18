package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Data Testing
	DataTesting := []int{561892, 1029, 320}

	for _, M := range DataTesting {
		Solution(M)
	}
}

func Solution(M int) {
	N := int64(M)
	binary := strconv.FormatInt(N, 2)
	fmt.Println(binary)
	var start = false
	var mostGap = 0
	var count = 0
	for _, d := range binary {
		b, _ := strconv.Atoi(string(d))
		if b == 1 && count == 0 {
			start = true
		}
		if b == 1 && count > 0 {
			if count >= mostGap {
				mostGap = count
				count = 0
			}
			if count < mostGap {
				count = 0
			}
		}
		if b == 0 && start {
			count++
		}
	}

	fmt.Println(mostGap)
}