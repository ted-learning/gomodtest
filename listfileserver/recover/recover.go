package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err.Error())
		} else {
			panic(fmt.Sprintf("I don't know what to do with error: %v", r))
		}
	}()

	panic("123")
	//panic(errors.New("this is a new error"))
	//a := 0
	//b := 5
	//fmt.Println(b / a)
}

func main() {
	tryRecover()
}
