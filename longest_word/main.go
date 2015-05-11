package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(longest(scanner.Text()))
	}
}

func longest(s string) string {
	parts := strings.Fields(s)
	m := make(map[int]string)
	for _, v := range parts {
		if _, ok := m[len(v)]; !ok {
			m[len(v)] = v
		}
	}
	longest := 0
	word := ""
	for k, v := range m {
		if k > longest {
			longest = k
			word = v
		}
	}
	return word
}
