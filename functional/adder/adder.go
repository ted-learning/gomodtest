package main

import (
	"fmt"
)

func adder() func(i int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("a(%d)=%d\n", i, a(i))
	}

	fmt.Println("-----------------")

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var v int
		v, a2 = a2(i)
		fmt.Printf("a2(%d)=%d\n", i, v)
	}
}
