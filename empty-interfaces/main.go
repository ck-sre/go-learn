package main

import "fmt"

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

type animal struct {
	sound string
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	critters := []interface{}{fifi, fido}
	specs(fido)
	specs(fifi)
	fmt.Println(critters)
}
