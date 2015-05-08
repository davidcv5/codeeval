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
		s := scanner.Text()
		fmt.Println(find(s))
	}
}

func find(s string) string {
	parts := strings.Split(s, "|")
	x := parts[0]
	c := strings.Fields(parts[1])
	result := ""
	for _, r := range c {
		i, _ := strconv.Atoi(r)
		result += string(x[i-1])
	}
	return result
}
