package main

import "fmt"

type myErr struct {
	msg string
}

func (me myErr) Error() string {
	return fmt.Sprintf("err is %v", me.msg)
}

func main() {
	re := myErr{
		"message of myerror from Main",
	}
	fMyErr(re)
}

func fMyErr(r error) {
	fmt.Println("error is ", r, "\n")
	fmt.Println("error info using assertion is ", r.(myErr).msg, "\n")

}
