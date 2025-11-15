package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)

	for _, file := range os.Args[1:] {
		w, err := os.Open(file)
		if err != nil {
			continue
		}

		input := bufio.NewScanner(w)
		for input.Scan() {
			if counts[input.Text()] == nil {
				counts[input.Text()] = make(map[string]bool)
			}
			counts[input.Text()][file] = true
		}
		w.Close()
	}

	for line, linePath := range counts {
		if len(linePath) > 1 {
			fmt.Println("Строка", line, "встречается в нескольких файлах")
			fmt.Println("Названия файлов: ")
			for k, _ := range linePath {
				fmt.Println(k, " ")
			}
			fmt.Println(" ")
		}
	}
}
