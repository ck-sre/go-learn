package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	var an1, an2 string

	fmt.Print("Animal: ")
	_, err = fmt.Scan(&an1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Food:")
	_, err = fmt.Scan(&an2)
	if err != nil {
		panic(err)
	}
	fmt.Println(an1, an2)

	f, err := os.Create("./animals.example")
	if err != nil {
		fmt.Println(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	r := strings.NewReader(strings.Join([]string{an1, an2}, " "))
	_, err = io.Copy(f, r)
	if err != nil {
		return
	}

	k, err := os.Open("./animals.example")
	defer func(k *os.File) {
		err := k.Close()
		if err != nil {

		}
	}(k)
	if err != nil {
		fmt.Println(err)
	}

	bs, err := io.ReadAll(k)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
