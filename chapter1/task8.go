package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
			fmt.Println(url)
		} else {
			fmt.Println("http://" + url)
		}
	}
}
