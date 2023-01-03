package main

import (
	"fmt"
	"sort"
)

//cmd shift e
//cmd e
//ctrl tab
//ctrl opt r

func main() {
	sg := []string{"Zeno", "John", "Jenny", "Rohith", "Shetty", "Ali"}
	fmt.Println(sg)
	s := sort.Reverse(sort.StringSlice(sg))
	sort.Sort(s)
	fmt.Printf("here's the Type of the result %T\n", s)
	fmt.Println("here's the sorted list", s)
}
