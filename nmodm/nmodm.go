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
		fmt.Println(mod(s))
	}
}

func mod(s string) int {
	parts := strings.Split(s, ",")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	for {
		x := n - m
		if x == 0 {
			return 0
		}
		if x < 0 {
			return n
		}
		n = x
	}
}
