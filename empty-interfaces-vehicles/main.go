package main

import (
	"fmt"
)

type plane struct {
	vehicle
	Jet bool
}

type car struct {
	vehicle
	Wheels int
}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    int
}

type vehicles interface{}

func (v vehicle) specs() {
	fmt.Printf("Seats are %v, color is %v", v.Seats, v.Color)
}

func main() {
	prius := car{}
	tacoma := car{}
	boeing := plane{}
	rides := []vehicles{prius, tacoma, boeing}
	for key, value := range rides {
		fmt.Println("for vehicle", "key - ", key, "value - ", value)
	}
}
