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
		fmt.Println(lowest(s))
	}
}

func lowest(s string) int {
	parts := strings.Fields(s)
	m := make(map[int]int)
	for i, v := range parts {
		x, _ := strconv.Atoi(v)
		if _, ok := m[x]; ok {
			m[x] = -1
		} else {
			m[x] = i + 1
		}
	}

	w := 10
	p := 0
	for k, v := range m {
		if v > 0 && k < w {
			w = k
			p = v
		}
	}
	if w == 10 {
		return 0
	}
	return p
}
