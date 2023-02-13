package main

import "fmt"

func main() {
	switch "moments" {
	case "passItThrough":
		fmt.Println("following through")
	case "momentswegothroug":
		fmt.Println("passed moments here")
		fallthrough
	case "minutes":
		fmt.Println("reached after moments fell through")
	case "hours", "days":
		fmt.Println("Shouldn't ever reach here")
	default:
		fmt.Println("wassup defaults")
	}
	SwitchOnType(5)
}
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("tha int")
	case string:
		fmt.Println("tha string")
	default:
		fmt.Println("unknown")
	}
}
