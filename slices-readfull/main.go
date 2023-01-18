package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nonce [20]byte
	fmt.Println(nonce)

	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
	io.ReadAtLeast(rand.Reader, nonce[:], 19)
	fmt.Println(nonce)
}
