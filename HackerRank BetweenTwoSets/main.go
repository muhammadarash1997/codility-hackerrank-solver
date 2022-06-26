package main

import (
	"fmt"
)

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}
	fmt.Println(getTotalX(a, b))
}

func getTotalX(a []int32, b []int32) int32 {

	// First look for a's GCD
	aGCD := findGCD(a...)
	fmt.Println(aGCD)

	// First look for b's GCD
	bGCD := findGCD(b...)
	fmt.Println(bGCD)

	// Collect all number between a's GCD and b's GCD that are the multiple of all of a's elements to list
	list := []int32{}
	for i := aGCD; i <= bGCD; i++ {

		status := 0
		for _, d := range a {

			if i%d == 0 {
				status++
			}
		}
		if status == len(a) {
			list = append(list, i)
		}
	}

	// Calculate modulo each of b's element to each of list's elements
	result := []int32{}
	for _, l := range list {

		status := 0
		for _, d := range b {

			if d%l == 0 {
				status++
			}
		}
		if status == len(b) {
			result = append(result, l)
		}
	}
	return int32(len(result))
}

func findGCD(input ...int32) int32 {

	list := []int{}
	keys := make(map[int]int)
	for _, d := range input {

		for i := 1; int32(i) <= d; i++ {

			if d%int32(i) == 0 {
				keys[i]++
				if keys[i] == len(input) {
					list = append(list, i)
				}
			}
		}
	}
	return int32(list[len(list)-1])
}
