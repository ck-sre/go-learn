package main

import (
	"fmt"
	"github.com/ck-sre/go-learn/cats-bench-test/cat"
)

type feline struct {
	name string
	age  int
}

func main() {
	murica := feline{
		name: "Murica",
		age:  cat.Years(10),
	}
	fmt.Println(murica)
	fmt.Println(cat.YearsSum(20))
}
