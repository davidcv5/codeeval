package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		s := scanner.Text()
		fmt.Println(list(s))
	}
}

func list(s string) string {
	x, _ := strconv.Atoi(s[:strings.Index(s, ",")])
	m := make(map[rune]bool)
	for _, r := range s[strings.Index(s, ",")+1:] {
		m[r] = true
	}
	return ""
}
