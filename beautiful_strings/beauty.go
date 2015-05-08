package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		fmt.Println(beauty(scanner.Text()))
	}
}

func beauty(s string) int {
	m := make(map[rune]int)
	for _, r := range []rune(strings.ToLower(s)) {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", r) {
			m[r]++
		}
	}
	c := []int{}
	for _, v := range m {
		c = append(c, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(c)))
	total := 0
	for i, v := range c {
		total += (26 - i) * v
	}
	return total
}
