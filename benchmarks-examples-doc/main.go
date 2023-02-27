package main

import (
	"benchmarks-examples-doc/quoter"
	"benchmarks-examples-doc/worder"
	"fmt"
)

func main() {
	fmt.Println(worder.Counter(quoter.Suntzu))
	for k, v := range worder.UseCount(quoter.Suntzu) {
		fmt.Println(v, k)
	}
}
