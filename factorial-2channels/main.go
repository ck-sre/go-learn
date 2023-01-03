package main

import "fmt"

func main() {
	c := decrementor(4)
	for d := range multiplier(c) {
		fmt.Println(d)
	}
}

func decrementor(b int) <-chan int {
	out := make(chan int)
	go func(f int) {
		for i := f; i > 0; i-- {
			out <- i
		}
		close(out)
	}(b)
	return out

}

func multiplier(c <-chan int) <-chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := range c {
			total *= i

		}
		out <- total
		close(out)
	}()
	return out
}
