package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *atomicInt) increase() {
	func() {
		i.lock.Lock()
		defer i.lock.Unlock()
		i.value++
	}()
}

func (i *atomicInt) get() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	a := atomicInt{0, sync.Mutex{}}
	go a.increase()
	a.increase()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
