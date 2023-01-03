package main

import (
	"bufio"
	"errors"
	"fmt"
)

// var ErrCustom error

func main() {
	fmt.Printf("%T\n", ErrCustom)
	_, err := sqrt(-10)
	fmt.Println(err)
}

func sqrt(n int) (float64, error) {

	if n < 0 {
		ErrCustom := fmt.Errorf("custom error type for negative sqrt%v", n)
		return 0, ErrCustom
	}
	return 42, nil
}
