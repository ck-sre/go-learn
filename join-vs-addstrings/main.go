package main

import (
	"fmt"
	"join-vs-addstrings/joinerstr"
)

func main() {
	var (
		k = []string{"girish", "barbat"}
	)
	fmt.Println(joinerstr.AddStr(k))
	fmt.Println(joinerstr.Joint(k))
}
