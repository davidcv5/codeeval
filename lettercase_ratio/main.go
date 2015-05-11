package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var lower = "abcdefghijklmnopqrstuvwxyz"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(ratio(scanner.Text()))
	}
}

func ratio(s string) string {
	u := 0
	l := 0
	for _, r := range []rune(s) {
		if i := strings.IndexRune(lower, r); i >= 0 {
			l++
		} else {
			u++
		}
	}
	return fmt.Sprintf("lowercase: %.2f uppercase: %.2f", float64(l*100)/float64(len(s)), float64(u*100)/float64(len(s)))
}
