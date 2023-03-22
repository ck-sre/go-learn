package main

import "fmt"

func main() {
	t := make(chan int)
	nt := make(chan int)
	q := make(chan int)

	go sendThrees(t, nt, q)
	getThrees(t, nt, q)
	fmt.Println("above is all")
}

func sendThrees(tr, ntr, qr chan<- int) {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			tr <- i
		} else {
			ntr <- i
		}
	}
	qr <- 0
}

func getThrees(tg, ntg, qg <-chan int) {
	for {
		select {
		case i := <-tg:
			fmt.Println("from the multiple of 3 ", i)
		case i := <-ntg:
			fmt.Println("from a non multiple of 3 ", i)
		case i := <-qg:
			fmt.Println("quit signal received ", i)
			return
		}
	}

}
