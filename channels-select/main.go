package main

import "fmt"

func main() {
	c := gen()
	receive(c)
	fmt.Println("about to exit")
}

func receive(d chan int) {

	for i := range d {
		fmt.Println(i)
	}

}

func gen() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	close(d)

	return c
}
