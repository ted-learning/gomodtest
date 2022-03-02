package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genMsg(service string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("Service: %s, send: %d", service, i)
			i++
		}
	}()
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v
			}
		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func main() {
	msg1 := genMsg("s1")
	msg2 := genMsg("s2")

	c1 := fanInBySelect(msg1, msg2)
	c2 := fanIn(msg1, msg2)
	go func() {
		for {
			fmt.Println("c1", <-c1)
		}
	}()

	for {
		fmt.Println("c2", <-c2)
	}
}
