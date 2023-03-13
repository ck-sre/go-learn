package main

import "fmt"

func main() {
	x := make(chan string)
	go func() {
		x <- "hello"
	}()
	fmt.Println(<-x)

	y := make(chan int, 1)
	y <- 2
	fmt.Println(<-y)

}
