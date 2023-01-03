package main

import (
	"errors"
	"fmt"
)

// var ErrCustom error
var ErrCustom = errors.New("custom error type for negative sqrt")

func main() {
	fmt.Printf("%T\n", ErrCustom)
	_, err := sqrt(-10)
	fmt.Println(err)
}

func sqrt(n int) (float64, error) {
	if n < 0 {
		return 0, ErrCustom
	}
	return 42, nil
}
