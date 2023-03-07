package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer checkFoo()
	f, err := os.Create("stdlog.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	fe, err := os.Open("unknown.txt")
	if err != nil {
		log.Fatal("error occurred ", err)
	}
	defer fe.Close()

	fmt.Println("check if the log is in the directory")
}

func checkFoo() {
	fmt.Println("checking if this defer ran")
	log.Println("checking if this defer ran")
}
