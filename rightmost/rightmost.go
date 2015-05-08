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
	for scanner.Scan() {
		in := scanner.Text()
		if len(in) != 0 {
			fmt.Println(rightmost(in))
		}
	}
}

func rightmost(in string) int {
	parts := strings.Split(in, ",")
	a := parts[0]
	b := parts[1]
	for i := len(a) - 1; i >= 0; i-- {
		if string(a[i]) == b {
			return i
		}
	}
	return -1
}
