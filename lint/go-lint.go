package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(x int) string {
	if x > 10 {
		return fmt.Sprint("x is greater than 10")
	}
	return fmt.Sprint("less than")

}
