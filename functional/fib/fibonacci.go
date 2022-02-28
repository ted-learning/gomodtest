package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Fibonacci2() Fin {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type Fin func() int

func (f Fin) Read(p []byte) (n int, err error) {
	next := f()
	if next > 100000 {
		return 0, io.EOF
	}
	return strings.NewReader(fmt.Sprintf("%d\n", next)).Read(p)
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := Fibonacci()

	fmt.Println(f()) //1
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //5
	fmt.Println(f()) //8
	fmt.Println(f()) //13
	fmt.Println(f()) //21
	fmt.Println(f()) //34

	printContents(Fibonacci2())
}
