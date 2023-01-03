package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 2000; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 2000; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
