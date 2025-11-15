package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	notEffective()
	effective()
}

func notEffective() {
	start := time.Now()
	var s string
	for i, v := range os.Args[1:] {
		if i > 0 {
			s += ", "
		}
		s += v
	}
	fmt.Println("result:", s)
	fmt.Println("time:", time.Since(start))
}

func effective() {
	start := time.Now()
	fmt.Println("result:", strings.Join(os.Args[1:], ", "))
	fmt.Println("time:", time.Since(start))
}
