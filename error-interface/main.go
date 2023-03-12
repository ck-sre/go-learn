package main

import (
	"errors"
	"fmt"
	"log"
)

type reduxError struct {
	sev string
	err error
}

func (k reduxError) Error() string {
	return fmt.Sprintf("an error of sev %v occurred: %v", k.sev, k.err)
}

func main() {
	_, err := cube(-22)

	if err != nil {
		log.Fatal(err)
	}
}

func cube(f float64) (float64, error) {
	if f < 0 {
		fmt.Printf("%T\n", errors.New("headlessmath"))
		//return 0, errors.New("headless math")
		//return 0, fmt.Errorf("headless math:  %v", f)
		re := fmt.Errorf("mindless error being thrown, this is the err object")
		return 0, reduxError{"error", re}
	}
	return 48, nil
}
