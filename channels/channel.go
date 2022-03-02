package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	go func() {
		for {
			received, ok := <-c
			if !ok {
				break
			}
			fmt.Printf("Worker %d, received %d\n", id, received)
		}
	}()
}

func worker2(id int, c chan int) {
	go func() {
		for received := range c {
			fmt.Printf("Worker %d, received %d\n", id, received)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func bufferChanDemo() {
	c := make(chan int, 3)
	go worker2(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func closeChanDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	bufferChanDemo()
	closeChanDemo()
}
