package main

import "fmt"

type Nodes []Node

type Node struct {
	a     int
	b     int
	group int
}

func main() {
	A := []int{1, 2, 3, 3}
	B := []int{2, 3, 1, 4}
	// A := []int{1, 2, 4, 5}
	// B := []int{2, 3, 5, 6}
	fmt.Println(Solution(A, B, 4))
}

func Solution(A []int, B []int, N int) int {
	count := 0
	iNodes := Nodes{}
	group := []int{}

	for i := 0; i < (len(A)+len(B))/2; i++ {
		iNode := Node{a: A[i], b: B[i]}
		iNodes = append(iNodes, iNode)
	}

	fmt.Println(iNodes)

	for i := 0; i < len(iNodes); i++ {

		for j := 0; j < len(iNodes); j++ {

			if i != j {

				if iNodes[i].a == iNodes[j].a || iNodes[i].a == iNodes[j].b || iNodes[i].b == iNodes[j].a || iNodes[i].b == iNodes[j].b {

					if iNodes[i].group == 0 && iNodes[j].group == 0 {
						count++
						group = append(group, count)
						iNodes[i].group, iNodes[j].group = count, count
					} else if iNodes[i].group == 0 && iNodes[j].group != 0 {
						iNodes[j].group = iNodes[i].group
					} else if iNodes[i].group != 0 && iNodes[j].group == 0 {
						iNodes[j].group = iNodes[i].group
					}
				}
			}
		}
	}

	fmt.Println(iNodes)
	fmt.Println(group)

	return count
}
