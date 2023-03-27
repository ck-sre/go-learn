package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fo := make(chan int)
	fi := make(chan int)

	go fillChan(fo)
	go spreadOutIn(fo, fi)

	for w := range fi {
		fmt.Println(w)
	}
	fmt.Println("exiting")
}

func fillChan(x chan int) {
	for i := 0; i < 90; i++ {
		x <- i
	}
	close(x)
}

func spreadOutIn(fo, fi chan int) {
	var wg sync.WaitGroup
	const numGos = 15
	wg.Add(numGos)

	for i := 0; i < numGos; i++ {
		go func() {
			for j := range fo {
				func(y int) {
					fi <- heavyComputing(y)
				}(j)
			}
			wg.Done()
			close(fi)
		}()
	}
	wg.Wait()
	close(fo)
}

func heavyComputing(z int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(5)))
	return z + rand.Intn(10)
}
