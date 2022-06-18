package main

type Input struct {
    X, Y, D int
}

func main() {
	// Data Testing
    DataTesting := []Input{
        {10, 85, 30},
        {10, 40, 30},
        {10, 40, 40},
        {30, 30, 40},
    }

    for _, d := range DataTesting {
        Solution(d.X, d.Y, d.D)
    }
}

func Solution(X int, Y int, D int) int {
    // write your code in Go 1.4
    distance := Y-X
    if distance <= 0 {
        return 0
    }
    
    if distance < D {
        return 1
    }

    if distance%D == 0 {
        return distance/D
    }

    r := distance/D+1
    return r
}