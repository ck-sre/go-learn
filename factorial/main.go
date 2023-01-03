package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	fmt.Println("hello")
	return total
}
