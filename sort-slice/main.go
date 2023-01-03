package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

//type ByAge []person

func (p person) String() string {
	return fmt.Sprintf("This is the new custom string type %s: %d", p.Name, p.Age)
}

//func (p people) Len() int           { return len(p) }
//func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	people := []person{
		{Name: "Bob", Age: 31},
		{Name: "Nina", Age: 3},
	}
	fmt.Println("From main:", people[0])
}
