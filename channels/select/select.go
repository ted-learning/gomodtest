package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	func() {
		for received := range c {
			time.Sleep(time.Second)
			fmt.Printf("Worker %d, received %d\n", id, received)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func genChan() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func main() {
	c1, c2 := genChan(), genChan()
	worker := createWorker(0)
	var values []int
	var activeWorker chan<- int
	var activeValue int
	end := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-end:
			fmt.Println("bye")
			return
		case <-tick:
			fmt.Println("values left:", len(values))
		case <-time.After(500 * time.Millisecond):
			fmt.Println("time out")
		}

	}
}
