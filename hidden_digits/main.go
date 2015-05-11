package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var letters = "abcdefghij"
var numbers = "0123456789"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(hidden(scanner.Text()))
	}
}

func hidden(s string) string {
	result := ""
	for _, v := range s {
		i := strings.IndexRune(letters, v)
		if i >= 0 {
			result += strconv.Itoa(i)
		} else if strings.ContainsRune(numbers, v) {
			result += string(v)
		}
	}
	if len(result) > 0 {
		return result
	}
	return "NONE"
}
