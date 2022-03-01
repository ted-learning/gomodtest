package test

import (
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	data := []struct {
		a, b, c int
	}{
		{0, 1, 1},
		{1, 2, 3},
		{0, 0, 0},
		{-10, 20, 10},
		{math.MaxInt, 1, math.MinInt},
	}

	for _, test := range data {
		if i := add(test.a, test.b); i != test.c {
			t.Errorf("test error: %d + %d != %d\n", test.a, test.b, test.c)
		}
	}

}

func TestTriangle(t *testing.T) {
	givens := []struct{ p1, p2, result int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, given := range givens {
		if result := triangle(given.p1, given.p2); result != given.result {
			t.Errorf("calculate triangle %d and %d, the actual result %d is not as expect: %d", given.p1, given.p2, result, given.result)
		}
	}
}

func Example_add() {
	sum := add(1, 2)
	fmt.Println(sum)

	//output:
	//3
}
