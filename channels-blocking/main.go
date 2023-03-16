package main

import "fmt"

func main() {

	c := make(chan int)

	go sender(c)
	receiver(c)

	//fmt.Println(len(x))

}

func sender(d chan<- int) {
	d <- 4
}

func receiver(e <-chan int) {
	fmt.Println(<-e)
}
