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
		fmt.Println(major(scanner.Text()))
	}
}

func major(s string) string {
	n := strings.Split(s, ",")
	m := make(map[string]int)
	for _, v := range n {
		m[v]++
		if m[v] > len(n)/2 {
			return v
		}
	}
	return "None"
}
