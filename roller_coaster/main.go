package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(roller(scanner.Text()))
	}
}

func roller(s string) string {
	runes := []rune(s)
	upper := true
	for i, r := range runes {
		if unicode.IsLetter(r) {
			if upper {
				runes[i] = unicode.ToUpper(r)
			} else {
				runes[i] = unicode.ToLower(r)
			}
			upper = !upper
		}
	}
	return string(runes)
}
