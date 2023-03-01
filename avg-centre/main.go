package main

import (
	avg_centred "avg-centre/avg-centred"
	"fmt"
)

func main() {
	fmt.Println(avg_centred.AvgCentre([]int{1, 2, 3, 4}))
	fmt.Println(avg_centred.AvgCentre([]int{4, 5, 1, 4}))
}
