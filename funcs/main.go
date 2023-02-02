package main

import (
	"fmt"
)

var x int

func wrapper() func() int {
	//x := 0
	return func() int {
		x++
		return x
	}
}

func buna() {
	fmt.Println("Buna")
}

func zioua() {
	fmt.Println("zioua")
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func callbackExample(x []int, y func(int) bool) []int {
	ni := []int{}
	for _, n := range x {
		if y(n) {
			ni = append(ni, n)
		}
	}
	return ni
}

func visit(p []int, callback2 func(int)) {
	for _, n := range p {
		callback2(n)
	}
}

func makePrata() func() string {
	return func() string {
		return "Te rog"
	}
}

func aver(i ...int) float64 {
	var total int
	for _, n := range i {
		total += n
	}
	return float64(total / len(i))
}

func bookName(author, book string) (string, string) {
	return fmt.Sprint(book, author), fmt.Sprint(author, book)
}

func changeMe(z *int) {
	*z = 20
	fmt.Println(*z)
}

func main() {
	defer buna()
	zioua()
	fmt.Println(factorial(4))

	nx := callbackExample([]int{1, 2, 3}, func(n int) bool { return n > 2 })
	fmt.Println(nx)

	visit([]int{7, 5, 2, 90}, func(n int) { fmt.Println(n) })

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	greeting := func() {
		fmt.Println("Cu Placere")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	prata := makePrata()
	fmt.Println(prata())

	monkeys := []int{1, 2, 6}
	avg := aver(monkeys...)
	fmt.Println(avg)

	fmt.Println(bookName("Dan", "Angels"))

	age := 44
	changeMe(&age)
	fmt.Println(age)

}
