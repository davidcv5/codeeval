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
		fmt.Println(multiply(scanner.Text()))
	}
}

func multiply(s string) string {
	parts := strings.Split(s, " | ")
	p1 := strings.Fields(parts[0])
	p2 := strings.Fields(parts[1])
	result := make([]string, len(p1))
	for i := 0; i < len(p1); i++ {
		x, _ := strconv.Atoi(p1[i])
		y, _ := strconv.Atoi(p2[i])
		result[i] = strconv.Itoa(x * y)
	}
	return strings.Join(result, " ")
}
