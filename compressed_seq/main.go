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
		fmt.Println(compressed(scanner.Text()))
	}
}

func compressed(s string) string {
	parts := strings.Fields(s)
	c := parts[0]
	n := 1
	result := ""
	for _, r := range parts[1:] {
		if r != c {
			result += strconv.Itoa(n) + " " + c + " "
			c = r
			n = 0
		}
		n++
	}
	result += strconv.Itoa(n) + " " + c + " "
	return strings.TrimSpace(result)
}
