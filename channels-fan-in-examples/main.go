package main

import (
	"fmt"
	"sync"
)

func main() {
	m3 := make(chan int)
	nm3 := make(chan int)
	p := make(chan int)

	go spreadChans(m3, nm3)
	go getToPipe(m3, nm3, p)
	for k := range p {
		fmt.Println(k)
	}
	fmt.Println("exiting")
}

func spreadChans(m3, nm3 chan<- int) {
	for i := 0; i < 15; i++ {
		if i%3 == 0 {
			m3 <- i
		} else {
			nm3 <- i
		}
	}
	close(m3)
	close(nm3)
}

func getToPipe(m3, nm3 <-chan int, p chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range m3 {
			p <- v
		}
		wg.Done()
	}()

	go func() {
		for y := range nm3 {
			p <- y
		}
		wg.Done()
	}()

	wg.Wait()
	close(p)
}
