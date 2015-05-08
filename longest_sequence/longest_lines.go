package main

import (
	"fmt"
	"strings"
)
import "log"
import "bufio"
import "os"

var count int

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
		fmt.Println(longest("", x, y))
	}
	fmt.Println("total steps: ", count)
}

func longest(current, x, y string) string {
	fmt.Printf("%s - %s - %s\n", current, x, y)
	count++

	if len(x) == 0 || len(y) == 0 {
		return current
	}
	if x[0] == y[0] {
		return longest(current+string(x[0]), x[1:], y[1:])
	}
	a := longest("", x[1:], y)
	b := longest("", x, y[1:])
	if len(a) > len(b) {
		return current + a
	}
	return current + b
}
