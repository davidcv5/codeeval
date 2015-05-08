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
		s := scanner.Text()
		if len(s) == 0 {
			continue
		}
		fmt.Println(trailing(s))
	}
}

func trailing(s string) int {
	parts := strings.Split(s, ",")
	if len(parts) == 0 || len(parts[0]) == 0 || len(parts[1]) == 0 || len(parts[0]) < len(parts[1]) {
		return 0
	}
	if strings.LastIndex(parts[0], parts[1]) == len(parts[0])-len(parts[1]) {
		return 1
	}
	return 0
}
