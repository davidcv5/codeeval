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
		fmt.Println(swap(scanner.Text()))
	}
}

func swap(s string) string {
	parts := strings.Split(s, " : ")
	c := strings.Fields(parts[0])
	swaps := strings.Split(parts[1], ", ")
	for _, v := range swaps {
		p := strings.Split(v, "-")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		c[x], c[y] = c[y], c[x]
	}
	return strings.Join(c, " ")
}
