package main

import "fmt"

func main() {
	c := make(<-chan int, 1)
	h := make(chan int, 1)
	f := make(chan<- int, 1)
	g := make(chan int, 1)
	f = h
	c = g
	fmt.Println(c, f)
	d := make(chan int, 1)
	//e := <-c
	d <- 4
	fmt.Println(<-d)
}
