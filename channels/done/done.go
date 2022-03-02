package main

import (
	"fmt"
	"sync"
)

type worker struct {
	id   int
	in   chan int
	done func()
}

func doWorker(w worker) {
	func() {
		for received := range w.in {
			fmt.Printf("Worker %d, received %d\n", w.id, received)
			w.done()
		}
	}()
}

func createWorker(id int, done func()) worker {
	c := make(chan int)
	worker := worker{
		id:   id,
		in:   c,
		done: done,
	}
	go doWorker(worker)
	return worker
}

func chanDemo() {
	wg := sync.WaitGroup{}
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, func() {
			wg.Done()
		})
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
