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
		fmt.Println(recover(scanner.Text()))
	}
}

func recover(s string) string {
	parts := strings.Split(s, ";")
	source := strings.Fields(parts[0])
	dest := make([]string, len(source))
	copy(dest, source)
	pos := strings.Fields(parts[1])
	for i, v := range pos {
		x, _ := strconv.Atoi(v)
		dest[x-1] = source[i]
	}
	n := missing(pos)
	dest[n-1] = source[len(source)-1]
	return strings.Join(dest, " ")
}

func missing(s []string) int {
	n := (len(s) + 1) * (len(s) + 2) / 2
	for _, r := range s {
		x, _ := strconv.Atoi(r)
		n -= x
	}
	return n
}
