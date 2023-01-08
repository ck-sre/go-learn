package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "This eBook is for the use of anyone anywhere in the United States and\nmost other parts of the" +
		" world at no cost and with almost no restrictions\nwhatsoever"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(
			os.Stderr,
			"reading error",
			err,
		)
	}

}
