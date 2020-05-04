package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Base Command: %s\n", os.Args[0])
	for i, c := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, c)
	}
}
