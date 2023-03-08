package main

import "fmt"

func main() {
	h()
	fmt.Println("returned normally from h")
}

func h() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("recovered in h", h)
		}
	}()
	fmt.Println("calling k")
	k(0)
	fmt.Println("returned normally from k")
}

func k(l int) {
	if l > 3 {
		fmt.Sprintf("%v Pankckign", l)
		panic(fmt.Sprintf("%v", l))
	}
	defer fmt.Println("defer in k", l)
	fmt.Println("Printing in k", l)
	k(l + 1)

}
