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
		u := strings.ToUpper(s)
		result := ""
		for i, _ := range u {
			if s[i] == u[i] {
				result += strings.ToLower(string(u[i]))
			} else {
				result += string(u[i])
			}
		}
		fmt.Println(result)
	}
}
