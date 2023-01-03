package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
