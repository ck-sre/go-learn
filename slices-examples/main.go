package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:5])
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[:2])

	mySlice2 := make([]int, 0, 5)

	fmt.Println("-------------")
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println("-------------")

	for i := 0; i < 80; i++ {
		mySlice2 = append(mySlice2, i)
		fmt.Println("Len:", len(mySlice2), "Cap:", cap(mySlice2))
	}

	partWeekDays := []string{"Monday", "Tuesday"}
	otherPartWeekDays := []string{"Wednesday", "Thursday", "Friday"}

	fullDays := append(partWeekDays, otherPartWeekDays...)
	fmt.Println(fullDays)

	missingWeekDays := append(partWeekDays[:1], otherPartWeekDays[:]...)
	fmt.Println(missingWeekDays)

}
