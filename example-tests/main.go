// Test if example works in the documentation
package main

import "fmt"

// MySum Sum adds everything passed in
func MySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(MySum(3, 4))
}
