package main

import "fmt"

func main() {
	c := gen(2, 3, 4)
	sqrt := sq(c)
	for x := range sqrt {
		fmt.Println(x)
	}
}

func gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range n {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}
