package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("filename.txt")
	if err != nil {
		fmt.Println("error creating")
	}
	log.SetOutput(nf)

}

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal("no such file", err)
		fmt.Println("println called")
	}
}
