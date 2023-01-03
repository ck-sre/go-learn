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
	s := []string{"Zeno", "John", "AI", "Jenny"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println("Sorted from stringSlice", s)
}
