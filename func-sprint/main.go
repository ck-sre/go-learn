package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, "of", lname), fmt.Sprint(lname, "of", fname)
}
