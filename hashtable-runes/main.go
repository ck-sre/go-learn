package main

import "fmt"

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	word := "Hello"
	letter = rune(word[0])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	for i := 65; i <= 122; i++ {
		fmt.Println(i, "-", string(i), "-", i%12)
	}
}
