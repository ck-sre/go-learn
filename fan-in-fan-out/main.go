package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(msgP("Foo"), msgP("Bar"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting...")

}

func msgP(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FanIN function
func fanIn(inp1, inp2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-inp1
		}
	}()

	go func() {
		for {
			c <- <-inp2
		}
	}()
	return c
}
