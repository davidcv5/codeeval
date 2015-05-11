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
	pos := -1
	for scanner.Scan() {
		s := scanner.Text()
		c := strings.Index(s, "C")
		g := strings.Index(s, "_")
		if c >= 0 {
			fmt.Println(replace(s, pos, c))
			pos = c
		} else {
			fmt.Println(replace(s, pos, g))
			pos = g
		}
	}
}

func replace(s string, pos, current int) string {
	if pos == -1 || pos == current {
		return s[:current] + "|" + s[current+1:]
	}
	if pos < current {
		return s[:current] + "\\" + s[current+1:]
	}
	return s[:current] + "/" + s[current+1:]
}
