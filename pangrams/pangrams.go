package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(pangrams(scanner.Text()))
	}
}

func pangrams(s string) string {
	runes := []rune(strings.ToLower(s))
	m := make(map[rune]bool)
	for _, r := range runes {
		m[r] = true
	}
	missing := ""
	runes = []rune(alphabet)
	for _, r := range runes {
		if _, ok := m[r]; ok {
			continue
		}
		missing += string(r)
	}
	if len(missing) == 0 {
		return "NULL"
	}
	return missing
}
