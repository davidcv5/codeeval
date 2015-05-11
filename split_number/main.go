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
		fmt.Println(split(scanner.Text()))
	}
}

func split(s string) int {
	parts := strings.Fields(s)
	i := strings.Index(parts[1], "+")
	op := ""
	if i > 0 {
		op = "+"
	} else {
		i = strings.Index(parts[1], "-")
		op = "-"
	}
	x, _ := strconv.Atoi(parts[0][:i])
	y, _ := strconv.Atoi(parts[0][i:])
	if op == "+" {
		return x + y
	}
	return x - y
}
