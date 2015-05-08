package main

import (
	"fmt"
	"strings"
)
import "log"
import "bufio"
import "os"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		parts := strings.Split(in, ", ")
		fmt.Println(scrub(parts[0], parts[1]))
	}
}

func scrub(source, scrub string) string {
	result := ""
	for _, r := range source {
		if strings.ContainsRune(scrub, r) {
			continue
		}
		result += string(r)
	}
	return result
}
