package main

import "fmt"

func main() {
	//myGreeting := map[string]string{}
	//var newGreeting = make(map[string]string)
	//myGreeting["Tim"] = "hello"
	//newGreeting["Tammy"] = "how do you do"
	//fmt.Println(myGreeting)
	//fmt.Println(len(newGreeting))
	tzMap := map[string]int{
		"SGT": 8,
		"PST": 12,
	}

	if value, exists := tzMap["IST"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("doesn't exist ", value)
	}
}
