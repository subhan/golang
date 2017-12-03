package main

import (
	"fmt"
	"os"
)

func main() {

	println("it's over 900")
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("it's not over", os.Args[1])
}
