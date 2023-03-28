package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	cx, breakFree := context.WithCancel(context.Background())

	fmt.Println("error check:", cx.Err())
	fmt.Println("total goroutines:", runtime.NumGoroutine())

	go func() {
		v := 0
		for {
			select {
			case <-cx.Done():
				return
			default:
				v++
				time.Sleep(time.Millisecond * 100)
				fmt.Println("Still chugging along", v)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("is there an error now?", cx.Err())
	fmt.Println("num routines next check:", runtime.NumGoroutine())

	fmt.Println("cancelling ctx now")
	breakFree()
	fmt.Println("cancelled cx")
	time.Sleep(time.Second * 2)
	fmt.Println("3rd error check", cx.Err())
	fmt.Println("num routines", runtime.NumGoroutine())
}
