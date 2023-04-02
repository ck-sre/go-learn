package main

import "fmt"

func main() {
	var x [50]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[2])
	for i := 560; i <= 600; i++ {
		x[i-560] = int(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))

}
