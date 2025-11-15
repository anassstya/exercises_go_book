package main

import (
	"fmt"
	"os"
)

func main() {
	for i, l := range os.Args[0:] {
		fmt.Println("index:", i, "value:", l)
	}
}
