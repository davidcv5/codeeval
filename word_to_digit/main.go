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
		fmt.Println(w2d(scanner.Text()))
	}
}

func w2d(s string) string {
	parts := strings.Split(s, ";")
	result := ""
	for _, d := range parts {
		result += digit(d)
	}
	return result
}

func digit(s string) string {
	switch s {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return ""
}
