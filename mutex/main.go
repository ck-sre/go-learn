package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3) * int(time.Millisecond)))

		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}
