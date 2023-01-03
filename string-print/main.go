package main

import (
	"fmt"
	"sort"
)

//cmd shift e
//cmd e
//ctrl tab
//ctrl opt r

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Jenny"}
	fmt.Println("here's the initial list", studyGroup)
	sort.Sort(studyGroup)
	fmt.Println("here's the sorted list", studyGroup)
}
