package test

import "math"

func add(a, b int) int {
	return a + b
}

func triangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
