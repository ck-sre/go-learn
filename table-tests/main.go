package main

import "fmt"

func MySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum + 33
}

func main() {
	fmt.Println(MySum(3, 4))
}
