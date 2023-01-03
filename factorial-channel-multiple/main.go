package main

import "fmt"

func main() {
	c := genNum(100)
	f := factorial(c)
	for d := range f {
		fmt.Println("for number ", d)
	}
}

func genNum(d int64) chan int64 {
	out := make(chan int64)
	go func() {
		for i := d; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(b chan int64) <-chan int64 {
	out := make(chan int64)
	go func() {
		for i := range b {
			out <- fact(i)
		}
		close(out)
	}()
	return out

}

func fact(g int64) int64 {
	var total int64
	total = 1
	for i := g; i > 0; i-- {
		total *= i
	}
	return total
}
