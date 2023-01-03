package main

import "fmt"

type Person struct {
	First string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("A regular person hi")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Hello to the secret agent i")
}

func main() {
	p1 := Person{
		First: "Michael",
		Age:   900,
	}
	p2 := DoubleZero{
		Person: Person{
			First: "James Bond",
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()

}
