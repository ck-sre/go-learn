package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()

	d, ok := <-c
	fmt.Println(d, ok)
	d, ok = <-c
	fmt.Println(d, ok)
}
