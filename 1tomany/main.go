package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool) // semaphore

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

}
