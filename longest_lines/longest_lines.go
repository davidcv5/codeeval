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
	var (
		a    []string
		x, y string
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		a = strings.Split(line, ";")
		x = a[0]
		y = a[1]
		if len(x) > 50 || len(y) > 50 {
			continue
		}
		if len(x) < len(y) {
			fmt.Println(longest("", x, y))
		} else {
			fmt.Println(longest("", y, x))
		}
	}
}

func longest(current, x, y string) string {
	if len(x) == 0 || len(y) == 0 {
		return current
	}
	if x[0] == y[0] {
		return longest(current+string(x[0]), x[1:], y[1:])
	}
	a := longest(current, x[1:], y)
	b := longest(current, y[1:], x)
	if len(a) > len(b) {
		return a
	}
	return b
}
