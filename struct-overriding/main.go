package main

import "fmt"

type Person struct {
	First string
	Age   int
}

func main() {
	p1 := Person{
		First: "Michael",
		Age:   900,
	}
	fmt.Println(p1.First)
	fmt.Println("%T\n", p1)

}
