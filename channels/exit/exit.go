package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genMsg(service string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("Service: %s, send: %d", service, i)
				i++
			case <-done:
				fmt.Println("clean up started.")
				time.Sleep(time.Second)
				fmt.Println("clean up finished.")
				done <- struct{}{}
				return
			}
		}
	}()
	return c
}

func timeoutWaiting(c chan string, timeout time.Duration) (string, bool) {
	select {
	case val := <-c:
		return val, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	c := genMsg("s1", done)
	for i := 0; i < 10; i++ {
		if waiting, ok := timeoutWaiting(c, time.Second); ok {
			fmt.Println(waiting)
		} else {
			fmt.Println("timeout")
		}
	}

	done <- struct{}{}
	<-done
}
