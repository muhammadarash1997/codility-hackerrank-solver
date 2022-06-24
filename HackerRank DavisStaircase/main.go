package main

import "fmt"

func main() {
	var count int
	var n int = 5
	var limit int = n
	var sum int = 0
	calculate(n, sum, limit, &count)
	fmt.Println(count)
}

func calculate(n int, sum int, limit int, count *int) {
	if limit == 0 {
		return
	}

	for i := 1; i <= 3; i++ {
		if i+sum == n {
			*count++
			return
		}

		calculate(n, i+sum, limit-1, count)
	}
}

// 1 = 1
// 2 = 2
// 3 = 4
// 4 = 7
// 5 = 13
// 6 = 24
// 7 = 44