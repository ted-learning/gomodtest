package main

import (
	"bufio"
	"fmt"
	"gomodtest/functional/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("stop")
		}
	}
}

func writeFibonacci2File(filename string) {
	createdFile, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL, 0666)
	//err = errors.New("Customer Error.")
	if err != nil {
		if p, ok := err.(*os.PathError); ok {
			fmt.Printf("Op: %s, path: %s, error: %s\n", p.Op, p.Path, p.Error())
		} else {
			panic(err)
		}
		return
	}
	defer handle(createdFile.Close)

	writer := bufio.NewWriter(createdFile)
	defer handle(writer.Flush)

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			panic(err)
		}
	}
}

func handle(f func() error) {
	err := f()
	if err != nil {
		panic(err)
	}
}

func main() {
	writeFibonacci2File("test_write_fibonacci.txt")
	tryDefer()
}
