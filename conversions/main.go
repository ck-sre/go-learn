package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x rune = 'b'
	var y int32 = 'c'
	fmt.Println(x, string(x))
	fmt.Println(y, string(y))
	fmt.Println([]byte{'h', 'e'})
	fmt.Println(string([]byte{'h', 'e'}))
	fmt.Println([]byte("hello"))
	z, e := strconv.Atoi("d")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(z)
	fmt.Println(strconv.Itoa(49))

	// Assertions

	var name interface{} = 12
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Println("not a string")
	}
	fmt.Println(name.(int) - 3)
	fmt.Printf("%T\n", name.(int))

}
