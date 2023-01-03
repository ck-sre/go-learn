package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)

	for n := range merge(c0, c1, c2) {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out

}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(clist ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()

	}

	wg.Add(len(clist))

	for _, c := range clist {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
