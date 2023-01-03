package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w, r
}

// cmd up
// opt enter - inject language reference, create var from implementation
// on function name, the above would give add  option
// opt up or down - selection scope
// ctrl G - multi-word refactor
// ctrl space - suggestions
// . and ctrl space - will give conversion suggestions - especially for variables
// cmd shift e  - this shows all recent file blobs - very interesting
// cmd e - switcher
// ctrl tab
// cmd f12 for structure of package
// ctrl h - call hierarchy
// ctrl T - refactor, multicursor change, - can change signature here
// cmd N - implement an interface
// ctrl T - can also abstract out an interface
// also can introduce a constant

type message struct {
	field1 int
	field2 string
	field3 string
}

func (m message) Method1() {

	fmt.Println("this is method1")
}

func (m message) Method2() {
	fmt.Println("this is method2")
}

func loopThrough() {
	for i := 0; i < 42; i++ {
		_, err := CheckTwenty(i)
		if err != nil {
			break
		}
		fmt.Println(i)
	}
}

//CheckTwenty returns int and error
func CheckTwenty(n int) (int, error) {
	if n == 21 {
		return 1, fmt.Errorf("its twenty one")
	}
	return n, nil
}

type demoType struct {
	Field1 int
}

// AddToField returns a Field1
func (d *demoType) AddToField(val int) int {
	d.Field1 += val
	return d.Field1
}

func main() {
	loopThrough()

	_ = handler
	_ = newHandler

	fmt.Println("running for newHandler")
	_ = myTypeStruct{
		field1: 1,
		field2: "hello",
	}
	msg := message{
		field1: 1,
		field2: "1",
		field3: "string2",
	}
	out, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal error")
	}
	fmt.Println(out)
	_ = addNumbers(2, 1)

	_ = demoType{}
}

func addNumbers(b int, a int) int {
	return a + b
}

type myTypeStruct struct {
	field1 int
	field2 string
}

// newHandler adds a ResponseWriter object to
func newHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w, r
}

type msgAbstract interface {
	Method1()
}
