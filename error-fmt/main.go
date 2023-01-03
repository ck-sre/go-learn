package main

import (
	"errors"
	"fmt"
)

// var ErrCustom error
var ErrCustom = errors.New("custom error type for negative sqrt")

type SqrtError struct {
	number int
	err    error
}

func (sqrterror *SqrtError) Error() string {
	return fmt.Sprintf("There you go, a sqrt error with the number %v, and the error string %v", sqrterror.number, sqrterror.err)
}

func main() {
	fmt.Printf("%T\n", ErrCustom)
	_, err := sqrt(-10)
	fmt.Println(err)
}

func sqrt(n int) (float64, error) {
	if n < 0 {
		downstreamError := fmt.Errorf("Error from downstream func")
		return 0, &SqrtError{n, downstreamError}
	}
	return 42, nil
}
