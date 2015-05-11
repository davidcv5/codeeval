package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(period(scanner.Text()))
	}
}

func period(s string) int {
	for i := 1; i <= len(s)/2; i++ {
		match := true
		for j := 0; j < len(s)/i; j++ {
			for k := 0; k < len(s)/i; k++ {
				if s[j] != s[i*k] {
					match = false
					break
				}
			}
			if match == true {
				return i
			}
		}
		if match == true {
			return i
		}
	}
	return len(s)
}
