package main

import "fmt"

func main() {
	f := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			f <- i
		}
		close(f)
	}()

	for k := range f {
		fmt.Println(k)
	}
	fmt.Println("done sending and receiving")
}
