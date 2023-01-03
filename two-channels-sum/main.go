package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	sum1 := puller(c1)
	sum2 := puller(c2)
	fmt.Println("Final result: ", <-sum1+<-sum2) // <-sum1+<-sum2 => this is supposed to be blocking main from exiting
}

func incrementor(s string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	var sum int
	go func() {
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
